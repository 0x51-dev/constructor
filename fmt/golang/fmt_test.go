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
			"first_name": "Alice",
			"username": "alice"
		},
		{
			"first_name": "Bob",
			"username": "bob",
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
	// FirstName string `json:"first_name"`
	// Username string `json:"username"`
	// } `json:"users"`
	// }
}
