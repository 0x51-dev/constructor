package constructor

import (
	"slices"
	"strings"
)

type Or struct {
	Types []Node
}

func NewOr(types []Node) *Or {
	slices.SortFunc(types, func(a, b Node) int {
		return strings.Compare(a.String(), b.String())
	})
	return &Or{
		Types: types,
	}
}

func (o *Or) Combine(n Node) (Node, error) {
	switch t := n.(type) {
	case *Or:
		// Combine them, check for duplicates.
		u := make([]Node, 0, len(o.Types)+len(t.Types))
		u = append(u, o.Types...)
		for _, v := range t.Types {
			for _, w := range u {
				if w.Equals(v) {
					break
				}
			}
			u = append(u, v)
		}
		return NewOr(u), nil
	default:
		for i, v := range o.Types {
			if v.Equals(n) {
				return o, nil
			}
			if n, err := v.Combine(n); err == nil {
				u := make([]Node, len(o.Types))
				copy(u, o.Types)
				u[i] = n
				return NewOr(u), nil
			}
		}
		u := make([]Node, len(o.Types)+1)
		copy(u, o.Types)
		u[len(o.Types)] = n
		return NewOr(u), nil
	}
}

func (o *Or) Equals(n Node) bool {
	t, ok := n.(*Or)
	if !ok {
		return false
	}
	if len(o.Types) != len(t.Types) {
		return false
	}
	for i, v := range o.Types {
		if !v.Equals(t.Types[i]) {
			return false
		}
	}
	return true
}
