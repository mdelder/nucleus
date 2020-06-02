package spoke

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"

	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/openshift/library-go/pkg/assets"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	operatorhelpers "github.com/openshift/library-go/pkg/operator/v1helpers"

	nucleusv1client "github.com/open-cluster-management/api/client/operator/clientset/versioned/typed/operator/v1"
	nucleusinformer "github.com/open-cluster-management/api/client/operator/informers/externalversions/operator/v1"
	nucleuslister "github.com/open-cluster-management/api/client/operator/listers/operator/v1"
	nucleusapiv1 "github.com/open-cluster-management/api/operator/v1"
	"github.com/open-cluster-management/nucleus/pkg/helpers"
	"github.com/open-cluster-management/nucleus/pkg/operators/spoke/bindata"
)

const (
	nucleusSpokeFinalizer        = "nucleus.open-cluster-management.io/spoke-core-cleanup"
	bootstrapHubKubeConfigSecret = "bootstrap-hub-kubeconfig"
	hubKubeConfigSecret          = "hub-kubeconfig-secret"
	nucleusKlusterletNamespace   = "open-cluster-management-agent"
	spokeCoreApplied             = "Applied"
	spokeRegistrationDegraded    = "SpokeRegistrationDegraded"
)

var (
	staticResourceFiles = []string{
		"manifests/spoke/spoke-registration-serviceaccount.yaml",
		"manifests/spoke/spoke-registration-clusterrole.yaml",
		"manifests/spoke/spoke-registration-clusterrolebinding.yaml",
		"manifests/spoke/spoke-registration-role.yaml",
		"manifests/spoke/spoke-registration-rolebinding.yaml",
		"manifests/spoke/spoke-work-serviceaccount.yaml",
		"manifests/spoke/spoke-work-clusterrole.yaml",
		"manifests/spoke/spoke-work-clusterrolebinding.yaml",
		"manifests/spoke/spoke-work-clusterrolebinding-addition.yaml",
	}
)

type nucleusSpokeController struct {
	nucleusClient          nucleusv1client.KlusterletInterface
	nucleusLister          nucleuslister.KlusterletLister
	kubeClient             kubernetes.Interface
	registrationGeneration int64
	workGeneration         int64
}

// NewNucleusSpokeController construct nucleus spoke controller
func NewNucleusSpokeController(
	kubeClient kubernetes.Interface,
	nucleusClient nucleusv1client.KlusterletInterface,
	nucleusInformer nucleusinformer.KlusterletInformer,
	recorder events.Recorder) factory.Controller {
	controller := &nucleusSpokeController{
		kubeClient:    kubeClient,
		nucleusClient: nucleusClient,
		nucleusLister: nucleusInformer.Lister(),
	}

	return factory.New().WithSync(controller.sync).
		WithInformersQueueKeyFunc(func(obj runtime.Object) string {
			accessor, _ := meta.Accessor(obj)
			return accessor.GetName()
		}, nucleusInformer.Informer()).
		ToController("NucleusSpokeController", recorder)
}

// spokeConfig is used to render the template of hub manifests
type spokeConfig struct {
	KlusterletName            string
	KlusterletNamespace       string
	RegistrationImage         string
	WorkImage                 string
	ClusterName               string
	ExternalServerURL         string
	HubKubeConfigSecret       string
	BootStrapKubeConfigSecret string
}

