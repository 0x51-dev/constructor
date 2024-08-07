package constructor

import (
	"fmt"
	"slices"
	"strings"
)

type Struct struct {
	SortedKeys []string
	Types      map[string]Node
}

func NewStruct(types map[string]Node) *Struct {
	keys := make([]string, 0, len(types))
	for k := range types {
		keys = append(keys, k)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return strings.Compare(a, b)
	})
	return &Struct{
		SortedKeys: keys,
		Types:      types,
	}
}

func (s *Struct) Combine(n Node) (Node, error) {
	switch t := n.(type) {
	case *Optional:
		if n, err := s.Combine(t.Type); err == nil {
			return &Optional{
				Type: n,
			}, nil
		}
	case *Struct:
		types := make(map[string]Node)
		for k, v := range s.Types {
			if _, ok := t.Types[k]; ok {
				types[k] = v
			} else {
				types[k] = &Optional{Type: v}
			}
		}
		for k, v := range t.Types {
			if v2, ok := types[k]; ok {
				if n, err := v2.Combine(v); err == nil {
					types[k] = n
				} else {
					types[k] = NewOr([]Node{v, v2})
				}
			} else {
				types[k] = &Optional{Type: v}
			}
		}
		return NewStruct(types), nil
	}
	return nil, fmt.Errorf("cannot combine %T with %T", s, n)
}

func (s *Struct) Equals(n Node) bool {
	t, ok := n.(*Struct)
	if !ok {
		return false
	}
	if len(s.SortedKeys) != len(t.SortedKeys) {
		return false
	}
	for k, v := range s.SortedKeys {
		if v != t.SortedKeys[k] {
			return false
		}
		if !s.Types[v].Equals(t.Types[v]) {
			return false
		}
	}
	return true
}
