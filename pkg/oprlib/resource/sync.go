package resource

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"io/ioutil"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"
)

type SyncManager struct {
	K8sResource map[string]SyncResource
}

type SyncResource struct {
	manage.Object
	setOwnerRef bool
	Sync        func(client.Client, manage.Object) error
}

func (m *SyncResource) SetOwnerRef() {
	m.setOwnerRef = true
}
func (m *SyncResource) IsChargedByOwnerRef() bool {
	return m.setOwnerRef
}
func (m *SyncManager) LoadFile(filePath string, res SyncResource) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return m.Load(string(bytes), res)
}
func (m *SyncManager) Load(objectStr string, res SyncResource) error {
	var err error
	object := res.Object
	err = yaml.Unmarshal([]byte(objectStr), object)
	if err != nil {
		return err
	}
	objKey := fmt.Sprintf("%s-%s-%s", object.GetObjectKind().GroupVersionKind().Kind,
		object.GetNamespace(),
		object.GetName(),
	)
	if m.K8sResource == nil {
		m.K8sResource = make(map[string]SyncResource)
	}
	m.K8sResource[objKey] = res
	return err
}

func (m *SyncManager) Sync(om *manage.OperatorManage) error {
	for k, res := range m.K8sResource {
		res11 := m.K8sResource[k]
		fmt.Println(res11.IsChargedByOwnerRef())
		if res.IsChargedByOwnerRef() {
			err := controllerutil.SetControllerReference(om.Object, res.Object, om.Scheme)
			if err != nil {
				return err
			}
		}
		if err := res.Sync(om.K8sClient, res.Object); err != nil {
			return err
		}
	}
	return nil
}
