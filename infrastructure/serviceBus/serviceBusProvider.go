package serviceBus

type ServiceBusProvider interface {
	Start() ServiceBusProvider
	Consume() ServiceBusProvider
	WithExchange(exchangeName string) ServiceBusProvider
	WithQueue(queueName string, exchange string) ServiceBusProvider
	Stop()
}
