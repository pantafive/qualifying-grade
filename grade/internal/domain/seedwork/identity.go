package seedwork

import "github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork/interfaces"

func NewIdentity[T comparable](value T) (Identity[T, interfaces.Identity[T]], error) {
	return Identity[T, interfaces.Identity[T]]{value: value}, nil
}

// The way to fix issue of generics:
// https://issuemode.com/issues/golang/go/105227904

type Identity[T comparable, C interfaces.Identity[T]] struct {
	value T
}

func (id Identity[T, C]) Equals(other C) bool {
	return id.CreateMemento() == other.CreateMemento()
}

func (id Identity[T, C]) CreateMemento() T {
	return id.value
}
