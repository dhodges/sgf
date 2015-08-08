package sgf

import "fmt"

type Point struct {
	X rune
	Y rune
}

func (point Point) String() string {
	return fmt.Sprintf("[%c%c]", point.X, point.Y)
}
