package resource

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
)

var _ model.ExecuteItem = ChartTask{}
var _ model.Chart = ChartTask{}

type ChartTask struct {
	FileTask
	////子类override 接口
	//implementor model.OverrideOperation
	//
	////资源mappings hook
	//ResourceMappings map[metav1.TypeMeta]*K8sResourceMapping
	//
	////ResourceOptions  map[string]YamlResource
	////归属任务的资源集
	//K8sResource map[string]SyncResource
	//
	////任务层面资源是否随operator删除的设定
	//KeepResourceAfterOperatorDeleted *bool
	//
	////任务级别的values数据，一般对应到一个chart的values
	////TemplateValues map[string]interface{}

	ChartDir string
}

func (m ChartTask) Reload(values map[string]interface{}) error {
	fmt.Println("chart reload")
	var (
		files map[string]string
		err   error
	)
	if files, err = resource2.GetChartResources(m.ChartDir, values); err != nil {
		return err
	}
	if err = m.LoadResources(files); err != nil {
		return err
	}
	return nil
}

func (m *ChartTask) Override(operation model.OverrideOperation) {
	m.implementor = operation
}
func (m ChartTask) Name() string {
	panic("implement me")
}
func (m ChartTask) GetImplementor() model.OverrideOperation {
	return m.implementor
}

//func (m ChartTask) Run(ctx *model.OperatorContext) error {
//
//	if m.GetImplementor().GetOperation() == model.Operations.Provision {
//		return m.Sync(ctx)
//	} else if m.GetImplementor().GetOperation() == model.Operations.Deletion {
//		return m.Delete(ctx)
//	} else {
//		return errors.New("UnSupport type of ResourceTask")
//	}
//}
//
//type SyncResource struct {
//	FileInfo
//	model.Object
//	Sync   SyncFunction
//	Delete SyncFunction
//}
//
//func NewSyncResource(resMapping *K8sResourceMapping) *SyncResource {
//	res := &SyncResource{FileInfo: FileInfo{}}
//	res.Object = resMapping.ObjectGenerator()
//	res.Sync = resMapping.Sync
//	res.Delete = resMapping.Deletion
//	return res
//}
//
//func (m *SyncResource) SetObject(o model.Object) {
//	m.Object = o
//}
//func (m *SyncResource) SetOwnerRef() {
//	t := true
//	m.ChargeByOperator = &t
//}
//func (m *SyncResource) IsChargedByOwnerRef() *bool {
//	return m.ChargeByOperator
//}
//
//func (m *ChartTask) Sync(ctx *model.OperatorContext) error {
//	for _, res := range m.K8sResource {
//		if err := res.Sync(ctx.K8sClient, res.Object); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (m *ChartTask) Delete(ctx *model.OperatorContext) error {
//	for _, res := range m.K8sResource {
//		//资源参数优先
//		if res.IsChargedByOwnerRef() != nil && *res.IsChargedByOwnerRef() ||
//			m.KeepResourceAfterOperatorDeleted != nil && !*m.KeepResourceAfterOperatorDeleted {
//			if err := res.Delete(ctx.K8sClient, res.Object); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
//
