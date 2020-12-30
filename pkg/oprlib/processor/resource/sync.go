package resource

import (
	"errors"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PointerTrue() *bool {
	var t = true
	return &t
}
func PointerFalse() *bool {
	var t = false
	return &t
}

var _ model.ExecuteItem = Task{}

type Task struct {
	//子类override 接口
	implementor model.OverrideOperation

	//资源mappings hook
	ResourceMappings map[metav1.TypeMeta]*K8sResourceMapping

	//ResourceOptions  map[string]YamlResource
	//归属任务的资源集
	K8sResource map[string]SyncResource

	//任务层面资源是否随operator删除的设定
	KeepResourceAfterOperatorDeleted *bool

	//任务级别的values数据，一般对应到一个chart的values
	TemplateValues map[string]interface{}
}

func (m *Task) Override(operation model.OverrideOperation) {
	m.implementor = operation
}
func (m Task) GetStageName() string {
	panic("implement me")
}
func (m Task) GetImplementor() model.OverrideOperation {
	return m.implementor
}

func (m Task) Run(manage *model.OperatorManage) error {
	if m.GetImplementor().GetOperation() == model.Operations.Provision {
		return m.Sync(manage)
	} else if m.GetImplementor().GetOperation() == model.Operations.Deletion {
		return m.Delete(manage)
	} else {
		return errors.New("UnSupport type of ResourceTask")
	}
}

type SyncResource struct {
	FileInfo
	model.Object
	Sync   SyncFunction
	Delete SyncFunction
}

func NewSyncResource(resMapping *K8sResourceMapping) *SyncResource {
	res := &SyncResource{FileInfo: FileInfo{}}
	res.Object = resMapping.ObjectGenerator()
	res.Sync = resMapping.Sync
	res.Delete = resMapping.Deletion
	return res
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
