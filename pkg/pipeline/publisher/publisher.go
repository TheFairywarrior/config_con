package publisher


type Publisher interface {
	Publish(map[string]any) error
}

