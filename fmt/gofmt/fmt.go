package gofmt

import (
	"fmt"
	. "github.com/0x51-dev/constructor"
	"strings"
)

func NodeToGo(name string, n Node) (string, error) {
	str, err := nodeToGo(n, 0)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("type %s %s", name, str), nil
}

func nodeToGo(n Node, indent int) (string, error) {
	sw := new(strings.Builder)
	switch n := n.(type) {
	case nil:
		sw.WriteString("null")
	case *Any:
		sw.WriteString("any")
	case *Array:
		sw.WriteString("[]")
		t, err := nodeToGo(n.Type, indent)
		if err != nil {
			return "", err
		}
		sw.WriteString(t)
	case *Number:
		sw.WriteString("json.Number")
	case *Optional:
		sw.WriteString("*")
		t, err := nodeToGo(n.Type, indent)
		if err != nil {
			return "", err
		}
		sw.WriteString(t)
	case *String:
		sw.WriteString("string")
	case *Struct:
		sw.WriteString("struct {\n")
		for _, k := range n.SortedKeys {
			t, err := nodeToGo(n.Types[k], indent+1)
			if err != nil {
				return "", err
			}
			for i := 0; i <= indent; i++ {
				sw.WriteString("\t")
			}
			sw.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", sanitiseName(k), t, k))
		}
		for i := 0; i < indent; i++ {
			sw.WriteString("\t")
		}
		sw.WriteString("}")
	case *Or:
		sw.WriteString("any /* ")
		for i, t := range n.Types {
			t, err := nodeToGo(t, indent)
			if err != nil {
				return "", err
			}
			if i != 0 {
				sw.WriteString(", ")
			}
			sw.WriteString(t)
		}
		sw.WriteString(" */")
	default:
		return "", fmt.Errorf("unexpected type %T", n)
	}
	return sw.String(), nil
}

func sanitiseName(s string) string {
	parts := strings.Split(s, "_")
	for i, p := range parts {
		parts[i] = fmt.Sprintf("%s%s", strings.ToUpper(p[:1]), p[1:])
	}
	return strings.Join(parts, "")
}
