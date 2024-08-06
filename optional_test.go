package constructor

import (
	"testing"
)

func TestOptional_Combine(t *testing.T) {
	o := &Optional{Type: new(Any)}

	t.Run("identical", func(t *testing.T) {
		c, err := o.Combine(o)
		if err != nil {
			t.Fatal(err)
		}
		if !c.Equals(o) {
			t.Errorf("expected %v, got %v", o, c)
		}
	})

	t.Run("string", func(t *testing.T) {
		a := &Optional{Type: new(String)}
		c, err := o.Combine(a)
		if err != nil {
			t.Fatal(err)
		}
		if !c.Equals(a) {
			t.Errorf("expected %v, got %v", a, c)
		}
	})
}
