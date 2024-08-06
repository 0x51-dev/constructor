package constructor

import (
	"testing"
)

func TestAny_Combine(t *testing.T) {
	a := new(Any)
	t.Run("string", func(t *testing.T) {
		n, err := a.Combine(new(String))
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := n.(*String); !ok {
			t.Error("unexpected type")
		}
	})

	for _, n := range []Node{new(Any), new(Array), new(Optional), new(Or), new(String), new(Struct)} {
		if _, err := a.Combine(n); err != nil {
			t.Error("expected error")
		}
	}
}
