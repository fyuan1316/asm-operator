/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const (
	DefaultInterval = 10 * time.Second
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CaseMonitorSpec defines the desired state of CaseMonitor
type CaseMonitorSpec struct {
	Interval string `json:"interval,omitempty"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	JobLabel string `json:"jobLabel,omitempty"`
	// A list of urls allowed as part of this ServiceMonitor.
	PingURLs []PingURL `json:"pingurls,omitempty"`
	// Selector transfers labels on the Kubernetes Pod onto the target.
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// NamespaceSelector to select which namespaces the Endpoints objects are discovered from.
	NamespaceSelector NamespaceSelector `json:"namespaceSelector,omitempty"`
	// SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
	MonitorResults MonitorResult `json:"monitorResults,omitempty"`
	// A list of crds allowed as part of this CaseMonitor.
	CRDs []CRD `json:"crds,omitempty"`
}

// Url defines a scrapeable url serving Prometheus metrics.
// +k8s:openapi-gen=true
type PingURL struct {
	// HTTP path to scrape for metrics.
	Name string `json:"name"`
	// HTTP path to scrape for metrics.
	Host string `json:"host,omitempty"`
	// Name of the service port this url refers to. Mutually exclusive with targetPort.
	Port string `json:"port,omitempty"`
	// HTTP path to scrape for metrics.
	Path string `json:"path,omitempty"`
	// HTTP scheme to use for scraping.
	Scheme string `json:"scheme,omitempty"`
	// Optional HTTP URL parameters
	Params map[string][]string `json:"params,omitempty"`
	// Interval at which metrics should be scraped
	Interval string `json:"interval,omitempty"`
	// Timeout after which the scrape is ended
	ScrapeTimeout string `json:"scrapeTimeout,omitempty"`
	// TLS config to use when scraping the url
	TLSConfig *TLSConfig `json:"tlsConfig,omitempty"`
	// url for whole url
	URL string `json:"url,omitempty"`
	//List of custom errors.
	CustomErrors []string `json:"customErrors,omitempty"`
}

// TLSConfig specifies TLS config parameters.
// +k8s:openapi-gen=true
type TLSConfig struct {
	// Disable target certificate validation.
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
}

// A list of crds allowed as part of this CaseMonitor.
type CRD struct {
	// Name of job.
	Name string `json:"name"`
	// CRD group for apigroups GroupVersionKind.
	Group string `json:"group"`
	// kind for GroupVersionKind.
	Kind string `json:"kind,omitempty"`
	// version for GroupVersionKind.
	Version string `json:"version,omitempty"`
	// name of crd.
	CRDName string `json:"crdname,omitempty"`
	// namespace of crd.
	Namespace string `json:"namespace,omitempty"`
	// operator of the crd, like Exists, HttpPing.
	Operator string `json:"operator,omitempty"`
	// TLS config to use when scraping the url
	BasicAuth *BasicAuth `json:"basicAuth,omitempty"`
	// url for whole url
	URL string `json:"url,omitempty"`
	//List of custom errors.
	CustomErrors []string `json:"customErrors,omitempty"`
}

// BasicAuth allow an http to authenticate over basic authentication.
type BasicAuth struct {
	// property name of password.
	Password string `json:"password,omitempty"`
	// property name of username.
	UserName string `json:"username,omitempty"`
}

// CaseMonitorStatus defines the observed state of CaseMonitor
type CaseMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// NamespaceSelector is a selector for selecting either all namespaces or a
// list of namespaces.
// +k8s:openapi-gen=true
type NamespaceSelector struct {
	// List of namespace names.
	MatchNames []string `json:"matchNames,omitempty"`

	// TODO(fabxc): this should embed metav1.LabelSelector eventually.
	// Currently the selector is only used for namespaces which require more complex
	// implementation to support label selections.
}

// MonitorResults is the results for monotors
// +k8s:openapi-gen=true
type MonitorResult struct {
	// +nullable: true
	LatestUpdated metav1.Time `json:"latestUpdated,omitempty"`
	Targets       []Target    `json:"targets,omitempty"`
}

// target for monitor
// +k8s:openapi-gen=true
type Target struct {
	Name       string   `json:"name,omitempty"`
	TargetType string   `json:"targetType,omitempty"`
	LastErrors []string `json:"lastErrors,omitempty"`
	Health     string   `json:"running,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// CaseMonitor is the Schema for the casemonitors API
type CaseMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaseMonitorSpec   `json:"spec,omitempty"`
	Status CaseMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CaseMonitorList contains a list of CaseMonitor
type CaseMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CaseMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CaseMonitor{}, &CaseMonitorList{})
}

// GetInterval returns the   interval (default 10s)
func (c *CaseMonitor) GetInterval() time.Duration {
	if c.Spec.Interval == "" {
		return DefaultInterval
	}

	interval, err := time.ParseDuration(fmt.Sprintf("%ss", c.Spec.Interval))
	if err != nil {
		return DefaultInterval
	}

	if interval < 2*time.Second {
		return DefaultInterval
	}

	return interval
}
