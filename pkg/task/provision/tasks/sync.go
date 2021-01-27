package tasks

import (
	"context"
	"fmt"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
	pkgerrors "github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

type ProvisionResourcesTask struct {
	*chart.ChartTask
}

func (p ProvisionResourcesTask) GetName() string {
	return "asm-controlplane-install"
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}
var _ model.PreCheck = ProvisionResourcesTask{}
var _ model.PreRun = ProvisionResourcesTask{}

func (p ProvisionResourcesTask) PreRun(ctx *model.OperatorContext) error {
	key := client.ObjectKey{Name: ctx.InstalledNamespace}
	current := corev1.Namespace{}
	current.Name = ctx.InstalledNamespace
	err := ctx.K8sClient.Get(context.TODO(), key, &current)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return pkgerrors.WithStack(err)
		}
		if cErr := ctx.K8sClient.Create(context.TODO(), &current); cErr != nil {
			return pkgerrors.WithStack(err)
		}
		return nil
	}
	return nil
}

var deps = []string{
	"virtualservices.networking.istio.io",
	"serviceentries.networking.istio.io",
	"peerauthentications.security.istio.io",
	"gateways.networking.istio.io",
}

func (p ProvisionResourcesTask) PreCheck(ctx *model.OperatorContext) (bool, error) {
	//check asm-controller dependency (istio crds)
	crdList := &apiextensionsv1.CustomResourceDefinitionList{}
	err := ctx.K8sClient.List(context.Background(), crdList)
	if err != nil {
		return false, pkgerrors.WithStack(err)
	}

	var count = 0
	var missCrds []string
	var crds = make(map[string]bool)
	for _, crd := range crdList.Items {
		crds[crd.Name] = false
	}
	for _, dep := range deps {
		if _, exist := crds[dep]; exist {
			count++
		} else {
			missCrds = append(missCrds, dep)
		}
	}
	if count != len(deps) {
		err = pkgerrors.New(fmt.Sprintf("not found: %s", strings.Join(missCrds, ",")))
		return false, err
	}
	return true, nil
}

func (p ProvisionResourcesTask) IsReady(oCtx *model.OperatorContext) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(oCtx *model.OperatorContext) bool {
	return false
}

func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

var ClusterAsmResDir = "files/provision/cluster-asm"

func SetUpProvision() {
	ProvisionResources = ProvisionResourcesTask{
		&chart.ChartTask{
			Dir: ClusterAsmResDir,
		},
	}
	ProvisionResources.Init()
	ProvisionResources.Override(ProvisionResources)
}
