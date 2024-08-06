package constructor

type Any struct{}

func (a *Any) Combine(n Node) (Node, error) {
	if _, ok := n.(*Any); ok {
		// Any combined with Any is still Any.
		return a, nil
	}
	// Specialize to the other type.
	return n, nil
}

func (*Any) Equals(n Node) bool {
	// Any is only equal to itself.
	_, ok := n.(*Any)
	return ok
}
