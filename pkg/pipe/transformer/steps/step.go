package steps


type Step interface {
	Process(any) (any, error)
}
