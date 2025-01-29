package serviceBus

type ServiceBusProvider interface {
	Start() (ServiceBusProvider, error)
	Consume(f func()) (ServiceBusProvider, error)
	Stop()
}
