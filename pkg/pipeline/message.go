package pipeline


type Message interface {
	Content() any
}


type DefaultMessage struct {
	content any
}

func (m *DefaultMessage) Content() any {
	return m.content
}

func NewDefaultMessage(content any) Message {
	return &DefaultMessage{
		content: content,
	}
}
