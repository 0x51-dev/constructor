package constructor

type Number struct{}

func (i *Number) Combine(n Node) (Node, error) {
	switch n.(type) {
	case *Number:
		return i, nil
	case *Any:
		return new(Number), nil
	default:
		return nil, NewCombinationError(i, n)
	}
}

func (i *Number) Equals(n Node) bool {
	switch t := n.(type) {
	case *Number:
		return true
	default:
		return t.Equals(i)
	}
}
