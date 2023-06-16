package internal

type BaseInterface interface {
	DoGet() (string, error)
	DoPost(data Data) (Data, error)
	DoPut() (string, error)
}

type BaseStruct struct {
}

func NewBase() *BaseStruct {
	return &BaseStruct{}
}

func (b *BaseStruct) DoGet() (string, error) {
	return "Hello DoGet", nil
}

func (b *BaseStruct) DoPost(data Data) (Data, error) {
	return data, nil
}

func (b *BaseStruct) DoPut() (string, error) {
	return "Hello DoPut", nil
}
