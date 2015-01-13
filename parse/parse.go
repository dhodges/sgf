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
	next       *Node
}

func (point *Point) String() string {
	return fmt.Sprintf("[%c%c]", point.x, point.y)
}

func (node *Node) AddProperty(prop Property) {
	switch prop.name {
	case "B", "W":
		node.point = prop
	default:
		node.properties = append(node.properties, prop)
	}
}

func (node *Node) NewNode() *Node {
	node.next = new(Node)
	return node.next
}

type SGFGame struct {
	gameInfo GameInfo
	gameTree *Node
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

func (sgf *SGFGame) NodeCount() int {
	count := 0
	for node := sgf.gameTree; node != nil; node = node.next {
		count += 1
	}
	return count
}

func (sgf *SGFGame) AddError(msg string) {
	sgf.errors = append(sgf.errors, errors.New(msg))
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
			sgf.AddError(i.val)
			break Loop
		case itemRightParen, itemEOF:
			break Loop
		case itemSemiColon:
			if !parsingSetup && !parsingGame {
				parsingSetup = true
			} else {
				if parsingSetup && !parsingGame {
					parsingSetup = false
					parsingGame = true
					sgf.gameTree = new(Node)
					currentNode = sgf.gameTree
				} else {
					currentNode = currentNode.NewNode()
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
