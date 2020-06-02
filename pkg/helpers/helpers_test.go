package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	nucleusfake "github.com/open-cluster-management/api/client/operator/clientset/versioned/fake"
	nucleusapiv1 "github.com/open-cluster-management/api/operator/v1"
	"github.com/openshift/library-go/pkg/operator/events/eventstesting"
	operatorhelpers "github.com/openshift/library-go/pkg/operator/v1helpers"
	admissionv1 "k8s.io/api/admissionregistration/v1"
	fakeapiextensions "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"

	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	fakekube "k8s.io/client-go/kubernetes/fake"
	fakeapiregistration "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/fake"
)

func TestUpdateStatusCondition(t *testing.T) {
	nowish := metav1.Now()
	beforeish := metav1.Time{Time: nowish.Add(-10 * time.Second)}
	afterish := metav1.Time{Time: nowish.Add(10 * time.Second)}

	cases := []struct {
		name               string
		startingConditions []nucleusapiv1.StatusCondition
		newCondition       nucleusapiv1.StatusCondition
		expectedUpdated    bool
		expectedConditions []nucleusapiv1.StatusCondition
	}{
		{
			name:               "add to empty",
			startingConditions: []nucleusapiv1.StatusCondition{},
			newCondition:       newCondition("test", "True", "my-reason", "my-message", nil),
			expectedUpdated:    true,
			expectedConditions: []nucleusapiv1.StatusCondition{newCondition("test", "True", "my-reason", "my-message", nil)},
		},
		{
			name: "add to non-conflicting",
			startingConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
			},
			newCondition:    newCondition("one", "True", "my-reason", "my-message", nil),
			expectedUpdated: true,
			expectedConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
				newCondition("one", "True", "my-reason", "my-message", nil),
			},
		},
		{
			name: "change existing status",
			startingConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
				newCondition("one", "True", "my-reason", "my-message", nil),
			},
			newCondition:    newCondition("one", "False", "my-different-reason", "my-othermessage", nil),
			expectedUpdated: true,
			expectedConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
				newCondition("one", "False", "my-different-reason", "my-othermessage", nil),
			},
		},
		{
			name: "leave existing transition time",
			startingConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
				newCondition("one", "True", "my-reason", "my-message", &beforeish),
			},
			newCondition:    newCondition("one", "True", "my-reason", "my-message", &afterish),
			expectedUpdated: false,
			expectedConditions: []nucleusapiv1.StatusCondition{
				newCondition("two", "True", "my-reason", "my-message", nil),
				newCondition("one", "True", "my-reason", "my-message", &beforeish),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fakeClusterClient := nucleusfake.NewSimpleClientset(
				&nucleusapiv1.ClusterManager{
					ObjectMeta: metav1.ObjectMeta{Name: "testspokecluster"},
					Status: nucleusapiv1.ClusterManagerStatus{
						Conditions: c.startingConditions,
					},
				},
				&nucleusapiv1.SpokeCore{
					ObjectMeta: metav1.ObjectMeta{Name: "testspokecluster"},
					Status: nucleusapiv1.KlusterletStatus{
						Conditions: c.startingConditions,
					},
				},
			)

			hubstatus, updated, err := UpdateNucleusHubStatus(
				context.TODO(),
				fakeClusterClient.NucleusV1().ClusterManagers(),
				"testspokecluster",
				UpdateNucleusHubConditionFn(c.newCondition),
			)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}
			if updated != c.expectedUpdated {
				t.Errorf("expected %t, but %t", c.expectedUpdated, updated)
			}

			spokestatus, updated, err := UpdateNucleusSpokeStatus(
				context.TODO(),
				fakeClusterClient.NucleusV1().SpokeCores(),
				"testspokecluster",
				UpdateNucleusSpokeConditionFn(c.newCondition),
			)
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}
			if updated != c.expectedUpdated {
				t.Errorf("expected %t, but %t", c.expectedUpdated, updated)
			}

			for i := range c.expectedConditions {
				expected := c.expectedConditions[i]
				hubactual := hubstatus.Conditions[i]
				if expected.LastTransitionTime == (metav1.Time{}) {
					hubactual.LastTransitionTime = metav1.Time{}
				}
				if !equality.Semantic.DeepEqual(expected, hubactual) {
					t.Errorf(diff.ObjectDiff(expected, hubactual))
				}

				spokeactual := spokestatus.Conditions[i]
				if expected.LastTransitionTime == (metav1.Time{}) {
					spokeactual.LastTransitionTime = metav1.Time{}
				}
				if !equality.Semantic.DeepEqual(expected, spokeactual) {
					t.Errorf(diff.ObjectDiff(expected, spokeactual))
				}
			}
		})
	}
}

