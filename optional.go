package constructor

type Optional struct {
	Type Node
}

func (o *Optional) Combine(n Node) (Node, error) {
	switch t := n.(type) {
	case *Optional:
		if n, err := o.Type.Combine(t.Type); err == nil {
			return &Optional{
				Type: n,
			}, nil
		}
	default:
		if n, err := o.Type.Combine(n); err == nil {
			return &Optional{
				Type: n,
			}, nil
		}
	}
	return &Or{
		Types: []Node{o, n},
	}, nil
}

func (o *Optional) Equals(n Node) bool {
	t, ok := n.(*Optional)
	if !ok {
		return false
	}
	return o.Type.Equals(t.Type)
}
