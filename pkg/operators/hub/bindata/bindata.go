// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/hub/0000_00_clusters.open-cluster-management.io_spokeclusters.crd.yaml
// manifests/hub/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml
// manifests/hub/0000_00_work.open-cluster-management.io_manifestworks.crd.yaml
// manifests/hub/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml
// manifests/hub/hub-clusterrolebinding.yaml
// manifests/hub/hub-namespace.yaml
// manifests/hub/hub-registration-clusterrole.yaml
// manifests/hub/hub-registration-clusterrolebinding.yaml
// manifests/hub/hub-registration-deployment.yaml
// manifests/hub/hub-registration-serviceaccount.yaml
// manifests/hub/hub-registration-webhook-apiservice.yaml
// manifests/hub/hub-registration-webhook-clusterrole.yaml
// manifests/hub/hub-registration-webhook-clusterrolebinding.yaml
// manifests/hub/hub-registration-webhook-deployment.yaml
// manifests/hub/hub-registration-webhook-secret.yaml
// manifests/hub/hub-registration-webhook-service.yaml
// manifests/hub/hub-registration-webhook-serviceaccount.yaml
// manifests/hub/hub-registration-webhook-validatingconfiguration.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: spokeclusters.cluster.open-cluster-management.io
spec:
  group: cluster.open-cluster-management.io
  names:
    kind: SpokeCluster
    listKind: SpokeClusterList
    plural: spokeclusters
    singular: spokecluster
  scope: "Cluster"
  subresources:
    status: {}
  preserveUnknownFields: false
  validation:
    openAPIV3Schema:
      description: "SpokeCluster represents the desired state and current status of
        spoke cluster. SpokeCluster is a cluster scoped resource. The name is the
        cluster UID. \n The cluster join process follows a double opt-in process:
        \n 1. agent on spoke cluster creates CSR on hub with cluster UID and agent
        name. 2. agent on spoke cluster creates spokecluster on hub. 3. cluster admin
        on hub approves the CSR for the spoke's cluster UID and agent name. 4. cluster
        admin set spec.acceptSpokeCluster of spokecluster to true. 5. cluster admin
        on spoke creates credential of kubeconfig to spoke. \n Once the hub creates
        the cluster namespace, the spoke agent pushes the credential to the hub to
        use against the spoke's kube-apiserver."
      type: object
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: Spec represents a desired configuration for the agent on the
            spoke cluster.
          type: object
          properties:
            hubAcceptsClient:
              description: AcceptSpokeCluster reprsents that hub accepts the join
                of spoke agent. Its default value is false, and can only be set true
                when the user on hub has an RBAC rule to UPDATE on the virtual subresource
                of spokeclusters/accept. When the vaule is set true, a namespace whose
                name is same as the name of SpokeCluster is created on hub representing
                the spoke cluster, also role/rolebinding is created on the namespace
                to grant the permision of access from agent on spoke. When the value
                is set false, the namespace representing the spoke cluster is deleted.
              type: boolean
            leaseDurationSeconds:
              description: LeaseDurationSeconds is used to coordinate the lease update
                time of spoke agents. If its value is zero, the spoke agent will update
                its lease per 60s by default
              type: integer
              format: int32
            spokeClientConfigs:
              description: SpokeClientConfigs represents a list of the apiserver address
                of the spoke cluster. If it is empty, spoke cluster has no accessible
                address to be visited from hub.
              type: array
              items:
                description: ClientConfig represents the apiserver address of the
                  spoke cluster. TODO include credential to connect to spoke cluster
                  kube-apiserver
                type: object
                properties:
                  caBundle:
                    description: CABundle is the ca bundle to connect to apiserver
                      of the spoke cluster. System certs are used if it is not set.
                    type: string
                    format: byte
                  url:
                    description: URL is the url of apiserver endpoint of the spoke
                      cluster.
                    type: string
        status:
          description: Status represents the current status of joined spoke cluster
          type: object
          properties:
            allocatable:
              description: Allocatable represents the total allocatable resources
                on the spoke cluster.
              type: object
              additionalProperties:
                type: string
            capacity:
              description: Capacity represents the total resource capacity from all
                nodeStatuses on the spoke cluster.
              type: object
              additionalProperties:
                type: string
            conditions:
              description: Conditions contains the different condition statuses for
                this spoke cluster.
              type: array
              items:
                description: StatusCondition contains condition information for a
                  spoke cluster.
                type: object
                properties:
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      changed from one status to another.
                    type: string
                    format: date-time
                  message:
                    description: Message is a human-readable message indicating details
                      about the last status change.
                    type: string
                  reason:
                    description: Reason is a (brief) reason for the condition's last
                      status change.
                    type: string
                  status:
                    description: Status is the status of the condition. One of True,
                      False, Unknown.
                    type: string
                  type:
                    description: Type is the type of the cluster condition.
                    type: string
            version:
              description: Version represents the kubernetes version of the spoke
                cluster.
              type: object
              properties:
                kubernetes:
                  description: Kubernetes is the kubernetes version of spoke cluster
                  type: string
  version: v1
  versions:
  - name: v1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYamlBytes() ([]byte, error) {
	return _manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYaml, nil
}

func manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYaml() (*asset, error) {
	bytes, err := manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/0000_00_clusters.open-cluster-management.io_spokeclusters.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: klusterlets.operator.open-cluster-management.io
spec:
  group: operator.open-cluster-management.io
  names:
    kind: Klusterlet
    listKind: KlusterletList
    plural: klusterlets
    singular: klusterlet
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: Klusterlet represents controllers on the managed cluster. When
        configured, the Klusterlet requires a secret named of bootstrap-hub-kubeconfig
        in the same namespace to allow API requests to the hub for the registration
        protocol.
      type: object
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: Spec represents the desired deployment configuration of klusterlet
            agent.
          type: object
          properties:
            clusterName:
              description: ClusterName is the name of the spoke cluster to be created
                on hub. The spoke agent generates a random name if it is not set,
                or discovers the appropriate cluster name on openshift.
              type: string
            externalServerURLs:
              description: ExternalServerURLs represents the a list of apiserver urls
                and ca bundles that is accessible externally If it is set empty, spoke
                cluster has no externally accessible url that hub cluster can visit.
              type: array
              items:
                description: ServerURL represents the apiserver url and ca bundle
                  that is accessible externally
                type: object
                properties:
                  caBundle:
                    description: CABundle is the ca bundle to connect to apiserver
                      of the spoke cluster. System certs are used if it is not set.
                    type: string
                    format: byte
                  url:
                    description: URL is the url of apiserver endpoint of the spoke
                      cluster.
                    type: string
            namespace:
              description: Namespace is the namespace to deploy the agent. The namespace
                must have a prefix of "open-cluster-management-", and if it is not
                set, the namespace of "open-cluster-management-spoke" is used to deploy
                agent.
              type: string
            registrationImagePullSpec:
              description: RegistrationImagePullSpec represents the desired image
                configuration of registration agent.
              type: string
            workImagePullSpec:
              description: WorkImagePullSpec represents the desired image configuration
                of work agent.
              type: string
        status:
          description: Status represents the current status of klusterlet agent.
          type: object
          properties:
            conditions:
              description: 'Conditions contain the different condition statuses for
                this spokecore. Valid condition types are: Applied: components in
                spoke is applied. Available: components in spoke are available and
                ready to serve. Progressing: components in spoke are in a transitioning
                state. Degraded: components in spoke do not match the desired configuration
                and only provide degraded service.'
              type: array
              items:
                description: StatusCondition contains condition information.
                type: object
                properties:
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      changed from one status to another.
                    type: string
                    format: date-time
                  message:
                    description: Message is a human-readable message indicating details
                      about the last status change.
                    type: string
                  reason:
                    description: Reason is a (brief) reason for the condition's last
                      status change.
                    type: string
                  status:
                    description: Status is the status of the condition. One of True,
                      False, Unknown.
                    type: string
                  type:
                    description: Type is the type of the cluster condition.
                    type: string
  version: v1
  versions:
  - name: v1
    served: true
    storage: true
  preserveUnknownFields: false
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYamlBytes() ([]byte, error) {
	return _manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYaml, nil
}

func manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYaml() (*asset, error) {
	bytes, err := manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: manifestworks.work.open-cluster-management.io
spec:
  group: work.open-cluster-management.io
  names:
    kind: ManifestWork
    listKind: ManifestWorkList
    plural: manifestworks
    singular: manifestwork
  scope: "Namespaced"
  preserveUnknownFields: false
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: ManifestWork represents a manifests workload that hub wants to
        deploy on the spoke cluster. A manifest workload is defined as a set of kubernetes
        resources. ManifestWork must be created in the cluster namespace on the hub,
        so that agent on the corresponding spoke cluster can access this resource
        and deploy on the spoke cluster.
      type: object
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: Spec represents a desired configuration of work to be deployed
            on the spoke cluster.
          type: object
          properties:
            workload:
              description: Workload represents the manifest workload to be deployed
                on spoke cluster
              type: object
              properties:
                manifests:
                  description: Manifests represents a list of kuberenetes resources
                    to be deployed on the spoke cluster.
                  type: array
                  items:
                    description: Manifest represents a resource to be deployed on
                      spoke cluster
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                    x-kubernetes-embedded-resource: true
        status:
          description: Status represents the current status of work
          type: object
          properties:
            appliedResources:
              description: AppliedResources represents a list of resources defined
                within the manifestwork that are applied. Only resources with valid
                GroupVersionResource, namespace, and name are suitable. An item in
                this slice is deleted when there is no mapped manifest in manifestwork.Spec
                or by finalizer. The resource relating to the item will also be removed
                from spoke cluster. The deleted resource may still be present until
                the finalizers for that resource are finished. However, the resource
                will not be undeleted, so it can be removed from this list and eventual
                consistency is preserved.
              type: array
              items:
                description: AppliedManifestResourceMeta represents the gvr, name
                  and namespace of a resource. Since these resources have been created,
                  they must have valid group, version, resource, namespace, and name.
                type: object
                properties:
                  group:
                    description: Group is the API Group of the kubernetes resource
                    type: string
                  name:
                    description: Name is the name of the kubernetes resource
                    type: string
                  namespace:
                    description: Name is the namespace of the kubernetes resource,
                      empty string indicates it is a cluster scoped resource.
                    type: string
                  resource:
                    description: Resource is the resource name of the kubernetes resource
                    type: string
                  version:
                    description: Version is the version of the kubernetes resource
                    type: string
            conditions:
              description: 'Conditions contains the different condition statuses for
                this work. Valid condition types are: 1. Applied represents workload
                in ManifestWork is applied successfully on spoke cluster. 2. Progressing
                represents workload in ManifestWork is being applied on spoke cluster.
                3. Available represents workload in ManifestWork exists on the spoke
                cluster. 4. Degraded represents the current state of workload does
                not match the desired state for a certain period.'
              type: array
              items:
                description: StatusCondition contains condition information for a
                  spoke work.
                type: object
                properties:
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      changed from one status to another.
                    type: string
                    format: date-time
                  message:
                    description: Message is a human-readable message indicating details
                      about the last status change.
                    type: string
                  reason:
                    description: Reason is a (brief) reason for the condition's last
                      status change.
                    type: string
                  status:
                    description: Status is the status of the condition. One of True,
                      False, Unknown.
                    type: string
                  type:
                    description: Type is the type of the spoke work condition.
                    type: string
            resourceStatus:
              description: ResourceStatus represents the status of each resource in
                manifestwork deployed on spoke cluster. The agent on spoke cluster
                syncs the condition from spoke to the hub.
              type: object
              properties:
                manifests:
                  description: 'Manifests represents the condition of manifests deployed
                    on spoke cluster. Valid condition types are: 1. Progressing represents
                    the resource is being applied on spoke cluster. 2. Applied represents
                    the resource is applied successfully on spoke cluster. 3. Available
                    represents the resource exists on the spoke cluster. 4. Degraded
                    represents the current state of resource does not match the desired
                    state for a certain period.'
                  type: array
                  items:
                    description: ManifestCondition represents the conditions of the
                      resources deployed on spoke cluster
                    type: object
                    properties:
                      conditions:
                        description: Conditions represents the conditions of this
                          resource on spoke cluster
                        type: array
                        items:
                          description: StatusCondition contains condition information
                            for a spoke work.
                          type: object
                          properties:
                            lastTransitionTime:
                              description: LastTransitionTime is the last time the
                                condition changed from one status to another.
                              type: string
                              format: date-time
                            message:
                              description: Message is a human-readable message indicating
                                details about the last status change.
                              type: string
                            reason:
                              description: Reason is a (brief) reason for the condition's
                                last status change.
                              type: string
                            status:
                              description: Status is the status of the condition.
                                One of True, False, Unknown.
                              type: string
                            type:
                              description: Type is the type of the spoke work condition.
                              type: string
                      resourceMeta:
                        description: ResourceMeta represents the gvk, name and namespace
                          of a resoure
                        type: object
                        properties:
                          group:
                            description: Group is the API Group of the kubernetes
                              resource
                            type: string
                          kind:
                            description: Kind is the kind of the kubernetes resource
                            type: string
                          name:
                            description: Name is the name of the kubernetes resource
                            type: string
                          namespace:
                            description: Name is the namespace of the kubernetes resource
                            type: string
                          ordinal:
                            description: Ordinal represents the index of the manifest
                              on spec
                            type: integer
                            format: int32
                          resource:
                            description: Resource is the resource name of the kubernetes
                              resource
                            type: string
                          version:
                            description: Version is the version of the kubernetes
                              resource
                            type: string
  version: v1
  versions:
  - name: v1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYamlBytes() ([]byte, error) {
	return _manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYaml, nil
}

func manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYaml() (*asset, error) {
	bytes, err := manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/0000_00_work.open-cluster-management.io_manifestworks.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: clustermanagers.operator.open-cluster-management.io
spec:
  group: operator.open-cluster-management.io
  names:
    kind: ClusterManager
    listKind: ClusterManagerList
    plural: clustermanagers
    singular: clustermanager
  scope: Cluster
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      description: ClusterManager configures the controllers on the hub that govern
        registration and work distribution for attached klusterlets. ClusterManager
        will be only deployed in open-cluster-management-hub namespace.
      type: object
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: Spec represents a desired deployment configuration of controllers
            that govern registration and work distribution for attached klusterlets.
          type: object
          properties:
            registrationImagePullSpec:
              description: RegistrationImagePullSpec represents the desired image
                of registration controller installed on hub.
              type: string
        status:
          description: Status represents the current status of controllers that govern
            the lifecycle of managed clusters.
          type: object
          properties:
            conditions:
              description: 'Conditions contain the different condition statuses for
                this ClusterManager. Valid condition types are: Applied: components
                in hub are applied. Available: components in hub are available and
                ready to serve. Progressing: components in hub are in a transitioning
                state. Degraded: components in hub do not match the desired configuration
                and only provide degraded service.'
              type: array
              items:
                description: StatusCondition contains condition information.
                type: object
                properties:
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      changed from one status to another.
                    type: string
                    format: date-time
                  message:
                    description: Message is a human-readable message indicating details
                      about the last status change.
                    type: string
                  reason:
                    description: Reason is a (brief) reason for the condition's last
                      status change.
                    type: string
                  status:
                    description: Status is the status of the condition. One of True,
                      False, Unknown.
                    type: string
                  type:
                    description: Type is the type of the cluster condition.
                    type: string
  version: v1
  versions:
  - name: v1
    served: true
    storage: true
  preserveUnknownFields: false
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYamlBytes() ([]byte, error) {
	return _manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYaml, nil
}

func manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYaml() (*asset, error) {
	bytes, err := manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .ClusterManagerName }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:open-cluster-management:{{ .ClusterManagerName }}
subjects:
- kind: ServiceAccount
  namespace: {{ .ClusterManagerNamespace }}
  name: {{ .ClusterManagerName }}-sa
`)

func manifestsHubHubClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsHubHubClusterrolebindingYaml, nil
}

func manifestsHubHubClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsHubHubClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: {{ .ClusterManagerNamespace }}
`)

func manifestsHubHubNamespaceYamlBytes() ([]byte, error) {
	return _manifestsHubHubNamespaceYaml, nil
}

func manifestsHubHubNamespaceYaml() (*asset, error) {
	bytes, err := manifestsHubHubNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
<<<<<<< 249448baea60bb7d36bf7ea1c432e54623de372d:manifests/hub/hub-registration-clusterrole.yaml
  name: system:open-cluster-management:{{ .HubCoreName }}-registration-controller
=======
  name: system:open-cluster-management:{{ .ClusterManagerName }}
>>>>>>> Adopt new API updates for operator:manifests/hub/hub-clusterrole.yaml
rules:
# Allow hub to monitor and update status of csr
- apiGroups: ["certificates.k8s.io"]
  resources: ["certificatesigningrequests"]
  verbs: ["create", "get", "list", "watch"]
- apiGroups: ["certificates.k8s.io"]
  resources: ["certificatesigningrequests/status", "certificatesigningrequests/approval"]
  verbs: ["update"]
# Allow hub to get/list/watch/create/delete namespace and service account
- apiGroups: [""]
  resources: ["namespaces", "serviceaccounts", "configmaps"]
  verbs: ["get", "list", "watch", "create", "delete", "update"]
- apiGroups: ["", "events.k8s.io"]
  resources: ["events"]
  verbs: ["create", "patch", "update"]
- apiGroups: ["authorization.k8s.io"]
  resources: ["subjectaccessreviews"]
  verbs: ["create"]
# Allow hub to manage clusterrole/clusterrolebinding/role/rolebinding
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterrolebindings", "rolebindings"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterroles", "roles"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete", "escalate", "bind"]
# Allow hub to manage spokeclusters
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["spokeclusters"]
  verbs: ["get", "list", "watch", "update", "patch"]
- apiGroups: ["cluster.open-cluster-management.io"]
  resources: ["spokeclusters/status"]
  verbs: ["update", "patch"]
`)

func manifestsHubHubRegistrationClusterroleYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationClusterroleYaml, nil
}

func manifestsHubHubRegistrationClusterroleYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .HubCoreName }}-registration-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:open-cluster-management:{{ .HubCoreName }}-registration-controller
subjects:
- kind: ServiceAccount
  namespace: {{ .HubCoreNamespace }}
  name: {{ .HubCoreName }}-registration-controller-sa
`)

func manifestsHubHubRegistrationClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationClusterrolebindingYaml, nil
}

func manifestsHubHubRegistrationClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationDeploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
<<<<<<< 249448baea60bb7d36bf7ea1c432e54623de372d:manifests/hub/hub-registration-deployment.yaml
  name: {{ .HubCoreName }}-registration-controller
  namespace: {{ .HubCoreNamespace }}
  labels:
    app: nucleushub-registration-controller
=======
  name: {{ .ClusterManagerName }}-controller
  namespace: {{ .ClusterManagerNamespace }}
  labels:
    app: clustermanager-controller
>>>>>>> Adopt new API updates for operator:manifests/hub/hub-deployment.yaml
spec:
  replicas: 3
  selector:
    matchLabels:
<<<<<<< 249448baea60bb7d36bf7ea1c432e54623de372d:manifests/hub/hub-registration-deployment.yaml
      app: nucleushub-registration-controller
  template:
    metadata:
      labels:
        app: nucleushub-registration-controller
    spec:
      serviceAccountName: {{ .HubCoreName }}-registration-controller-sa
=======
      app: clustermanager-controller
  template:
    metadata:
      labels:
        app: clustermanager-controller
    spec:
      serviceAccountName: {{ .ClusterManagerName }}-sa
>>>>>>> Adopt new API updates for operator:manifests/hub/hub-deployment.yaml
      containers:
      - name: hub-registration-controller
        image: {{ .RegistrationImage }}
        imagePullPolicy: IfNotPresent
        args:
          - "/registration"
          - "controller"
        livenessProbe:
          httpGet:
            path: /healthz
            scheme: HTTPS
            port: 8443
          initialDelaySeconds: 2
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /healthz
            scheme: HTTPS
            port: 8443
          initialDelaySeconds: 2
`)

func manifestsHubHubRegistrationDeploymentYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationDeploymentYaml, nil
}

func manifestsHubHubRegistrationDeploymentYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .HubCoreName }}-registration-controller-sa
  namespace: {{ .HubCoreNamespace }}
`)

func manifestsHubHubRegistrationServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationServiceaccountYaml, nil
}

func manifestsHubHubRegistrationServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookApiserviceYaml = []byte(`apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  name: v1.admission.cluster.open-cluster-management.io
spec:
  group: admission.cluster.open-cluster-management.io
  version: v1
  service:
    name: {{ .HubCoreWebhookRegistrationService }}
    namespace: {{ .HubCoreNamespace }}
  caBundle: {{ .RegistrationAPIServiceCABundle }}
  groupPriorityMinimum: 10000
  versionPriority: 20
`)

func manifestsHubHubRegistrationWebhookApiserviceYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookApiserviceYaml, nil
}

func manifestsHubHubRegistrationWebhookApiserviceYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:open-cluster-management:{{ .HubCoreName }}-registration-webhook
rules:
# Allow spokecluster admission to get/list/watch configmaps
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get", "list", "watch"]
# Allow spokecluster admission to create subjectaccessreviews
- apiGroups: ["authorization.k8s.io"]
  resources: ["subjectaccessreviews"]
  verbs: ["create"]
`)

func manifestsHubHubRegistrationWebhookClusterroleYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookClusterroleYaml, nil
}

func manifestsHubHubRegistrationWebhookClusterroleYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .HubCoreName }}-registration-webhook
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: {{ .HubCoreName }}-registration-webhook
subjects:
  - kind: ServiceAccount
    name: {{ .HubCoreName }}-registration-webhook-sa
    namespace: {{ .HubCoreNamespace }}
`)

func manifestsHubHubRegistrationWebhookClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookClusterrolebindingYaml, nil
}

func manifestsHubHubRegistrationWebhookClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .HubCoreName }}-registration-webhook
  namespace: {{ .HubCoreNamespace }}
  labels:
    app: {{ .HubCoreName }}-registration-webhook
spec:
  replicas: 3
  selector:
    matchLabels:
      app: {{ .HubCoreName }}-registration-webhook
  template:
    metadata:
      labels:
        app: {{ .HubCoreName }}-registration-webhook
    spec:
      serviceAccountName: {{ .HubCoreName }}-registration-webhook-sa
      containers:
      - name: {{ .HubCoreName }}-registration-webhook-sa
        image: {{ .RegistrationImage }}
        imagePullPolicy: IfNotPresent
        args:
          - "/registration"
          - "webhook"
          - "--secure-port=6443"
          - "--tls-cert-file=/serving-cert/tls.crt"
          - "--tls-private-key-file=/serving-cert/tls.key"
        livenessProbe:
          httpGet:
            path: /healthz
            scheme: HTTPS
            port: 6443
          initialDelaySeconds: 2
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /healthz
            scheme: HTTPS
            port: 6443
          initialDelaySeconds: 2
        volumeMounts:
        - name: webhook-secret
          mountPath: "/serving-cert"
          readOnly: true
      volumes:
      - name: webhook-secret
        secret:
          secretName: {{ .HubCoreWebhookSecret }}

