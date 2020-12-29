package resource

import (
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

type FileInfo struct {
	ChargeByOperator *bool
}

type FileOption func(spec *FileInfo)

func KeepResourceAfterOperatorDeleted() FileOption {
	return func(spec *FileInfo) {
		b := true
		spec.ChargeByOperator = &b
	}
}

func (m *Task) LoadFile(filePath string, opts ...FileOption) error {

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return m.Load(string(bytes))
}
func (m *Task) Load(objectStr string, opts ...FileOption) error {
	var err error

	//meta := manage.TypeObjectMeta{}
	// render values TODO mergeDefaults
	renderedObjectStr, err := Parse(objectStr, m.TemplateValues)
	if err != nil {
		return err
	}
	unStruct := unstructured.Unstructured{}
	err = yaml.Unmarshal([]byte(renderedObjectStr), &unStruct)

	if err != nil {
		return err
	}
	if &unStruct == nil {
		fmt.Println()
	}
	if res, err := findResource(unStruct, m.ResourceMappings); err != nil {
		return err
	} else {
		for _, opt := range opts {
			opt(&res.FileInfo)
		}
		//if err = updateSyncResource(unStruct, res); err != nil {
		//	return err
		//}

		object := res.Object
		err = yaml.Unmarshal([]byte(renderedObjectStr), object)
		if err != nil {
			return err
		}
		//TODO fy 按文件名顺序排序
		objKey := fmt.Sprintf("-%s-%s-%s", object.GetObjectKind().GroupVersionKind().Kind,
			object.GetNamespace(),
			object.GetName(),
		)
		if m.K8sResource == nil {
			m.K8sResource = make(map[string]SyncResource)
		}

		m.K8sResource[objKey] = *res
		return err
	}
}
