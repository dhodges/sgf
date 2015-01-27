package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dhodges/sgf"
)

type SGFGame struct {
	gameInfo sgf.GameInfo
	gameTree *sgf.Node
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
	for node := sgf.gameTree; node != nil; node = node.Next {
		treeString += node.String()
	}
	return treeString
}

func (sgf SGFGame) String() string {
	return "(" + sgf.gameInfo.String() + sgf.GameTreeString() + ")"
}

func (sgf SGFGame) NodeCount() int {
	count := 0
	for node := sgf.gameTree; node != nil; node = node.Next {
		count += 1
	}
	return count
}

func (sgf SGFGame) NthNode(n int) (node *sgf.Node, err error) {
	if n < 1 {
		return nil, errors.New("n less than 1")
	}
	nodeCount := sgf.NodeCount()

	if n > nodeCount {
		return nil, errors.New(fmt.Sprintf("n greater than node count (%d)", nodeCount))
	}
	for node = sgf.gameTree; n > 1; n -= 1 {
		node = node.Next
	}
	return node, nil
}

func (sgf *SGFGame) AddError(msg string) {
	sgf.errors = append(sgf.errors, errors.New(msg))
}

func Parse(input string) (games []*SGFGame) {
	var currentNode *sgf.Node
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
				game.gameInfo = make(sgf.GameInfo)
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
				currentNode = node.(*sgf.Node)
			} else {
				parsingSetup = false
				parsingGametree = false
			}
		case itemSemiColon:
			if parsingSetup {
				if len(game.gameInfo) > 0 {
					parsingSetup = false
					parsingGametree = true
					game.gameTree = new(sgf.Node)
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
