package constructor

type Array struct {
	Type Node
}

func (a *Array) Combine(n Node) (Node, error) {
	switch t := n.(type) {
	case *Optional:
		if n, err := a.Type.Combine(t.Type); err == nil {
			return &Optional{
				Type: n,
			}, nil
		}
	case *Array:
		if n, err := a.Type.Combine(t.Type); err == nil {
			return &Array{
				Type: n,
			}, nil
		}
	}
	return &Or{
		Types: []Node{a, n},
	}, nil
}

func (a *Array) Equals(n Node) bool {
	t, ok := n.(*Array)
	if !ok {
		return false
	}
	return a.Type.Equals(t.Type)
}
