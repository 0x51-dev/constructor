package constructor

import (
	"fmt"
	"strings"
)

func (*Any) String() string {
	return "any"
}

func (a *Array) String() string {
	return fmt.Sprintf("[]%s", a.Type)
}

func (*Number) String() string {
	return "number"
}

func (o *Optional) String() string {
	return fmt.Sprintf("*%s", o.Type)
}

func (o *Or) String() string {
	var types []string
	for _, t := range o.Types {
		types = append(types, t.String())
	}
	return fmt.Sprintf("(%s)", strings.Join(types, ", "))
}

func (*String) String() string {
	return "string"
}

func (s *Struct) String() string {
	var fields []string
	for _, k := range s.SortedKeys {
		fields = append(fields, fmt.Sprintf("%s %s", k, s.Types[k]))
	}
	return fmt.Sprintf("struct{%s}", strings.Join(fields, "; "))
}
