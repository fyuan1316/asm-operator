package resource

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

func PointerTrue() *bool {
	var t = true
	return &t
}
func PointerFalse() *bool {
	var t = false
	return &t
}

type SyncManager struct {
	K8sResource map[string]SyncResource
	//top setting
	ChargeByOperator *bool
	////key:gvk
	//DynamicClients map[string]dynamic.NamespaceableResourceInterface
}

type SyncResource struct {
	model.Object
	SetOwnerReference *bool
	Sync              func(client.Client, model.Object) error
	Delete            func(client.Client, model.Object) error
}

func (m *SyncResource) SetObject(o model.Object) {
	m.Object = o
}
func (m *SyncResource) SetOwnerRef() {
	t := true
	m.SetOwnerReference = &t
}
func (m *SyncResource) IsChargedByOwnerRef() *bool {
	return m.SetOwnerReference
}
func (m *SyncManager) LoadFile(filePath string, res *SyncResource, values map[string]interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return m.Load(string(bytes), res, values)
}
func (m *SyncManager) Load(objectStr string, res *SyncResource, values map[string]interface{}) error {
	var err error
	//meta := manage.TypeObjectMeta{}
	// render values TODO mergeDefaults
	renderedObjectStr, err := Parse(objectStr, values)
	if err != nil {
		return err
	}
	unStruct := unstructured.Unstructured{}
	err = yaml.Unmarshal([]byte(renderedObjectStr), &unStruct)

	if err != nil {
		return err
	}
	if &unStruct == nil {
		fmt.Println()
	}
	if err = updateSyncResource(unStruct, res); err != nil {
		return err
	}
	object := res.Object

	err = yaml.Unmarshal([]byte(renderedObjectStr), object)
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
	m.K8sResource[objKey] = *res
	return err
}
func (m *SyncManager) Sync(om *model.OperatorManage) error {
	for _, res := range m.K8sResource {
		//资源参数优先
		/*
			if res.IsChargedByOwnerRef() != nil && *res.IsChargedByOwnerRef() ||
				m.ChargeByOperator != nil && *m.ChargeByOperator {
				err := controllerutil.SetControllerReference(om.Object, res.Object, om.Scheme)
				if err != nil {
					return err
				}
			}
		*/
		if err := res.Sync(om.K8sClient, res.Object); err != nil {
			return err
		}
	}
	return nil
}

func (m *SyncManager) Delete(om *model.OperatorManage) error {
	for _, res := range m.K8sResource {
		if res.IsChargedByOwnerRef() != nil && *res.IsChargedByOwnerRef() ||
			m.ChargeByOperator != nil && *m.ChargeByOperator {
			if err := res.Delete(om.K8sClient, res.Object); err != nil {
				return err
			}
		}
	}
	return nil
}
