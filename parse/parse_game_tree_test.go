package parse

import (
	"fmt"
	"testing"
)

func errorUnlessStrEqual(t *testing.T, found string, expected string, errMsg string) {
	if found != expected {
		t.Error(fmt.Sprintf("%s, found: '%s', expected: '%s'", errMsg, found, expected))
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
	sgf, err := parseString(gameTreeString)
	if err != nil {
		t.Error(err)
		return
	}

	if sgf.NodeCount() != 6 {
		t.Errorf("game tree is incorrect, found %d nodes, expected 6", sgf.NodeCount())
	}

	node := sgf.gameTree
	errorUnlessStrEqual(t, node.point.String(), "B[qc]", "1st node move is incorrect")

	node = node.next
	errorUnlessStrEqual(t, node.point.String(), "W[cd]", "2nd node move is incorrect")

	node = node.next
	errorUnlessStrEqual(t, node.point.String(), "B[dp]", "3rd node move is incorrect")

	node = node.next
	errorUnlessStrEqual(t, node.point.String(), "W[pq]", "4th node move is incorrect")

	node = node.next
	errorUnlessStrEqual(t, node.point.String(), "B[jj]", "5th node move is incorrect")

	node = node.next
	errorUnlessStrEqual(t, node.point.String(), "W[pd]", "6th node move is incorrect")
}

func TestParsingFullGameTree(t *testing.T) {
	sgf, err := parseFixture("2014.07.06_WAGC-Rd1-Lithuania-Canada-var.sgf")
	if err != nil {
		t.Error(err)
		return
	}

	nodeCount := sgf.NodeCount()

	if nodeCount != 7 {
		t.Error(fmt.Printf("wrong number of game nodes (expected 7, found %d)\n", nodeCount))
	}

	node, err := sgf.NthNode(7)
	if err != nil {
		t.Error(err)
		return
	}

	if node.point.String() != "B[hd]" {
		t.Error(fmt.Sprintf("wrong node: expected 'B[hd]', found '%s'", node.point))
	}

	if len(node.variations) != 3 {
		t.Error(fmt.Sprintf("wrong number of variations: expected 3, found '%d'", len(node.variations)))
	}
}
