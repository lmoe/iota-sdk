package methods

type NoType any

type BaseRequest[T any] struct {
	Name string `json:"name" yaml:"name" mapstructure:"name"`
	Data T      `json:"data" yaml:"data" mapstructure:"data"`
}

func NewBaseRequest[T any](name string, data T) BaseRequest[T] {
	return BaseRequest[T]{
		Name: name,
		Data: data,
	}
}

func NewBaseRequestNoData(name string) BaseRequest[NoType] {
	return BaseRequest[NoType]{
		Name: name,
	}
}
