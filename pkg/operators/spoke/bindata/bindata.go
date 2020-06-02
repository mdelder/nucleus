// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/spoke/spoke-registration-clusterrole.yaml
// manifests/spoke/spoke-registration-clusterrolebinding.yaml
// manifests/spoke/spoke-registration-deployment.yaml
// manifests/spoke/spoke-registration-role.yaml
// manifests/spoke/spoke-registration-rolebinding.yaml
// manifests/spoke/spoke-registration-serviceaccount.yaml
// manifests/spoke/spoke-work-clusterrole.yaml
// manifests/spoke/spoke-work-clusterrolebinding-addition.yaml
// manifests/spoke/spoke-work-clusterrolebinding.yaml
// manifests/spoke/spoke-work-deployment.yaml
// manifests/spoke/spoke-work-serviceaccount.yaml
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

var _manifestsSpokeSpokeRegistrationClusterroleYaml = []byte(`# Clusterrole for work agent in addition to admin clusterrole.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
rules:
# Allow agent to get/list/watch nodes.
- apiGroups: [""]
  resources: ["nodes", "configmaps", "secrets"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["authorization.k8s.io"]
  resources: ["subjectaccessreviews"]
  verbs: ["create"]
`)

func manifestsSpokeSpokeRegistrationClusterroleYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationClusterroleYaml, nil
}

func manifestsSpokeSpokeRegistrationClusterroleYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeRegistrationClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
subjects:
  - kind: ServiceAccount
    name: {{ .KlusterletName }}-registration-sa
    namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeRegistrationClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationClusterrolebindingYaml, nil
}

func manifestsSpokeSpokeRegistrationClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeRegistrationDeploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{ .KlusterletName }}-registration-agent
  namespace: {{ .KlusterletNamespace }}
  labels:
    app: klusterlet-registration-agent
spec:
  replicas: 3
  selector:
    matchLabels:
      app: klusterlet-registration-agent
  template:
    metadata:
      labels:
        app: klusterlet-registration-agent
    spec:
      serviceAccountName: {{ .KlusterletName }}-registration-sa
      containers:
      - name: spoke-agent
        image: {{ .RegistrationImage }}
        imagePullPolicy: IfNotPresent
        args:
          - "/registration"
          - "agent"
          - "--cluster-name={{ .ClusterName }}"
          - "--bootstrap-kubeconfig=/spoke/bootstrap/kubeconfig"
          - "--spoke-external-server-urls={{ .ExternalServerURL }}"
        volumeMounts:
        - name: bootstrap-secret
          mountPath: "/spoke/bootstrap"
          readOnly: true
        - name: hub-kubeconfig-secret
          mountPath: "/spoke/hub-kubeconfig"
          readOnly: true
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
      volumes:
      - name: bootstrap-secret
        secret:
          secretName: {{ .BootStrapKubeConfigSecret }}
      - name: hub-kubeconfig-secret
        secret:
          secretName: {{ .HubKubeConfigSecret }}
`)

func manifestsSpokeSpokeRegistrationDeploymentYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationDeploymentYaml, nil
}

func manifestsSpokeSpokeRegistrationDeploymentYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeRegistrationRoleYaml = []byte(`# Role for registration agent.
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
  namespace: {{ .KlusterletNamespace }}
rules:
- apiGroups: [""]
  resources: ["configmaps", "secrets"]
  verbs: ["get", "list", "watch", "create", "delete", "update", "patch"]
- apiGroups: ["", "events.k8s.io"]
  resources: ["events"]
  verbs: ["create", "patch", "update"]
`)

func manifestsSpokeSpokeRegistrationRoleYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationRoleYaml, nil
}

func manifestsSpokeSpokeRegistrationRoleYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeRegistrationRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
  namespace: {{ .KlusterletNamespace }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: system:open-cluster-management:{{ .KlusterletName }}-registration-agent
subjects:
  - kind: ServiceAccount
    name: {{ .KlusterletName }}-registration-sa
    namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeRegistrationRolebindingYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationRolebindingYaml, nil
}

func manifestsSpokeSpokeRegistrationRolebindingYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeRegistrationServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .KlusterletName }}-registration-sa
  namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeRegistrationServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeRegistrationServiceaccountYaml, nil
}

func manifestsSpokeSpokeRegistrationServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeRegistrationServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-registration-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeWorkClusterroleYaml = []byte(`# Clusterrole for work agent in addition to admin clusterrole.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-work-agent
rules:
# Allow agent to get/list/watch/create/delete crds.
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["get", "list", "watch", "create", "delete", "update"]
# Allow agent to create/delete namespaces, get/list are contained in admin role already
- apiGroups: [""]
  resources: ["namespaces"]
  verbs: ["create", "delete"]
