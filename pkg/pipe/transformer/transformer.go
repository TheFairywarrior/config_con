package transformer

type Step interface {
	Process(any) (any, error)
}

type Transformer struct {
	Name  string
	Steps []Step
}
