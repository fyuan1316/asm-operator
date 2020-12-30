package resource

import (
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

type FileInfo struct {
	FilePath         string
	ChargeByOperator *bool
}

type FileOption func(spec *FileInfo)

func SetFilePath(path string) FileOption {
	return func(spec *FileInfo) {
		spec.FilePath = path
	}
}
func KeepResourceAfterOperatorDeleted() FileOption {
	return func(spec *FileInfo) {
		b := true
		spec.ChargeByOperator = &b
	}
}

func (m *Task) LoadFile(filePath string, opts ...FileOption) error {
	//execution order by fileName order
	opts = append(opts, SetFilePath(filePath))

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return m.Load(string(bytes), opts...)
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

		object := res.Object
		err = yaml.Unmarshal([]byte(renderedObjectStr), object)
		if err != nil {
			return err
		}
		//TODO fy 按文件名顺序排序
		objKey := fmt.Sprintf("%s-%s-%s-%s",
			res.FilePath,
			object.GetObjectKind().GroupVersionKind().Kind,
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

func (m *Task) LoadResources(objectStrs map[string]string, opts ...FileOption) error {
	var err error

	//meta := manage.TypeObjectMeta{}
	// render values TODO mergeDefaults
	//renderedObjectStr, err := Parse(objectStr, m.TemplateValues)
	//if err != nil {
	//	return err
	//}

	for _, objectStr := range objectStrs {
		unStruct := unstructured.Unstructured{}
		err = yaml.Unmarshal([]byte(objectStr), &unStruct)
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

			object := res.Object
			err = yaml.Unmarshal([]byte(objectStr), object)
			if err != nil {
				return err
			}
			//TODO fy 按文件名顺序排序
			objKey := fmt.Sprintf("%s-%s-%s-%s",
				res.FilePath,
				object.GetObjectKind().GroupVersionKind().Kind,
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
	return nil
}
