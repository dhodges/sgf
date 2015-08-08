package tests

import (
	"fmt"
	"testing"
)

func errorUnlessStrEqual(t *testing.T, found string, expected string, errMsg string) {
	if found != expected {
		t.Errorf("%s, found: '%s', expected: '%s'", errMsg, found, expected)
	}
}

var gameTreeString = "(" +
	";PB[Lee Sedol]BR[6p]BT[South Korea]PW[Gu Li]WR[9p]WT[China]RE[B+2]" +
	";B[qc]C[At that time, Go Seigen was just 20 years' old...]" +
	";W[cd]" +
	";B[dp]" +
	";W[pq]" +
	";B[jj]C[In this opening, Black takes two corners with B1 and B3...]" +
	";W[pd]C[W6 is the most powerful tactic when your opponent has a stone at 3-3.]" +
	")"

func TestParsingGameTree(t *testing.T) {
	games, err := parseString(gameTreeString)
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	if game.NodeCount() != 6 {
		t.Errorf("game tree is incorrect, found %d nodes, expected 6", game.NodeCount())
	}

	node := game.GameTree
	errorUnlessStrEqual(t, node.Point.String(), "B[qc]", "1st node move is incorrect")

	node = node.Next
	errorUnlessStrEqual(t, node.Point.String(), "W[cd]", "2nd node move is incorrect")

	node = node.Next
	errorUnlessStrEqual(t, node.Point.String(), "B[dp]", "3rd node move is incorrect")

	node = node.Next
	errorUnlessStrEqual(t, node.Point.String(), "W[pq]", "4th node move is incorrect")

	node = node.Next
	errorUnlessStrEqual(t, node.Point.String(), "B[jj]", "5th node move is incorrect")

	node = node.Next
	errorUnlessStrEqual(t, node.Point.String(), "W[pd]", "6th node move is incorrect")
}

func TestParsingFullGameTree(t *testing.T) {
	games, err := parseFixture("2014.07.06_WAGC-Rd1-Lithuania-Canada-var.sgf")
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	nodeCount := game.NodeCount()

	if nodeCount != 7 {
		t.Error(fmt.Printf("wrong number of game nodes (expected 7, found %d)\n", nodeCount))
	}

	node, err := game.NthNode(7)
	if err != nil {
		t.Error(err)
		return
	}

	if node.Point.String() != "B[hd]" {
		t.Errorf("wrong node: expected \"B[hd]\", found %q", node.Point)
	}

	if len(node.Variations) != 3 {
		t.Errorf("wrong number of variations: expected 3, found %d", len(node.Variations))
	}
}
