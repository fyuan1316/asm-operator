package resource

import (
	"errors"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func PointerTrue() *bool {
	var t = true
	return &t
}
func PointerFalse() *bool {
	var t = false
	return &t
}

type Task struct {
	//子类override 接口
	Implementor model.Operation

	//资源mappings hook
	ResourceMappings map[metav1.TypeMeta]K8sResourceMapping

	//ResourceOptions  map[string]YamlResource
	//归属任务的资源集
	K8sResource map[string]SyncResource

	//任务层面资源是否随operator删除的设定
	KeepResourceAfterOperatorDeleted *bool

	//任务级别的values数据，一般对应到一个chart的values
	TemplateValues map[string]interface{}
}

func (m Task) GetStageName() string {
	panic("implement me")
}

func (m Task) Run(manage *model.OperatorManage) error {
	if m.Implementor.GetOperation() == model.Operations.Provision {
		return m.Sync(manage)
	} else if m.Implementor.GetOperation() == model.Operations.Deletion {
		return m.Delete(manage)
	} else {
		return errors.New("UnSupport type of ResourceTask")
	}
}

var _ model.ExecuteItem = Task{}

//type YamlResource struct {
//	model.Object
//	ChargeByOperator *bool
//}

type SyncResource struct {
	FileInfo
	model.Object
	//ChargeByOperator *bool
	Sync   func(client.Client, model.Object) error
	Delete func(client.Client, model.Object) error
}

func NewSyncResource(
	object model.Object,
	sync func(client.Client, model.Object) error,
	delete func(client.Client, model.Object) error) *SyncResource {
	res := &SyncResource{FileInfo: FileInfo{}}
	res.Object = object
	res.Sync = sync
	res.Delete = delete
	return res
}

func (m *SyncResource) FromMappings(resMapping *K8sResourceMapping) *SyncResource {
	return NewSyncResource(resMapping.Object, resMapping.Sync, resMapping.Deletion)
}

func (m *SyncResource) SetObject(o model.Object) {
	m.Object = o
}
func (m *SyncResource) SetOwnerRef() {
	t := true
	m.ChargeByOperator = &t
}
func (m *SyncResource) IsChargedByOwnerRef() *bool {
	return m.ChargeByOperator
}

/*
func (m *Task) LoadFile(filePath string, res *SyncResource, values map[string]interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return m.Load(string(bytes), res, values)
}
func (m *Task) Load(objectStr string, res *SyncResource, values map[string]interface{}) error {
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
*/
func (m *Task) Sync(om *model.OperatorManage) error {
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

func (m *Task) Delete(om *model.OperatorManage) error {
	for _, res := range m.K8sResource {
		if res.IsChargedByOwnerRef() != nil && *res.IsChargedByOwnerRef() ||
			m.KeepResourceAfterOperatorDeleted != nil && !*m.KeepResourceAfterOperatorDeleted {
			if err := res.Delete(om.K8sClient, res.Object); err != nil {
				return err
			}
		}
	}
	return nil
}
