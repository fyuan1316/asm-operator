package task

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"io/ioutil"
	"os"
	"testing"
	"text/template"
)

func Test_load(t *testing.T) {

	f, err := ioutil.ReadFile("cluster-asm/resources/asm-metrics.yaml")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("test").Parse(string(f))
	if err != nil {
		panic(err)
	}
	//values := map[string]interface{}{
	//	"Release": map[string]interface{}{
	//		"Namespace": "test-fy",
	//	},
	//}
	values := resource.TreeValue(map[string]interface{}{
		"Release.Namespace": "test-fy",
		"Release.Service": "test-fy",
	})

	err = tmpl.Execute(os.Stdout, values)
	if err != nil {
		panic(err)
	}
}
