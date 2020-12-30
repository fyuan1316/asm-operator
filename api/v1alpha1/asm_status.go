package v1alpha1

func (in *AsmStatus) setState(state AsmState) {
	in.State = state
}

func (in *AsmStatus) SetState(isReady, isHealthy bool) {
	if isHealthy {
		in.setState(AsmStates.Health)
		return
	}
	if isReady {
		in.setState(AsmStates.Ready)
		return
	}
	in.setState(AsmStates.NotReady)
	return
}
