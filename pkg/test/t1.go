package test

type TestStruct struct {
	Name string
}

type Human interface {
	HasName() string
}

func (t TestStruct) HasName() string{
	return "mars"
}
