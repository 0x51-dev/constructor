package constructor

import (
	"encoding/json"
	"fmt"
	"io"
)

func Construct(r io.Reader) (Node, error) {
	var v any
	dec := json.NewDecoder(r)
	dec.UseNumber()
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	return constructAny(v)
}

func constructAny(v any) (Node, error) {
	switch v := v.(type) {
	case nil:
		return &Optional{
			Type: new(Any),
		}, nil
	case string:
		return new(String), nil
	case json.Number:
		return new(Number), nil
	case map[string]any:
		types := make(map[string]Node)
		for k, v := range v {
			v, err := constructAny(v)
			if err != nil {
				return nil, err
			}
			types[k] = v
		}
		return NewStruct(types), nil
	case []any:
		var a Array
		for _, v := range v {
			v, err := constructAny(v)
			if err != nil {
				return nil, err
			}
			if a.Type != nil {
				if !a.Type.Equals(v) {
					t, err := a.Type.Combine(v)
					if err != nil {
						return nil, err
					}
					a.Type = t
				}
			} else {
				a.Type = v
			}
		}
		return &a, nil
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	}
}
