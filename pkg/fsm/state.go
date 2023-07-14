package fsm

import "fmt"

type Transitor[T comparable] interface {
	OnTransition(to State[T]) error
}

type State[T comparable] struct {
	Name      T
	To        []T
	From      []T
	Transitor Transitor[T]
}

func (s *State[T]) CanGoTo(to State[T]) bool {
	for _, state := range s.To {
		fmt.Println(state)
		fmt.Println(to.Name)
		if state == to.Name {
			return true
		}
	}

	return false
}

func (s *State[T]) CanGoFrom(from State[T]) bool {
	for _, state := range s.From {
		if state == from.Name {
			return true
		}
	}

	return false
}
func (s *State[T]) GoTo(to State[T]) error {
	if !s.CanGoTo(to) {
		fmt.Println("unable to go to the next status")
		return fmt.Errorf("unable to go to the next status")
	}

	if !to.CanGoFrom(*s) {
		fmt.Println("unable to go from the current status")
		return fmt.Errorf("unable to go from the current status")
	}

	for _, state := range s.To {
		if state == to.Name {
			return s.Transitor.OnTransition(to)
		}
	}

	return fmt.Errorf("unable to go to the next status")
}
