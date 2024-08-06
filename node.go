package constructor

import "fmt"

func (*Any) node() {}

func (*Array) node() {}

type Node interface {
	fmt.Stringer

	Combine(Node) (Node, error)
	Equals(Node) bool
	node() // Make `Node` non-extendable.
}

func (*Number) node() {}

func (*Optional) node() {}

func (*Or) node() {}

func (*String) node() {}

func (*Struct) node() {}
