package strategy

type Assertion struct {
	value     int
	cellIndex int
}

type Implementation interface {
	Process(message *Assertion) (done bool)
}

type Strategy struct {
	impl     *Implementation
	incoming <-chan *Assertion
	outgoing chan<- *Assertion
}

func New(impl *Implementation, awaitingProcessing chan<- *Assertion) *Strategy {
	return &Strategy{
		impl,
		make(chan *Assertion),
		awaitingProcessing,
	}
}
