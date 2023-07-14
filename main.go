package main

import (
	"fmt"
	"github.com/chalfel/gofsm/pkg/fsm"
)

type TransactionState string

const (
	TransactionStateBlocked   TransactionState = "blocked"
	TransactionStateUnblocked TransactionState = "unblocked"
)

type TransactionStateTransitor[T TransactionState] struct {
}

func (t *TransactionStateTransitor[T]) OnTransition(to fsm.State[T]) error {
	return nil
}

func main() {
	f := fsm.NewFSM[TransactionState]()

	f.AddState(fsm.State[TransactionState]{
		Name:      TransactionStateBlocked,
		To:        []TransactionState{TransactionStateUnblocked},
		From:      []TransactionState{TransactionStateUnblocked},
		Transitor: &TransactionStateTransitor[TransactionState]{},
	})

	f.AddState(fsm.State[TransactionState]{
		Name:      TransactionStateUnblocked,
		To:        []TransactionState{TransactionStateBlocked},
		From:      []TransactionState{TransactionStateBlocked},
		Transitor: &TransactionStateTransitor[TransactionState]{},
	})

	err := f.GoTo(TransactionStateBlocked, TransactionStateUnblocked)

	if err != nil {
		fmt.Println(err.Error())
	}

}
