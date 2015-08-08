package sgf

import "fmt"

type Property struct {
	Name  string
	Value string
}

func (p Property) String() string {
	return fmt.Sprintf("%s[%s]", p.Name, p.Value)
}
