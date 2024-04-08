package transformer


type TransformStep interface {
	Transform(data any) (any, error)
}