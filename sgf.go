package sgf

import "fmt"

type Property struct {
	Name  string
	Value string
}

func (p Property) String() string {
	return fmt.Sprintf("%s[%s]", p.Name, p.Value)
}

type Point struct {
	X rune
	Y rune
}

func (point Point) String() string {
	return fmt.Sprintf("[%c%c]", point.X, point.Y)
}

type Node struct {
	Point      Property
	Properties []Property
	Variations []*Node
	Next       *Node
}

func (node Node) variationString() string {
	if len(node.Variations) == 0 {
		return ""
	}

	str := ""
	for _, nodevar := range node.Variations {
		nodestr := ""
		for nptr := nodevar; nptr != nil; nptr = nptr.Next {
			nodestr += nptr.String()
		}
		str += "(" + nodestr + ")"
	}
	return str
}

func (node Node) propertiesString() string {
	str := ""
	for _, prop := range node.Properties {
		str += prop.String()
	}
	return str
}

func (node Node) String() string {
	return ";" +
		node.Point.String() +
		node.propertiesString() +
		node.variationString()
}

func (node *Node) AddProperty(prop Property) {
	switch prop.Name {
	case "B", "W":
		node.Point = prop
	default:
		node.Properties = append(node.Properties, prop)
	}
}

func (node *Node) NewNode() *Node {
	node.Next = new(Node)
	return node.Next
}

func (n *Node) NewVariation() *Node {
	node := new(Node)
	n.Variations = append(n.Variations, node)
	return node
}

const BlackPlayerName = "PB"
const BlackPlayerRank = "BR"
const BlackPlayerTeam = "BT"
const WhitePlayerName = "PW"
const WhitePlayerRank = "WR"
const WhitePlayerTeam = "WT"
const Annotator = "AN"
const Copyright = "CP"
const Date = "DT"
const Event = "EV"
const GameComment = "GC"
const Comment = "C"
const GameName = "GN"
const Handicap = "HA"
const Opening = "ON"
const Overtime = "OT"
const Place = "PC"
const Result = "RE"
const Round = "RO"
const Rules = "RU"
const Source = "SO"
const TimeLimits = "TM"
const User = "US"
const Charset = "CA"
const Boardsize = "SZ"
const Komi = "KM"
