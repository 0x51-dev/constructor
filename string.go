package constructor

type String struct{}

func (s *String) Combine(n Node) (Node, error) {
	switch n.(type) {
	case *String:
		return s, nil
	case *Any:
		return new(String), nil
	default:
		return nil, NewCombinationError(s, n)
	}
}

func (s *String) Equals(n Node) bool {
	switch t := n.(type) {
	case *String:
		return true
	default:
		return t.Equals(s)
	}
}
