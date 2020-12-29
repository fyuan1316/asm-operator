package shell

import (
	"strings"
)

type ScriptFunc func() error
type Scripts struct {
	preCheck  []string
	postCheck []string

	preRun  []string
	postRun []string

	run []string
}

func (s *ScriptManager) Load(filePath string) error {
	lowercaseFilePath := strings.ToLower(filePath)
	if strings.HasPrefix(lowercaseFilePath, "precheck") {
		s.preCheck = append(s.preCheck, filePath)
	} else if strings.HasPrefix(lowercaseFilePath, "postcheck") {
		s.postCheck = append(s.postCheck, filePath)
	} else if strings.HasPrefix(lowercaseFilePath, "prerun") {
		s.preRun = append(s.preRun, filePath)
	} else if strings.HasPrefix(lowercaseFilePath, "postrun") {
		s.postRun = append(s.postRun, filePath)
	} else {
		s.run = append(s.run, filePath)
	}
	return nil
}

type ScriptManager struct {
	Scripts
}

//
//var _  model.ExecuteItem  = ScriptManager{}
//
//func (s *ScriptManager) Execute() error {
//
//}