func newCondition(name, status, reason, message string, lastTransition *metav1.Time) nucleusapiv1.StatusCondition {
	ret := nucleusapiv1.StatusCondition{
		Type:    name,
		Status:  metav1.ConditionStatus(status),
		Reason:  reason,
		Message: message,
	}
	if lastTransition != nil {
		ret.LastTransitionTime = *lastTransition
	}
	return ret
}

func newValidatingWebhookConfiguration(name, svc, svcNameSpace string) *admissionv1.ValidatingWebhookConfiguration {
	return &admissionv1.ValidatingWebhookConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Webhooks: []admissionv1.ValidatingWebhook{
			{
				ClientConfig: admissionv1.WebhookClientConfig{
					Service: &admissionv1.ServiceReference{
						Name:      svc,
						Namespace: svcNameSpace,
					},
				},
			},
		},
	}
}

func newUnstructured(
	apiVersion, kind, namespace, name string, content map[string]interface{}) *unstructured.Unstructured {
	object := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": apiVersion,
			"kind":       kind,
			"metadata": map[string]interface{}{
				"namespace": namespace,
				"name":      name,
			},
		},
	}
	for key, val := range content {
		object.Object[key] = val
	}

	return object
}

func TestApplyValidatingWebhookConfiguration(t *testing.T) {
	testcase := []struct {
		name          string
		existing      []runtime.Object
		expected      *admissionv1.ValidatingWebhookConfiguration
		expectUpdated bool
	}{
		{
			name:          "Create a new configuration",
			expectUpdated: true,
			existing:      []runtime.Object{},
			expected:      newValidatingWebhookConfiguration("test", "svc1", "svc1"),
		},
		{
			name:          "update an existing configuration",
			expectUpdated: true,
			existing:      []runtime.Object{newValidatingWebhookConfiguration("test", "svc1", "svc1")},
			expected:      newValidatingWebhookConfiguration("test", "svc2", "svc2"),
		},
		{
			name:          "skip update",
			expectUpdated: false,
			existing:      []runtime.Object{newValidatingWebhookConfiguration("test", "svc1", "svc1")},
			expected:      newValidatingWebhookConfiguration("test", "svc1", "svc1"),
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			fakeKubeClient := fakekube.NewSimpleClientset(c.existing...)
			_, updated, err := ApplyValidatingWebhookConfiguration(fakeKubeClient.AdmissionregistrationV1(), c.expected)
			if err != nil {
				t.Errorf("Expected no error when applying: %v", err)
			}

			if updated != c.expectUpdated {
				t.Errorf("Expect update is %t, but got %t", c.expectUpdated, updated)
			}
		})
	}
}

