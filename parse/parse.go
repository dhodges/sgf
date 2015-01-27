package parse

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/dhodges/sgf"
)

type Point struct {
	x rune
	y rune
}

func (point Point) String() string {
	return fmt.Sprintf("[%c%c]", point.x, point.y)
}

type Node struct {
	point      sgf.Property
	properties []sgf.Property
	variations []*Node
	next       *Node
}

func (node Node) variationString() string {
	if len(node.variations) == 0 {
		return ""
	}

	str := ""
	for _, nodevar := range node.variations {
		nodestr := ""
		for nptr := nodevar; nptr != nil; nptr = nptr.next {
			nodestr += nptr.String()
		}
		str += "(" + nodestr + ")"
	}
	return str
}

func (node Node) propertiesString() string {
	str := ""
	for _, prop := range node.properties {
		str += prop.String()
	}
	return str
}

func (node Node) String() string {
	return ";" +
		node.point.String() +
		node.propertiesString() +
		node.variationString()
}

func (node *Node) AddProperty(prop sgf.Property) {
	switch prop.Name {
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

func (n *Node) NewVariation() *Node {
	node := new(Node)
	n.variations = append(n.variations, node)
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
	gameInfo GameInfo
	gameTree *Node
	errors   []error
}

func (sgf *SGFGame) AddInfo(prop sgf.Property) {
	sgf.gameInfo[strings.ToUpper(prop.Name)] = prop.Value
}

func (sgf *SGFGame) GetInfo(name string) (value string, ok bool) {
	value, ok = sgf.gameInfo[strings.ToUpper(name)]
	return value, ok
}

func (sgf SGFGame) GameTreeString() string {
	treeString := ""
	for node := sgf.gameTree; node != nil; node = node.next {
		treeString += node.String()
	}
	return treeString
}

func (sgf SGFGame) String() string {
	return "(" + sgf.gameInfo.String() + sgf.GameTreeString() + ")"
}

func (sgf SGFGame) NodeCount() int {
	count := 0
	for node := sgf.gameTree; node != nil; node = node.next {
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
	for node = sgf.gameTree; n > 1; n -= 1 {
		node = node.next
	}
	return node, nil
}

func (sgf *SGFGame) AddError(msg string) {
	sgf.errors = append(sgf.errors, errors.New(msg))
}

func Parse(input string) (games []*SGFGame) {
	var currentNode *Node
	var game *SGFGame
	l := lex(input)
	prop := sgf.Property{}
	parsingSetup := false
	parsingGametree := false
	nodeStack := new(Stack)

Loop:
	for {
		i := l.nextItem()
		switch i.typ {
		case itemLeftParen:
			if !parsingSetup && !parsingGametree {
				game = new(SGFGame)
				game.gameInfo = make(GameInfo)
				games = append(games, game)
				parsingSetup = true
			} else if parsingSetup {
				game.AddError(l.QuoteErrorContext("unexpected left parenthesis"))
				break Loop
			} else {
				if len(game.gameInfo) == 0 {
					parsingSetup = true
				} else {
					nodeStack.Push(currentNode)
					currentNode = currentNode.NewVariation()
					if l.nextItem().typ != itemSemiColon {
						game.AddError(l.QuoteErrorContext("semi-colon expected here"))
						break Loop
					}
				}
			}
		case itemRightParen:
			node := nodeStack.Pop()
			if node != nil {
				currentNode = node.(*Node)
			} else {
				parsingSetup = false
				parsingGametree = false
			}
		case itemSemiColon:
			if parsingSetup {
				if len(game.gameInfo) > 0 {
					parsingSetup = false
					parsingGametree = true
					game.gameTree = new(Node)
					currentNode = game.gameTree
				}
			} else {
				currentNode = currentNode.NewNode()
			}
		case itemPropertyName:
			prop = sgf.Property{Name: i.val, Value: ""}
		case itemPropertyValue:
			prop.Value = i.val
			if parsingSetup {
				game.AddInfo(prop)
			} else {
				currentNode.AddProperty(prop)
			}
		case itemError:
			game.AddError(i.val)
			break Loop
		case itemEOF:
			break Loop
		}
	}

	return
}
