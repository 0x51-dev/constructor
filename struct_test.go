package constructor

import "testing"

func TestStruct_Combine(t *testing.T) {
	t.Run("complex", func(t *testing.T) {
		left := NewStruct(map[string]Node{
			"first_name": new(String),
			"username":   new(String),
			"age":        new(String),
		})
		right := NewStruct(map[string]Node{
			"username": new(String),
			"age":      new(Number),
		})
		expected := NewStruct(map[string]Node{
			"first_name": &Optional{Type: new(String)},
			"username":   new(String),
			"age":        NewOr([]Node{new(Number), new(String)}),
		})
		actual, err := left.Combine(right)
		if err != nil {
			t.Fatal(err)
		}
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
		inverse, err := right.Combine(left)
		if err != nil {
			t.Fatal(err)
		}
		if !expected.Equals(inverse) {
			t.Errorf("expected %v, got %v", expected, inverse)
		}
	})
}