# Allow agent to manage role/rolebinding/clusterrole/clusterrolebinding
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterrolebindings", "rolebindings"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
- apiGroups: ["rbac.authorization.k8s.io"]
  resources: ["clusterroles", "roles"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete", "escalate", "bind"]
# Allow agent to create sar
- apiGroups: ["authorization.k8s.io"]
  resources: ["subjectaccessreviews"]
  verbs: ["create"]
# Allow agent to create events
- apiGroups: ["", "events.k8s.io"]
  resources: ["events"]
  verbs: ["create", "patch", "update"]
`)

func manifestsSpokeSpokeWorkClusterroleYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeWorkClusterroleYaml, nil
}

func manifestsSpokeSpokeWorkClusterroleYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeWorkClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-work-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeWorkClusterrolebindingAdditionYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-work-agent-addition
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:open-cluster-management:{{ .KlusterletName }}-work-agent
subjects:
  - kind: ServiceAccount
    name: {{ .KlusterletName }}-work-sa
    namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeWorkClusterrolebindingAdditionYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeWorkClusterrolebindingAdditionYaml, nil
}

func manifestsSpokeSpokeWorkClusterrolebindingAdditionYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeWorkClusterrolebindingAdditionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-work-clusterrolebinding-addition.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeWorkClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system:open-cluster-management:{{ .KlusterletName }}-work-agent
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  # We deploy a controller that could work with permission lower than cluster-admin, the tradeoff is
  # responsivity because list/watch cannot be maintained over too many namespaces.
  name: admin
subjects:
  - kind: ServiceAccount
    name: {{ .KlusterletName }}-work-sa
    namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeWorkClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeWorkClusterrolebindingYaml, nil
}

func manifestsSpokeSpokeWorkClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeWorkClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-work-clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeWorkDeploymentYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{ .KlusterletName }}-work-agent
  namespace: {{ .KlusterletNamespace }}
  labels:
    app: klusterlet-manifestwork-agent
spec:
  replicas: 3
  selector:
    matchLabels:
      app: klusterlet-manifestwork-agent
  template:
    metadata:
      labels:
        app: klusterlet-manifestwork-agent
    spec:
      serviceAccountName: {{ .KlusterletName }}-work-sa
      containers:
      - name: spoke-agent
        image: {{ .WorkImage }}
        imagePullPolicy: IfNotPresent
        args:
          - "/work"
          - "agent"
          - "--spoke-cluster-name={{ .ClusterName }}"
          - "--hub-kubeconfig=/spoke/hub-kubeconfig/kubeconfig"
        volumeMounts:
        - name: hub-kubeconfig-secret
          mountPath: "/spoke/hub-kubeconfig"
          readOnly: true
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
      volumes:
      - name: hub-kubeconfig-secret
        secret:
          secretName: {{ .HubKubeConfigSecret }}
`)

func manifestsSpokeSpokeWorkDeploymentYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeWorkDeploymentYaml, nil
}

func manifestsSpokeSpokeWorkDeploymentYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeWorkDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-work-deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsSpokeSpokeWorkServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .KlusterletName }}-work-sa
  namespace: {{ .KlusterletNamespace }}
`)

func manifestsSpokeSpokeWorkServiceaccountYamlBytes() ([]byte, error) {
	return _manifestsSpokeSpokeWorkServiceaccountYaml, nil
}

func manifestsSpokeSpokeWorkServiceaccountYaml() (*asset, error) {
	bytes, err := manifestsSpokeSpokeWorkServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/spoke/spoke-work-serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"manifests/spoke/spoke-registration-clusterrole.yaml":         manifestsSpokeSpokeRegistrationClusterroleYaml,
	"manifests/spoke/spoke-registration-clusterrolebinding.yaml":  manifestsSpokeSpokeRegistrationClusterrolebindingYaml,
	"manifests/spoke/spoke-registration-deployment.yaml":          manifestsSpokeSpokeRegistrationDeploymentYaml,
	"manifests/spoke/spoke-registration-role.yaml":                manifestsSpokeSpokeRegistrationRoleYaml,
	"manifests/spoke/spoke-registration-rolebinding.yaml":         manifestsSpokeSpokeRegistrationRolebindingYaml,
	"manifests/spoke/spoke-registration-serviceaccount.yaml":      manifestsSpokeSpokeRegistrationServiceaccountYaml,
	"manifests/spoke/spoke-work-clusterrole.yaml":                 manifestsSpokeSpokeWorkClusterroleYaml,
	"manifests/spoke/spoke-work-clusterrolebinding-addition.yaml": manifestsSpokeSpokeWorkClusterrolebindingAdditionYaml,
	"manifests/spoke/spoke-work-clusterrolebinding.yaml":          manifestsSpokeSpokeWorkClusterrolebindingYaml,
	"manifests/spoke/spoke-work-deployment.yaml":                  manifestsSpokeSpokeWorkDeploymentYaml,
	"manifests/spoke/spoke-work-serviceaccount.yaml":              manifestsSpokeSpokeWorkServiceaccountYaml,
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
		"spoke": {nil, map[string]*bintree{
			"spoke-registration-clusterrole.yaml":         {manifestsSpokeSpokeRegistrationClusterroleYaml, map[string]*bintree{}},
			"spoke-registration-clusterrolebinding.yaml":  {manifestsSpokeSpokeRegistrationClusterrolebindingYaml, map[string]*bintree{}},
			"spoke-registration-deployment.yaml":          {manifestsSpokeSpokeRegistrationDeploymentYaml, map[string]*bintree{}},
			"spoke-registration-role.yaml":                {manifestsSpokeSpokeRegistrationRoleYaml, map[string]*bintree{}},
			"spoke-registration-rolebinding.yaml":         {manifestsSpokeSpokeRegistrationRolebindingYaml, map[string]*bintree{}},
			"spoke-registration-serviceaccount.yaml":      {manifestsSpokeSpokeRegistrationServiceaccountYaml, map[string]*bintree{}},
			"spoke-work-clusterrole.yaml":                 {manifestsSpokeSpokeWorkClusterroleYaml, map[string]*bintree{}},
			"spoke-work-clusterrolebinding-addition.yaml": {manifestsSpokeSpokeWorkClusterrolebindingAdditionYaml, map[string]*bintree{}},
			"spoke-work-clusterrolebinding.yaml":          {manifestsSpokeSpokeWorkClusterrolebindingYaml, map[string]*bintree{}},
			"spoke-work-deployment.yaml":                  {manifestsSpokeSpokeWorkDeploymentYaml, map[string]*bintree{}},
			"spoke-work-serviceaccount.yaml":              {manifestsSpokeSpokeWorkServiceaccountYaml, map[string]*bintree{}},
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
