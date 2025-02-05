package interfaces

//somehow abstract that to be more generic
type ServiceBusConsumer interface {
	Consume(body []byte) error
	GetExchange() string
	GetQueueName() string
}

// some types of messages consumer
// just method Consume
// Use method of provider to setup but it still hidden behind abstraction
