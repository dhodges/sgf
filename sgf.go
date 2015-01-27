package sgf

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

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

type GameInfo map[string]string

func (gi GameInfo) SortedKeys() []string {
	var keys sort.StringSlice
	for k, _ := range gi {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

func (gi GameInfo) String() string {
	str := ""
	for _, k := range gi.SortedKeys() {
		str += k + "[" + gi[k] + "]"
	}
	return ";" + str
}

type SGFGame struct {
	GameInfo GameInfo
	GameTree *Node
	Errors   []error
}

func (sgf *SGFGame) AddInfo(prop Property) {
	sgf.GameInfo[strings.ToUpper(prop.Name)] = prop.Value
}

func (sgf *SGFGame) GetInfo(name string) (value string, ok bool) {
	value, ok = sgf.GameInfo[strings.ToUpper(name)]
	return value, ok
}

func (sgf SGFGame) GameTreeString() string {
	treeString := ""
	for node := sgf.GameTree; node != nil; node = node.Next {
		treeString += node.String()
	}
	return treeString
}

func (sgf SGFGame) String() string {
	return "(" + sgf.GameInfo.String() + sgf.GameTreeString() + ")"
}

func (sgf SGFGame) NodeCount() int {
	count := 0
	for node := sgf.GameTree; node != nil; node = node.Next {
		count += 1
	}
	return count
}

func (sgf SGFGame) NthNode(n int) (node *Node, err error) {
	if n < 1 {
		return nil, errors.New("n less than 1")
	}
	nodeCount := sgf.NodeCount()

	if n > nodeCount {
		return nil, errors.New(fmt.Sprintf("n greater than node count (%d)", nodeCount))
	}
	for node = sgf.GameTree; n > 1; n -= 1 {
		node = node.Next
	}
	return node, nil
}

func (sgf *SGFGame) AddError(msg string) {
	sgf.Errors = append(sgf.Errors, errors.New(msg))
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
