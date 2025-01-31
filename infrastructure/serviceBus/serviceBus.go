package serviceBus

// will be extendended if needed
type ServiceBusProvider interface {
	Start() ServiceBusProvider
	Consume() ServiceBusProvider
	WithExchange(exchangeName string) ServiceBusProvider
	WithQueue(queueName string, exchange string) ServiceBusProvider
	Stop()
}
