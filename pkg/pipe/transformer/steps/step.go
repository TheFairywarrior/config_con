package steps


type Step interface {
	Process(any) (any, error)
}

type MultipleOutput interface {
	Key() string
	ListData() ([]any, error)
}
