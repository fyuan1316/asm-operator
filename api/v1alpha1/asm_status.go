package v1alpha1

import "github.com/fyuan1316/asm-operator/pkg/oprlib/api"

func (in *AsmStatus) setState(state api.OperatorState) {
	in.State = state
}

func (in *AsmStatus) SetState(isReady, isHealthy bool) {
	if isHealthy {
		in.setState(api.OperatorStates.Health)
		return
	}
	if isReady {
		in.setState(api.OperatorStates.Ready)
		return
	}
	in.setState(api.OperatorStates.NotReady)
	return
}
