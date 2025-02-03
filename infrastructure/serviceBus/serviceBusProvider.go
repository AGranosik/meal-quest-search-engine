package serviceBus

type ServiceBusProvider interface {
	Start() ServiceBusProvider
	Consume() ServiceBusProvider //consumer as parameter here, it will have parameters and consume method and doesnt care about type of channel
	WithExchange(exchangeName string) ServiceBusProvider
	WithQueue(queueName string, exchange string) ServiceBusProvider
	Stop()
}
