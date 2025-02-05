package interfaces

type ServiceBusProvider interface {
	Start() ServiceBusProvider
	Consume(consumer ServiceBusConsumer) ServiceBusProvider //consumer as parameter here, it will have parameters and consume method and doesnt care about type of channel
	Stop()
}
