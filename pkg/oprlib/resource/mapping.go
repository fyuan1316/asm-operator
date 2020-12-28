package resource

import (
	"errors"
	"fmt"
	promv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	depv1alpha1 "github.com/fyuan1316/asm-operator/api/dep/v1alpha1"
	depv1beta1 "github.com/fyuan1316/asm-operator/api/dep/v1beta1"
	depv1beta2 "github.com/fyuan1316/asm-operator/api/dep/v1beta2"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource/sync"
	admissionv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func updateSyncResource(unStruct unstructured.Unstructured, res *SyncResource) error {
	var obj model.Object
	var syncFn func(client.Client, model.Object) error
	switch unStruct.GetKind() {
	case "CustomResourceDefinition":
		obj = &apiextensionsv1.CustomResourceDefinition{}
		syncFn = sync.FnCrd
	case "Deployment":
		obj = &appsv1.Deployment{}
		syncFn = sync.FnDeployment
	case "Service":
		obj = &corev1.Service{}
		syncFn = sync.FnService
	case "ClusterRoleBinding":
		obj = &rbacv1.ClusterRoleBinding{}
		syncFn = sync.FnClusterRoleBinding
	case "ClusterRole":
		obj = &rbacv1.ClusterRole{}
		syncFn = sync.FnClusterRole
	case "RoleBinding":
		obj = &rbacv1.RoleBinding{}
		syncFn = sync.FnRoleBinding
	case "ClusterConfig":
		obj = &depv1beta1.ClusterConfig{}
		syncFn = sync.FnCreateClusterConfig
	case "CaseMonitor":
		if unStruct.GroupVersionKind().Version == "v1beta2" {
			obj = &depv1beta2.CaseMonitor{}
			syncFn = sync.FnCreateCaseMonitorv1beta2
		} else {
			obj = &depv1beta1.CaseMonitor{}
			syncFn = sync.FnCreateCaseMonitorv1beta1
		}

	case "CanaryTemplate":
		obj = &depv1alpha1.CanaryTemplate{}
		syncFn = sync.FnCanaryTemplate
	case "Ingress":
		obj = &extv1beta1.Ingress{}
		syncFn = sync.FnIngress
	case "ServiceMonitor":
		obj = &promv1.ServiceMonitor{}
		syncFn = sync.FnServiceMonitor
	case "PodMonitor":
		obj = &promv1.PodMonitor{}
		syncFn = sync.FnPodMonitor
	case "ServiceAccount":
		obj = &corev1.ServiceAccount{}
		syncFn = sync.FnServiceAccount
	case "ValidatingWebhookConfiguration":
		obj = &admissionv1.ValidatingWebhookConfiguration{}
		syncFn = sync.FnValidatingWebhookConfiguration
	default:
		return errors.New(fmt.Sprintf("UnSupport Resource Kind %s", unStruct.GetKind()))
	}
	res.Object = obj
	res.Sync = syncFn
	res.Delete = sync.FnDelete
	return nil
}
