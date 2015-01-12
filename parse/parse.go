package parse

import (
	"errors"
	"fmt"
)

type PlayerInfo struct {
	name string
	rank string
	team string
}

type GameInfo struct {
	black             PlayerInfo
	white             PlayerInfo
	result            string
	komi              string
	handicap          string
	timeLimits        string
	date              string
	event             string
	round             string
	place             string
	rules             string
	gameName          string
	opening           string
	overtime          string
	gameInfo          string
	comment           string
	source            string
	user              string
	annotator         string
	annotation        string
	copyright         string
	charset           string
	boardsize         string
	unknownProperties []Property
}

type Property struct {
	name  string
	value string
}

func (p Property) String() string {
	return fmt.Sprintf("%s[%s]", p.name, p.value)
}

type Point struct {
	x rune
	y rune
}

type Node struct {
	point      Property
	properties []Property
}

func (point *Point) String() string {
	return fmt.Sprintf("[%c%c]", point.x, point.y)
}

type GameTree struct {
	nodes []Node
}

func (node *Node) AddProperty(prop Property) {
	switch prop.name {
	case "B", "W":
		node.point = prop
	default:
		node.properties = append(node.properties, prop)
	}
}

func (gt *GameTree) NewNode() *Node {
	gt.nodes = append(gt.nodes, *new(Node))
	return &gt.nodes[len(gt.nodes)-1]
}

type SGFGame struct {
	gameInfo GameInfo
	gameTree GameTree
	errors   []error
}

func (gi *GameInfo) AddProperty(prop Property) {
	value := prop.value
	switch prop.name {
	case "PB":
		gi.black.name = value
	case "BR":
		gi.black.rank = value
	case "BT":
		gi.black.team = value
	case "PW":
		gi.white.name = value
	case "WR":
		gi.white.rank = value
	case "WT":
		gi.white.team = value
	case "AN":
		gi.annotator = value
	case "CP":
		gi.copyright = value
	case "DT":
		gi.date = value
	case "EV":
		gi.event = value
	case "GC":
		gi.gameInfo = value
	case "C":
		gi.comment = value
	case "GN":
		gi.gameName = value
	case "HA":
		gi.handicap = value
	case "ON":
		gi.opening = value
	case "OT":
		gi.overtime = value
	case "PC":
		gi.place = value
	case "RE":
		gi.result = value
	case "RO":
		gi.round = value
	case "RU":
		gi.rules = value
	case "SO":
		gi.source = value
	case "TM":
		gi.timeLimits = value
	case "US":
		gi.user = value
	case "CA":
		gi.charset = value
	case "SZ":
		gi.boardsize = value
	case "KM":
		gi.komi = value
	default:
		gi.unknownProperties = append(gi.unknownProperties, prop)
	}
}

func (sgf *SGFGame) Parse(input string) *SGFGame {
	var currentNode *Node
	l := lex(input)
	prop := Property{}
	parsingSetup := false
	parsingGame := false
Loop:
	for {
		i := l.nextItem()
		switch i.typ {
		case itemError:
			sgf.errors = append(sgf.errors, errors.New(i.val))
			break Loop
		case itemRightParen, itemEOF:
			break Loop
		case itemSemiColon:
			if !parsingSetup && !parsingGame {
				parsingSetup = true
			} else {
				currentNode = sgf.gameTree.NewNode()
				if parsingSetup && !parsingGame {
					parsingSetup = false
					parsingGame = true
				}
			}
		case itemPropertyName:
			prop = Property{i.val, ""}
		case itemPropertyValue:
			prop.value = i.val
			if parsingSetup {
				sgf.gameInfo.AddProperty(prop)
			} else {
				currentNode.AddProperty(prop)
			}
		}
	}
	return sgf
}
