package parse

import "github.com/dhodges/sgf"

func Parse(input string) (games []*sgf.SGFGame) {
	var currentNode *sgf.Node
	var game *sgf.SGFGame
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
				game = new(sgf.SGFGame)
				game.GameInfo = make(sgf.GameInfo)
				games = append(games, game)
				parsingSetup = true
			} else if parsingSetup {
				game.AddError(l.QuoteErrorContext("unexpected left parenthesis"))
				break Loop
			} else {
				if len(game.GameInfo) == 0 {
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
				if len(game.GameInfo) > 0 {
					parsingSetup = false
					parsingGametree = true
					game.GameTree = new(sgf.Node)
					currentNode = game.GameTree
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