`)

func manifestsHubHubRegistrationWebhookDeploymentYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookDeploymentYaml, nil
}

func manifestsHubHubRegistrationWebhookDeploymentYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookSecretYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  name: {{ .HubCoreWebhookSecret }}
  namespace: {{ .HubCoreNamespace }}
data:
  tls.crt: {{ .RegistrationServingCert }}
  tls.key: {{ .RegistrationServingKey }}
  ca.crt: {{ .RegistrationAPIServiceCABundle }}
type: Opaque
`)

func manifestsHubHubRegistrationWebhookSecretYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookSecretYaml, nil
}

func manifestsHubHubRegistrationWebhookSecretYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookSecretYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-secret.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookServiceYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  name: {{ .HubCoreWebhookRegistrationService }}
  namespace: {{ .HubCoreNamespace }}
spec:
  selector:
    app: {{ .HubCoreName }}-registration-webhook
  ports:
  - port: 443
    targetPort: 6443
`)

func manifestsHubHubRegistrationWebhookServiceYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookServiceYaml, nil
}

func manifestsHubHubRegistrationWebhookServiceYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
<<<<<<< 249448baea60bb7d36bf7ea1c432e54623de372d:manifests/hub/hub-registration-webhook-serviceaccount.yaml
  name: {{ .HubCoreName }}-registration-webhook-sa
  namespace: {{ .HubCoreNamespace }}
=======
  name: {{ .ClusterManagerName }}-sa
  namespace: {{ .ClusterManagerNamespace }}
>>>>>>> Adopt new API updates for operator:manifests/hub/hub-serviceaccount.yaml
`)

func manifestsHubHubRegistrationWebhookServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookServiceaccountYaml, nil
}

func manifestsHubHubRegistrationWebhookServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsHubHubRegistrationWebhookValidatingconfigurationYaml = []byte(`apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: spokeclustervalidators.admission.cluster.open-cluster-management.io
webhooks:
- name: spokeclustervalidators.admission.cluster.open-cluster-management.io
  failurePolicy: Fail
  clientConfig:
    service:
      # reach the webhook via the registered aggregated API
      namespace: default
      name: kubernetes
      path: /apis/admission.cluster.open-cluster-management.io/v1/spokeclustervalidators
  rules:
  - operations:
    - CREATE
    - UPDATE
    apiGroups:
    - cluster.open-cluster-management.io
    apiVersions:
    - "*"
    resources:
    - spokeclusters
  admissionReviewVersions: ["v1beta1"]
  sideEffects: None
  timeoutSeconds: 3
`)

func manifestsHubHubRegistrationWebhookValidatingconfigurationYamlBytes() ([]byte, error) {
	return _manifestsHubHubRegistrationWebhookValidatingconfigurationYaml, nil
}

func manifestsHubHubRegistrationWebhookValidatingconfigurationYaml() (*asset, error) {
	bytes, err := manifestsHubHubRegistrationWebhookValidatingconfigurationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/hub/hub-registration-webhook-validatingconfiguration.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"manifests/hub/0000_00_clusters.open-cluster-management.io_spokeclusters.crd.yaml":   manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYaml,
	"manifests/hub/0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml":     manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYaml,
	"manifests/hub/0000_00_work.open-cluster-management.io_manifestworks.crd.yaml":       manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYaml,
	"manifests/hub/0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml": manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYaml,
	"manifests/hub/hub-clusterrolebinding.yaml":                                          manifestsHubHubClusterrolebindingYaml,
	"manifests/hub/hub-namespace.yaml":                                                   manifestsHubHubNamespaceYaml,
	"manifests/hub/hub-registration-clusterrole.yaml":                                    manifestsHubHubRegistrationClusterroleYaml,
	"manifests/hub/hub-registration-clusterrolebinding.yaml":                             manifestsHubHubRegistrationClusterrolebindingYaml,
	"manifests/hub/hub-registration-deployment.yaml":                                     manifestsHubHubRegistrationDeploymentYaml,
	"manifests/hub/hub-registration-serviceaccount.yaml":                                 manifestsHubHubRegistrationServiceaccountYaml,
	"manifests/hub/hub-registration-webhook-apiservice.yaml":                             manifestsHubHubRegistrationWebhookApiserviceYaml,
	"manifests/hub/hub-registration-webhook-clusterrole.yaml":                            manifestsHubHubRegistrationWebhookClusterroleYaml,
	"manifests/hub/hub-registration-webhook-clusterrolebinding.yaml":                     manifestsHubHubRegistrationWebhookClusterrolebindingYaml,
	"manifests/hub/hub-registration-webhook-deployment.yaml":                             manifestsHubHubRegistrationWebhookDeploymentYaml,
	"manifests/hub/hub-registration-webhook-secret.yaml":                                 manifestsHubHubRegistrationWebhookSecretYaml,
	"manifests/hub/hub-registration-webhook-service.yaml":                                manifestsHubHubRegistrationWebhookServiceYaml,
	"manifests/hub/hub-registration-webhook-serviceaccount.yaml":                         manifestsHubHubRegistrationWebhookServiceaccountYaml,
	"manifests/hub/hub-registration-webhook-validatingconfiguration.yaml":                manifestsHubHubRegistrationWebhookValidatingconfigurationYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"manifests": {nil, map[string]*bintree{
		"hub": {nil, map[string]*bintree{
			"0000_00_clusters.open-cluster-management.io_spokeclusters.crd.yaml":   {manifestsHub0000_00_clustersOpenClusterManagementIo_spokeclustersCrdYaml, map[string]*bintree{}},
			"0000_00_operator.open-cluster-management.io_klusterlets.crd.yaml":     {manifestsHub0000_00_operatorOpenClusterManagementIo_klusterletsCrdYaml, map[string]*bintree{}},
			"0000_00_work.open-cluster-management.io_manifestworks.crd.yaml":       {manifestsHub0000_00_workOpenClusterManagementIo_manifestworksCrdYaml, map[string]*bintree{}},
			"0000_01_operator.open-cluster-management.io_clustermanagers.crd.yaml": {manifestsHub0000_01_operatorOpenClusterManagementIo_clustermanagersCrdYaml, map[string]*bintree{}},
			"hub-clusterrolebinding.yaml":                                          {manifestsHubHubClusterrolebindingYaml, map[string]*bintree{}},
			"hub-namespace.yaml":                                                   {manifestsHubHubNamespaceYaml, map[string]*bintree{}},
			"hub-registration-clusterrole.yaml":                                    {manifestsHubHubRegistrationClusterroleYaml, map[string]*bintree{}},
			"hub-registration-clusterrolebinding.yaml":                             {manifestsHubHubRegistrationClusterrolebindingYaml, map[string]*bintree{}},
			"hub-registration-deployment.yaml":                                     {manifestsHubHubRegistrationDeploymentYaml, map[string]*bintree{}},
			"hub-registration-serviceaccount.yaml":                                 {manifestsHubHubRegistrationServiceaccountYaml, map[string]*bintree{}},
			"hub-registration-webhook-apiservice.yaml":                             {manifestsHubHubRegistrationWebhookApiserviceYaml, map[string]*bintree{}},
			"hub-registration-webhook-clusterrole.yaml":                            {manifestsHubHubRegistrationWebhookClusterroleYaml, map[string]*bintree{}},
			"hub-registration-webhook-clusterrolebinding.yaml":                     {manifestsHubHubRegistrationWebhookClusterrolebindingYaml, map[string]*bintree{}},
			"hub-registration-webhook-deployment.yaml":                             {manifestsHubHubRegistrationWebhookDeploymentYaml, map[string]*bintree{}},
			"hub-registration-webhook-secret.yaml":                                 {manifestsHubHubRegistrationWebhookSecretYaml, map[string]*bintree{}},
			"hub-registration-webhook-service.yaml":                                {manifestsHubHubRegistrationWebhookServiceYaml, map[string]*bintree{}},
			"hub-registration-webhook-serviceaccount.yaml":                         {manifestsHubHubRegistrationWebhookServiceaccountYaml, map[string]*bintree{}},
			"hub-registration-webhook-validatingconfiguration.yaml":                {manifestsHubHubRegistrationWebhookValidatingconfigurationYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
