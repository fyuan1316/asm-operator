package mock

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource/sync"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeployTask struct {
	resource.SyncManager
}

func (m DeployTask) GetStageName() string {
	panic("implement me")
}

var deployTask DeployTask
var deploy1 = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: sleep-fy
  namespace: default
spec:
  selector:
    matchLabels:
      app: sleep-fy
  template:
    metadata:
      labels:
        app: sleep-fy
    spec:
      containers:
      - command:
        - /bin/sleep
        - 3650d
        image: governmentpaas/curl-ssl
        name: sleep

`
var svc1 = `apiVersion: v1
kind: Service
metadata:
  labels:
    app: sleep-fy
  name: sleep-fy
  namespace: default
spec:
  ports:
  - name: http
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: sleep-fy
  type: ClusterIP
`

func init() {
	deployTask = DeployTask{}
	res := resource.SyncResource{
		Object: &appsv1.Deployment{},
		Sync:   sync.FnDeployment,
		//Sync: func(client client.Client, object manage.Object) error {
		//	deploy := appsv1.Deployment{}
		//	err := client.Get(context.Background(),
		//		types.NamespacedName{Namespace: object.GetNamespace(), Name: object.GetName()},
		//		&deploy,
		//	)
		//	if err != nil {
		//		if errors.IsNotFound(err) {
		//			errCreate := client.Objecteate(context.Background(), object)
		//			if errCreate != nil {
		//				return errCreate
		//			}
		//		}
		//		return err
		//	} else {
		//		//update
		//		wanted := object.(*appsv1.Deployment)
		//		if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		//			deploy.Spec = wanted.Spec
		//			if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
		//				return errUpd
		//			}
		//		}
		//	}
		//	return nil
		//},
	}
	res.SetOwnerRef()

	resSvc := resource.SyncResource{
		Object: &corev1.Service{},
		Sync:   sync.FnService,
	}
	resSvc.SetOwnerRef()
	err := deployTask.Load(deploy1, &res, map[string]interface{}{})
	if err != nil {
		panic(err)
	}
	err = deployTask.Load(svc1, &resSvc, map[string]interface{}{})
	if err != nil {
		panic(err)
	}
}

var _ model.ExecuteItem = DeployTask{}

func (m DeployTask) PreRun(client client.Client) error {
	fmt.Println("DeployTask prerun")
	return nil
}

func (m DeployTask) PostRun(client client.Client) error {
	fmt.Println("DeployTask PostRun")
	return nil
}

func (m DeployTask) PreCheck(client client.Client) (bool, error) {
	fmt.Println("DeployTask PreCheck")
	return true, nil
}

func (m DeployTask) PostCheck(client client.Client) (bool, error) {
	fmt.Println("DeployTask PostCheck")
	return true, nil
}

func (m DeployTask) Run(om *model.OperatorManage) error {
	fmt.Println("DeployTask Run")
	err := m.Sync(om)
	return err
}

func (m DeployTask) LiveNess() bool {
	return true
}
