package task

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"os"
	"testing"
	"text/template"
)

var metrics = `apiVersion: v1
kind: Service
metadata:
  name: asm-controller
  namespace: {{ .Release.Namespace }}
  labels:
    heritage: {{ $.Release.Service }}
    release: {{ $.Release.Name }}
    service_name: asm-controller
spec:
  ports:
  - name: metrics
    port: 8080
    protocol: TCP
    targetPort: 8080
  - name: http-webservice
    port: 8099
    protocol: TCP
    targetPort: 8099
  selector:
    app: asm-controller
  type: ClusterIP
`

func Test_load(t *testing.T) {
	tmpl, err := template.New("test").Parse(metrics)
	if err != nil {
		panic(err)
	}
	values := resource.TreeValue(map[string]interface{}{
		"Release.Namespace": "test-fy",
		"Release.Service":   "test-fy",
	})

	err = tmpl.Execute(os.Stdout, values)
	if err != nil {
		panic(err)
	}
}
