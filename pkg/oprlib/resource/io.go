package resource

import (
	"os"
	"path/filepath"
	"sort"
)

type Order string

var ResourceOrder = struct {
	ASC  Order
	DESC Order
}{"ASC", "DESC"}

type Option func(spec *ResSpec)
type ResSpec struct {
	Order  Order
	Suffix string
}

func Asc() Option {
	return func(spec *ResSpec) {
		spec.Order = ResourceOrder.ASC
	}
}
func Desc() Option {
	return func(spec *ResSpec) {
		spec.Order = ResourceOrder.DESC
	}
}
func Suffix(s string) Option {
	return func(spec *ResSpec) {
		spec.Suffix = s
	}
}

func GetFilesInFolder(folderPath string, opts ...Option) ([]string, error) {
	resSpec := &ResSpec{Order: ResourceOrder.ASC}
	for _, opt := range opts {
		opt(resSpec)
	}
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if resSpec.Suffix != "" {
			if filepath.Ext(path) != resSpec.Suffix {
				return nil
			}
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.SliceStable(files, func(i, j int) bool {
		var cmp bool
		if files[i] < files[j] {
			cmp = true
		}
		if resSpec.Order == ResourceOrder.ASC {
			return cmp
		} else {
			return !cmp
		}
	})
	return files, nil
}
