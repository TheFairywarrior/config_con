package override


type FiberContext interface {
	BodyParser(interface{}) error 
	GetReqHeaders() map[string]string
	Status(int) FiberContext
	JSON(interface{}) error
}
