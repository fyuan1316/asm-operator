package resource

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"sigs.k8s.io/yaml"
	"text/template"
)

func toYaml(v interface{}) string {
	data, err := yaml.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template.
		fmt.Printf("ToYaml err:%q", err)
		return ""
	}
	return string(data)
}

func Parse(tpl string /*unStruct unstructured.Unstructured,*/, values map[string]interface{}) (string, error) {
	//key := fmt.Sprintf("%s-%s-%s", unStruct.GetKind(), unStruct.GetNamespace(), unStruct.GetName())
	fm := sprig.GenericFuncMap()
	fm["toYaml"] = toYaml
	tmpl, err := template.New("key").Funcs(fm).Parse(string(tpl))
	//tmpl, err := template.New("key").Parse(string(tpl))
	if err != nil {
		return "", err
	}
	//values := map[string]interface{}{
	//	"Release": map[string]interface{}{
	//		"Namespace": "test-fy",
	//	},
	//}
	//values := TreeValue(map[string]interface{}{
	//	"Release.Namespace": "test-fy",
	//	"Release.Service":   "test-fy",
	//})
	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, values)
	return buf.String(), nil
}
