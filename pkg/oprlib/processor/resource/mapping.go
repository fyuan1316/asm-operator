package resource

import (
	"errors"
	"fmt"
	promv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	depv1alpha1 "github.com/fyuan1316/asm-operator/api/dep/v1alpha1"
	depv1beta1 "github.com/fyuan1316/asm-operator/api/dep/v1beta1"
	depv1beta2 "github.com/fyuan1316/asm-operator/api/dep/v1beta2"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource/sync"
	admissionv1 "k8s.io/api/admissionregistration/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SyncFunction func(client.Client, model.Object) error

type K8sResourceMapping struct {
	Object   model.Object
	Sync     SyncFunction
	Deletion SyncFunction
}

var internalMappings = map[metav1.TypeMeta]K8sResourceMapping{
	metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1",
	}: {
		Object: &apiextensionsv1.CustomResourceDefinition{},
		Sync:   sync.FnCrd,
	},
	metav1.TypeMeta{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
	}: {
		Object: &appsv1.Deployment{},
		Sync:   sync.FnDeployment,
	},
	metav1.TypeMeta{
		Kind:       "Service",
		APIVersion: "v1",
	}: {
		Object: &corev1.Service{},
		Sync:   sync.FnService,
	},
	metav1.TypeMeta{
		Kind:       "ClusterRoleBinding",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		Object: &rbacv1.ClusterRoleBinding{},
		Sync:   sync.FnClusterRoleBinding,
	},
	metav1.TypeMeta{
		Kind:       "ClusterRole",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		Object: &rbacv1.ClusterRole{},
		Sync:   sync.FnClusterRole,
	},
	metav1.TypeMeta{
		Kind:       "RoleBinding",
		APIVersion: "rbac.authorization.k8s.io/v1beta1",
	}: {
		Object: &rbacv1.RoleBinding{},
		Sync:   sync.FnRoleBinding,
	},
	metav1.TypeMeta{
		Kind:       "ClusterConfig",
		APIVersion: "asm.alauda.io/v1beta1",
	}: {
		Object: &depv1beta1.ClusterConfig{},
		Sync:   sync.FnCreateClusterConfig,
	},
	metav1.TypeMeta{
		Kind:       "CaseMonitor",
		APIVersion: "asm.alauda.io/v1beta2",
	}: {
		Object: &depv1beta2.CaseMonitor{},
		Sync:   sync.FnCreateCaseMonitorv1beta2,
	},
	metav1.TypeMeta{
		Kind:       "CaseMonitor",
		APIVersion: "asm.alauda.io/v1beta1",
	}: {
		Object: &depv1beta1.CaseMonitor{},
		Sync:   sync.FnCreateCaseMonitorv1beta1,
	},
	metav1.TypeMeta{
		Kind:       "CanaryTemplate",
		APIVersion: "asm.alauda.io/v1alpha1",
	}: {
		Object: &depv1alpha1.CanaryTemplate{},
		Sync:   sync.FnCanaryTemplate,
	},
	metav1.TypeMeta{
		Kind:       "Ingress",
		APIVersion: "extensions/v1beta1",
	}: {
		Object: &extv1beta1.Ingress{},
		Sync:   sync.FnIngress,
	},
	metav1.TypeMeta{
		Kind:       "ServiceMonitor",
		APIVersion: "monitoring.coreos.com/v1",
	}: {
		Object: &promv1.ServiceMonitor{},
		Sync:   sync.FnServiceMonitor,
	},
	metav1.TypeMeta{
		Kind:       "PodMonitor",
		APIVersion: "monitoring.coreos.com/v1",
	}: {
		Object: &promv1.PodMonitor{},
		Sync:   sync.FnPodMonitor,
	},
	metav1.TypeMeta{
		Kind:       "ServiceAccount",
		APIVersion: "v1",
	}: {
		Object: &corev1.ServiceAccount{},
		Sync:   sync.FnServiceAccount,
	},
	metav1.TypeMeta{
		Kind:       "ValidatingWebhookConfiguration",
		APIVersion: "admissionregistration.k8s.io/v1beta1",
	}: {
		Object: &admissionv1.ValidatingWebhookConfiguration{},
		Sync:   sync.FnValidatingWebhookConfiguration,
	},
}

func GetInternalMappings() map[metav1.TypeMeta]K8sResourceMapping {
	for k := range internalMappings {
		m := internalMappings[k]
		m.Deletion = sync.FnDelete
	}
	return internalMappings
}

func findResource(unStruct unstructured.Unstructured, userMapping map[metav1.TypeMeta]K8sResourceMapping) (*SyncResource, error) {
	key := metav1.TypeMeta{
		Kind:       unStruct.GetKind(),
		APIVersion: unStruct.GetAPIVersion(),
	}
	var found *K8sResourceMapping
	if len(userMapping) > 0 {
		if v, ok := userMapping[key]; ok {
			found = &v
		}
	}
	if found == nil {
		if v, ok := GetInternalMappings()[key]; ok {
			found = &v
		}
	}
	if found == nil {
		return nil, errors.New(fmt.Sprintf("NotFound type %s in mappings", key))
	}
	var convertor *SyncResource
	return convertor.FromMappings(found), nil
}

/*
func updateSyncResource(unStruct unstructured.Unstructured, res *SyncResource) error {
	var obj model.Object
	var syncFn SyncFunction
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
*/
