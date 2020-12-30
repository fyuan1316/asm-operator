package manage

import (
	"github.com/fyuan1316/asm-operator/pkg/logging"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/options"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	logger = logging.RegisterScope("controller.oprlib")
)

func NewOperatorManage(client client.Client, cr model.Object, opts ...options.Option) *model.OperatorManage {
	managerSpec := &model.OperatorManage{
		K8sClient: client,
		CR:        cr,
	}
	for _, opt := range opts {
		opt(managerSpec)
	}
	return managerSpec
}
