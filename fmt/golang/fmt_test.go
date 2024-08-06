package golang_test

import (
	"bytes"
	"fmt"
	"github.com/0x51-dev/constructor"
	"github.com/0x51-dev/constructor/fmt/golang"
)

func ExampleNodeToGo() {
	raw := `{
	"users": [
		{
			"name": "Alice"
		},
		{
			"name": "Bob",
			"age": 42
		}
	]
}`
	n, err := constructor.Construct(bytes.NewReader([]byte(raw)))
	if err != nil {
		panic(err)
	}
	str, err := golang.NodeToGo("Users", n)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	// Output:
	// type Users struct {
	// Users []struct {
	// Age json.Number `json:"age"`
	// Name string `json:"name"`
	// } `json:"users"`
	// }
}
