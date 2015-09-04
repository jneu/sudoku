package strategy

import "github.com/jneu/sudoku/common"

type Implementation interface {
	Process(message *common.Assertion) (done bool)
}

type Strategy struct {
	impl     *Implementation
	incoming <-chan *common.Assertion
	outgoing chan<- *common.Assertion
}

func New(impl *Implementation, awaitingProcessing chan<- *common.Assertion) *Strategy {
	return &Strategy{
		impl,
		make(chan *common.Assertion),
		awaitingProcessing,
	}
}
