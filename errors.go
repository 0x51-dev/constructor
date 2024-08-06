package constructor

import "fmt"

func NewCombinationError(a, b Node) error {
	return CombinationError{A: a, B: b}
}

type CombinationError struct {
	A, B Node
}

func (e CombinationError) Error() string {
	return fmt.Sprintf("cannot combine %T with %T", e.A, e.B)
}