func (n *nucleusSpokeController) sync(ctx context.Context, controllerContext factory.SyncContext) error {
	spokeCoreName := controllerContext.QueueKey()
	klog.V(4).Infof("Reconciling Klusterlet %q", spokeCoreName)
	spokeCore, err := n.nucleusLister.Get(spokeCoreName)
	if errors.IsNotFound(err) {
		// AgentCore not found, could have been deleted, do nothing.
		return nil
	}
	if err != nil {
		return err
	}
	spokeCore = spokeCore.DeepCopy()

	config := spokeConfig{
		KlusterletName:            spokeCore.Name,
		KlusterletNamespace:       spokeCore.Spec.Namespace,
		RegistrationImage:         spokeCore.Spec.RegistrationImagePullSpec,
		WorkImage:                 spokeCore.Spec.WorkImagePullSpec,
		ClusterName:               spokeCore.Spec.ClusterName,
		BootStrapKubeConfigSecret: bootstrapHubKubeConfigSecret,
		HubKubeConfigSecret:       hubKubeConfigSecret,
		ExternalServerURL:         getServersFromKlusterlet(spokeCore),
	}
	// If namespace is not set, use the default namespace
	if config.KlusterletNamespace == "" {
		config.KlusterletNamespace = nucleusKlusterletNamespace
	}

	// Update finalizer at first
	if spokeCore.DeletionTimestamp.IsZero() {
		hasFinalizer := false
		for i := range spokeCore.Finalizers {
			if spokeCore.Finalizers[i] == nucleusSpokeFinalizer {
				hasFinalizer = true
				break
			}
		}
		if !hasFinalizer {
			spokeCore.Finalizers = append(spokeCore.Finalizers, nucleusSpokeFinalizer)
			_, err := n.nucleusClient.Update(ctx, spokeCore, metav1.UpdateOptions{})
			return err
		}
	}

	// Klusterlet is deleting, we remove its related resources on spoke
	if !spokeCore.DeletionTimestamp.IsZero() {
		if err := n.cleanUp(ctx, controllerContext, config); err != nil {
			return err
		}
		return n.removeWorkFinalizer(ctx, spokeCore)
	}

	// Start deploy spoke core components
	// Check if namespace exists
	_, err = n.kubeClient.CoreV1().Namespaces().Get(ctx, config.KlusterletNamespace, metav1.GetOptions{})
	switch {
	case errors.IsNotFound(err):
		_, createErr := n.kubeClient.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{Name: config.KlusterletNamespace},
		}, metav1.CreateOptions{})
		if createErr != nil {
			helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
				Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
				Message: fmt.Sprintf("Failed to create namespace %q: %v", config.KlusterletNamespace, createErr),
			}))
			return createErr
		}
	case err != nil:
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: fmt.Sprintf("Failed to get namespace %q: %v", config.KlusterletNamespace, err),
		}))
		return err
	}

	// Check if bootstrap secret exists
	_, err = n.kubeClient.CoreV1().Secrets(config.KlusterletNamespace).Get(
		ctx, config.BootStrapKubeConfigSecret, metav1.GetOptions{})
	if err != nil {
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: fmt.Sprintf("Failed to get bootstrap secret -n %q %q: %v", config.KlusterletNamespace, config.BootStrapKubeConfigSecret, err),
		}))
		return err
	}

	// Deploy the static resources
	// Apply static files
	resourceResults := resourceapply.ApplyDirectly(
		resourceapply.NewKubeClientHolder(n.kubeClient),
		controllerContext.Recorder(),
		func(name string) ([]byte, error) {
			return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
		},
		staticResourceFiles...,
	)
	errs := []error{}
	for _, result := range resourceResults {
		if result.Error != nil {
			errs = append(errs, fmt.Errorf("%q (%T): %v", result.File, result.Type, result.Error))
		}
	}

	if len(errs) > 0 {
		applyErrors := operatorhelpers.NewMultiLineAggregate(errs)
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: applyErrors.Error(),
		}))
		return applyErrors
	}

	// Create hub config secret
	hubSecret, err := n.kubeClient.CoreV1().Secrets(config.KlusterletNamespace).Get(ctx, hubKubeConfigSecret, metav1.GetOptions{})
	switch {
	case errors.IsNotFound(err):
		// Create an empty secret with placeholder
		hubSecret = &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      hubKubeConfigSecret,
				Namespace: config.KlusterletNamespace,
			},
			Data: map[string][]byte{"placeholder": []byte("placeholder")},
		}
		hubSecret, err = n.kubeClient.CoreV1().Secrets(config.KlusterletNamespace).Create(ctx, hubSecret, metav1.CreateOptions{})
		if err != nil {
			helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
				Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
				Message: fmt.Sprintf("Failed to create hub kubeconfig secret -n %q %q: %v", hubSecret.Namespace, hubSecret.Name, err),
			}))
			return err
		}
	case err != nil:
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: fmt.Sprintf("Failed to get hub kubeconfig secret with error %v", err),
		}))
		return err
	}

	// Deploy registration agent
	generation, err := helpers.ApplyDeployment(
		n.kubeClient,
		n.registrationGeneration,
		func(name string) ([]byte, error) {
			return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
		},
		controllerContext.Recorder(),
		"manifests/spoke/spoke-registration-deployment.yaml")
	if err != nil {
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: fmt.Sprintf("Failed to deploy registration deployment with error %v", err),
		}))
		return err
	}
	// TODO store this in the status of the spokecore itself
	n.registrationGeneration = generation

	// Deploy work agent
	generation, err = helpers.ApplyDeployment(
		n.kubeClient,
		n.workGeneration,
		func(name string) ([]byte, error) {
			return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
		},
		controllerContext.Recorder(),
		"manifests/spoke/spoke-work-deployment.yaml")
	if err != nil {
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeCoreApplied, Status: metav1.ConditionFalse, Reason: "KlusterletApplyFailed",
			Message: fmt.Sprintf("Failed to deploy work deployment with error %v", err),
		}))
		return err
	}
	// TODO store this in the status of the spokecore itself
	n.workGeneration = generation

	// if we get here, we have successfully applied everything and should indicate that
	helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
		Type: spokeCoreApplied, Status: metav1.ConditionTrue, Reason: "KlusterletApplied",
		Message: "Spoke Core Component Applied",
	}))

	// now that we have applied all of our logic, we can check to see if the data we expect to have present as indications of
	// proper functioning of registration controller is working
	// TODO this should be moved into a separate loop since it is independent of the application of the eventually consistent
	//  resources above

	// If cluster name is empty, read cluster name from hub config secret
	if config.ClusterName == "" {
		clusterName := hubSecret.Data["cluster-name"]
		if clusterName == nil {
			helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
				Type: spokeRegistrationDegraded, Status: metav1.ConditionTrue, Reason: "ClusterNameMissing",
				Message: fmt.Sprintf("Failed to get cluster name from `kubectl get secret -n %q %q -ojsonpath='{.data.cluster-name}`.  This is set by the spoke registration deployment.", hubSecret.Namespace, hubSecret.Name),
			}))
			return fmt.Errorf("Failed to get cluster name")
		}
		config.ClusterName = string(clusterName)
	}

	// If hub kubeconfig does not exist, return err.
	if hubSecret.Data["kubeconfig"] == nil {
		helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
			Type: spokeRegistrationDegraded, Status: metav1.ConditionTrue, Reason: "HubKubeconfigMissing",
			Message: fmt.Sprintf("Failed to get cluster name from `kubectl get secret -n %q %q -ojsonpath='{.data.kubeconfig}`.  This is set by the spoke registration deployment, but the CSR must be approved by the cluster-admin on the hub.", hubSecret.Namespace, hubSecret.Name),
		}))
		return fmt.Errorf("Failed to get kubeconfig from hub kubeconfig secret")
	}
	// TODO it is possible to verify the kubeconfig actually works.

	helpers.UpdateNucleusSpokeStatus(ctx, n.nucleusClient, spokeCoreName, helpers.UpdateNucleusSpokeConditionFn(nucleusapiv1.StatusCondition{
		Type: spokeRegistrationDegraded, Status: metav1.ConditionFalse, Reason: "RegistrationFunctional",
		Message: "Registration is managing credentials",
	}))
	return nil
}

