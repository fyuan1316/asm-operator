package ordered

import "github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"

type ObjectMap struct {
	maps     map[string]model.Object
	position map[string]int
}
