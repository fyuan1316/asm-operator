package v1

import (
	"fmt"
	"github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "monitoring.coreos.com", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	fmt.Println("register serviceMonitor")
	SchemeBuilder.Register(
		&v1.ServiceMonitor{}, &v1.ServiceMonitorList{},
		&v1.PodMonitor{}, &v1.PodMonitorList{},
	)
}
