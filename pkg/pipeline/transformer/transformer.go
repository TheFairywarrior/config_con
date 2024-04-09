package transformer


type TransformStep interface {
	Transform(data any) (any, error)
}

type TransformedData interface {
	Get(string) any
	List() []string
}

type BaseTransformedData struct {
	data map[string]any
}

func (b *BaseTransformedData) Get(key string) any {
	return b.data[key]
}

func (b *BaseTransformedData) List() []string {
	keys := make([]string, 0, len(b.data))
	for k := range b.data {
		keys = append(keys, k)
	}
	return keys
}
