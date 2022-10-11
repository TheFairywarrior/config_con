package api
//  Package to check and handle the creation of consumers.

var ApiRoutes chan string
var amountOfConsumers int

func InitRoutes(size int) {
	ApiRoutes = make(chan string, size)
	amountOfConsumers = size
}


func WhenReady() {
	count := 0
	for {
		<-ApiRoutes // Wait for a route to be added.
		count++
		if count == amountOfConsumers {
			break;
		}
	}
}

