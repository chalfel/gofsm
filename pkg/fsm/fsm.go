package fsm

import (
	"fmt"
	"sync"
)

type FSM[T comparable] struct {
	States map[T]State[T]
	mu     sync.RWMutex
}

func NewFSM[T comparable]() *FSM[T] {
	return &FSM[T]{
		States: make(map[T]State[T]),
		mu:     sync.RWMutex{},
	}
}

func (f *FSM[T]) AddState(state State[T]) {
	f.mu.Lock()
	f.States[state.Name] = state
	f.mu.Unlock()
}

func (f *FSM[T]) GoTo(from T, to T) error {
	f.mu.RLock()
	ts := f.States[to]

	if ts.Name != to {
		return fmt.Errorf("to state does not exist")
	}

	fs := f.States[from]

	if fs.Name != from {
		return fmt.Errorf("from state does not exist")
	}

	defer f.mu.RUnlock()
	err := fs.GoTo(ts)

	return err
}