func TestApplyDirectly(t *testing.T) {
	testcase := []struct {
		name           string
		applyFiles     map[string]runtime.Object
		applyFileNames []string
		expectErr      bool
	}{
		{
			name: "Apply webhooks & apiservice & secret",
			applyFiles: map[string]runtime.Object{
				"webhooks":   newUnstructured("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", "", "", map[string]interface{}{"webhooks": []interface{}{}}),
				"apiservice": newUnstructured("apiregistration.k8s.io/v1", "APIService", "", "", map[string]interface{}{"spec": map[string]interface{}{"service": map[string]string{"name": "svc1", "namespace": "svc1"}}}),
				"secret":     newUnstructured("v1", "Secret", "ns1", "n1", map[string]interface{}{"data": map[string]interface{}{"key1": []byte("key1")}}),
			},
			applyFileNames: []string{"webhooks", "apiservice", "secret"},
			expectErr:      false,
		},
		{
			name: "Apply unhandled object",
			applyFiles: map[string]runtime.Object{
				"kind1": newUnstructured("v1", "Kind1", "ns1", "n1", map[string]interface{}{"spec": map[string]interface{}{"key1": []byte("key1")}}),
			},
			applyFileNames: []string{"kind1"},
			expectErr:      true,
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			fakeKubeClient := fakekube.NewSimpleClientset()
			fakeResgistrationClient := fakeapiregistration.NewSimpleClientset()
			fakeExtensionClient := fakeapiextensions.NewSimpleClientset()
			results := ApplyDirectly(
				fakeKubeClient, fakeExtensionClient, fakeResgistrationClient.ApiregistrationV1(),
				eventstesting.NewTestingEventRecorder(t),
				func(name string) ([]byte, error) {
					if c.applyFiles[name] == nil {
						return nil, fmt.Errorf("Failed to find file")
					}

					return json.Marshal(c.applyFiles[name])
				},
				c.applyFileNames...,
			)
			aggregatedErr := []error{}
			for _, r := range results {
				if r.Error != nil {
					aggregatedErr = append(aggregatedErr, r.Error)
				}
			}

			if len(aggregatedErr) == 0 && c.expectErr {
				t.Errorf("Expect an apply error")
			}
			if len(aggregatedErr) != 0 && !c.expectErr {
				t.Errorf("Expect no apply error, %v", operatorhelpers.NewMultiLineAggregate(aggregatedErr))
			}
		})
	}
}

func TestDeleteStaticObject(t *testing.T) {
	applyFiles := map[string]runtime.Object{
		"webhooks":   newUnstructured("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", "", "", map[string]interface{}{"webhooks": []interface{}{}}),
		"apiservice": newUnstructured("apiregistration.k8s.io/v1", "APIService", "", "", map[string]interface{}{"spec": map[string]interface{}{"service": map[string]string{"name": "svc1", "namespace": "svc1"}}}),
		"secret":     newUnstructured("v1", "Secret", "ns1", "n1", map[string]interface{}{"data": map[string]interface{}{"key1": []byte("key1")}}),
		"crd":        newUnstructured("apiextensions.k8s.io/v1beta1", "CustomResourceDefinition", "", "", map[string]interface{}{}),
		"kind1":      newUnstructured("v1", "Kind1", "ns1", "n1", map[string]interface{}{"spec": map[string]interface{}{"key1": []byte("key1")}}),
	}
	testcase := []struct {
		name          string
		applyFileName string
		expectErr     bool
	}{
		{
			name:          "Delete webhooks",
			applyFileName: "webhooks",
			expectErr:     false,
		},
		{
			name:          "Delete apiservice",
			applyFileName: "apiservice",
			expectErr:     false,
		},
		{
			name:          "Delete secret",
			applyFileName: "secret",
			expectErr:     false,
		},
		{
			name:          "Delete crd",
			applyFileName: "crd",
			expectErr:     false,
		},
		{
			name:          "Delete unhandled object",
			applyFileName: "kind1",
			expectErr:     true,
		},
	}

	for _, c := range testcase {
		t.Run(c.name, func(t *testing.T) {
			fakeKubeClient := fakekube.NewSimpleClientset()
			fakeResgistrationClient := fakeapiregistration.NewSimpleClientset()
			fakeExtensionClient := fakeapiextensions.NewSimpleClientset()
			err := CleanUpStaticObject(
				context.TODO(),
				fakeKubeClient, fakeExtensionClient, fakeResgistrationClient.ApiregistrationV1(),
				func(name string) ([]byte, error) {
					if applyFiles[name] == nil {
						return nil, fmt.Errorf("Failed to find file")
					}

					return json.Marshal(applyFiles[name])
				},
				c.applyFileName,
			)

			if err == nil && c.expectErr {
				t.Errorf("Expect an apply error")
			}
			if err != nil && !c.expectErr {
				t.Errorf("Expect no apply error, %v", err)
			}
		})
	}
}
