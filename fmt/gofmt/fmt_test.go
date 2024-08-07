package gofmt_test

import (
	"bytes"
	"fmt"
	"github.com/0x51-dev/constructor"
	"github.com/0x51-dev/constructor/fmt/gofmt"
)

func ExampleNodeToGo() {
	raw := `{
	"users": [
		{
			"first_name": "Alice",
			"username": "alice",
			"age": "23"
		},
		{
			"username": "bob",
			"age": 42
		}
	]
}`
	n, err := constructor.Construct(bytes.NewReader([]byte(raw)))
	if err != nil {
		panic(err)
	}
	str, err := gofmt.NodeToGo("Users", n)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	// Output:
	// type Users struct {
	// 	Users []struct {
	// 		Age any /* json.Number, string */ `json:"age"`
	// 		FirstName string `json:"first_name"`
	// 		Username string `json:"username"`
	// 	} `json:"users"`
	// }
}
