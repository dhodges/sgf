package sgf

import (
	"errors"
	"fmt"
	"strings"
)

type Game struct {
	GameInfo GameInfo
	GameTree *Node
	Errors   []error
}

func (sgf *Game) AddInfo(prop Property) {
	sgf.GameInfo[strings.ToUpper(prop.Name)] = prop.Value
}

func (sgf *Game) GetInfo(name string) (value string, ok bool) {
	value, ok = sgf.GameInfo[strings.ToUpper(name)]
	return value, ok
}

func (sgf Game) GameTreeString() string {
	treeString := ""
	for node := sgf.GameTree; node != nil; node = node.Next {
		treeString += node.String()
	}
	return treeString
}

func (sgf Game) String() string {
	return "(" + sgf.GameInfo.String() + sgf.GameTreeString() + ")"
}

func (sgf Game) NodeCount() int {
	count := 0
	for node := sgf.GameTree; node != nil; node = node.Next {
		count += 1
	}
	return count
}

func (sgf Game) NthNode(n int) (node *Node, err error) {
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

func (sgf *Game) AddError(msg string) {
	sgf.Errors = append(sgf.Errors, errors.New(msg))
}
