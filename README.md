# Constructor

Provides a quick and dirty way to convert JSON data dumps to a Golang struct.

## Usage

Given the following JSON data dump:

```json
{
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
}
```

You can convert it to a Golang struct like this:

```go
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/0x51-dev/constructor"
	"github.com/0x51-dev/constructor/fmt/gofmt"
)

func main() {
	data, _ := os.ReadFile("data.json")
	n, _ := constructor.Construct(bytes.NewReader(data))
	str, _ := gofmt.NodeToGo("Users", n)
	fmt.Println(str) // Prints out formatted Golang struct.
}

```

Output:

```go
package users

type Users struct {
	Users []struct {
		Age       any /* json.Number, string */ `json:"age"`
		FirstName *string                       `json:"first_name"`
		Username  string                        `json:"username"`
	} `json:"users"`
}

```
