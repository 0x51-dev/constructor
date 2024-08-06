package golang

import (
	"fmt"
	. "github.com/0x51-dev/constructor"
	"strings"
)

func NodeToGo(name string, n Node) (string, error) {
	str, err := nodeToGo(n)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("type %s %s", name, str), nil
}

func nodeToGo(n Node) (string, error) {
	switch n := n.(type) {
	case nil:
		return "null", nil
	case *Any:
		return "any", nil
	case *Array:
		t, err := nodeToGo(n.Type)
		if err != nil {
			return "", err
		}
		return "[]" + t, nil
	case *Number:
		return "json.Number", nil
	case *Optional:
		t, err := nodeToGo(n.Type)
		if err != nil {
			return "", err
		}
		return "*" + t, nil
	case *String:
		return "string", nil
	case *Struct:
		fields := ""
		for _, k := range n.SortedKeys {
			t, err := nodeToGo(n.Types[k])
			if err != nil {
				return "", err
			}
			fields += fmt.Sprintf("%s %s `json:\"%s\"`\n", strings.Title(k), t, k)
		}
		return "struct {\n" + fields + "}", nil
	case *Or:
		types := ""
		for _, t := range n.Types {
			t, err := nodeToGo(t)
			if err != nil {
				return "", err
			}
			types += t + ", "
		}
		return "any /* " + types + " */", nil
	default:
		return "", fmt.Errorf("unexpected type %T", n)
	}
}