func (n *nucleusSpokeController) cleanUp(ctx context.Context, controllerContext factory.SyncContext, config spokeConfig) error {
	// Remove deployment
	registrationDeployment := fmt.Sprintf("%s-registration-agent", config.KlusterletName)
	err := n.kubeClient.AppsV1().Deployments(config.KlusterletNamespace).Delete(ctx, registrationDeployment, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	controllerContext.Recorder().Eventf("DeploymentDeleted", "deployment %s is deleted", registrationDeployment)
	workDeployment := fmt.Sprintf("%s-work-agent", config.KlusterletName)
	err = n.kubeClient.AppsV1().Deployments(config.KlusterletNamespace).Delete(ctx, workDeployment, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	// Remove secret
	err = n.kubeClient.CoreV1().Secrets(config.KlusterletNamespace).Delete(ctx, config.HubKubeConfigSecret, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	controllerContext.Recorder().Eventf("SecretDeleted", "secret %s is deleted", config.HubKubeConfigSecret)

	// Remove Static files
	for _, file := range staticResourceFiles {
		err := helpers.CleanUpStaticObject(
			ctx,
			n.kubeClient,
			nil,
			nil,
			func(name string) ([]byte, error) {
				return assets.MustCreateAssetFromTemplate(name, bindata.MustAsset(filepath.Join("", name)), config).Data, nil
			},
			file,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *nucleusSpokeController) removeWorkFinalizer(ctx context.Context, deploy *nucleusapiv1.Klusterlet) error {
	copiedFinalizers := []string{}
	for i := range deploy.Finalizers {
		if deploy.Finalizers[i] == nucleusSpokeFinalizer {
			continue
		}
		copiedFinalizers = append(copiedFinalizers, deploy.Finalizers[i])
	}

	if len(deploy.Finalizers) != len(copiedFinalizers) {
		deploy.Finalizers = copiedFinalizers
		_, err := n.nucleusClient.Update(ctx, deploy, metav1.UpdateOptions{})
		return err
	}

	return nil
}

func readClusterNameFromSecret(secret *corev1.Secret) (string, error) {
	if secret.Data["cluster-name"] == nil {
		return "", fmt.Errorf("Unable to find cluster name in secret")
	}

	return string(secret.Data["cluster-name"]), nil
}

func readKubuConfigFromSecret(secret *corev1.Secret, config spokeConfig) (string, error) {
	if secret.Data["kubeconfig"] == nil {
		return "", fmt.Errorf("Unable to find kubeconfig in secret")
	}

	return string(secret.Data["kubeconfig"]), nil
}

// TODO also read CABundle from ExternalServerURLs and set into registration deployment
func getServersFromKlusterlet(spokeCore *nucleusapiv1.Klusterlet) string {
	if spokeCore.Spec.ExternalServerURLs == nil {
		return ""
	}
	serverString := make([]string, 0, len(spokeCore.Spec.ExternalServerURLs))
	for _, server := range spokeCore.Spec.ExternalServerURLs {
		serverString = append(serverString, server.URL)
	}
	return strings.Join(serverString, ",")
}
