package serviceBus

type ServiceBusConsumer interface {
	Consume(body []byte) error
}

// some types of messages consumer
// just method Consume
// Use method of provider to setup but it still hidden behind abstraction
