package v1alpha1

import "github.com/fyuan1316/operatorlib/api"

func (in AsmStatus) GetOperatorStatus() api.OperatorStatus {
	return in.OperatorStatus
}
