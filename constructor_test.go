package constructor

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleConstruct_nullable() {
	n, _ := Construct(strings.NewReader(`[ { "nullable": null }, { "nullable": "null" } ]`))
	fmt.Println(n)
	// Output:
	// []struct{nullable *string}
}

func TestConstruct(t *testing.T) {
	t.Run("nullable", func(t *testing.T) {
		raw := `[ { "nullable": null }, { "nullable": "null" } ]`
		n, err := Construct(strings.NewReader(raw))
		if err != nil {
			t.Fatal(err)
		}
		expected := &Array{Type: NewStruct(map[string]Node{
			"nullable": &Optional{Type: new(String)},
		})}
		if !n.Equals(expected) {
			t.Errorf("expected %v, got %v", expected, n)
		}
	})
}
