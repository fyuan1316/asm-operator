package task

import "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"

func GetDefaults() map[string]interface{} {
	values := resource.TreeValue(map[string]interface{}{
		"Release.Namespace": "default",
		"Release.Name":      "asm-operator",
		"Release.Service":   "asm-operator",
	})

	return values
}
