package constructor

import (
	"testing"
)

func TestString_Combine(t *testing.T) {
	s := new(String)

	t.Run("string", func(t *testing.T) {
		n, err := s.Combine(new(String))
		if err != nil {
			t.Fatal(err)
		}
		if !n.Equals(s) {
			t.Errorf("expected %v, got %v", s, n)
		}
	})

	t.Run("any", func(t *testing.T) {
		switch n, _ := s.Combine(new(Any)); n.(type) {
		case *String:
		default:
			t.Errorf("unexpected type %T", n)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		for _, n := range []Node{new(Array), new(Optional), new(Or), new(Struct)} {
			if _, err := s.Combine(n); err == nil {
				t.Error("expected error")
			}
		}
	})
}
