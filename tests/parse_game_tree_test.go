package tests

import (
	"testing"

  "github.com/stretchr/testify/assert"
)

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
	assert.Equal(t, err, nil, "problem parsing gametree string")

	game := games[0]
	assert.Equal(t, game.NodeCount(), 6, "game tree is incorrect")

	node := game.GameTree
	assert.Equal(t, node.Point.String(), "B[qc]", "1st node move is incorrect")

	node = node.Next
	assert.Equal(t, node.Point.String(), "W[cd]", "2nd node move is incorrect")

	node = node.Next
	assert.Equal(t, node.Point.String(), "B[dp]", "3rd node move is incorrect")

	node = node.Next
	assert.Equal(t, node.Point.String(), "W[pq]", "4th node move is incorrect")

	node = node.Next
	assert.Equal(t, node.Point.String(), "B[jj]", "5th node move is incorrect")

	node = node.Next
	assert.Equal(t, node.Point.String(), "W[pd]", "6th node move is incorrect")
}

func TestParsingFullGameTree(t *testing.T) {
	games, err := parseFixture("2014.07.06_WAGC-Rd1-Lithuania-Canada-var.sgf")
	assert.Equal(t, err, nil, "problem loading fixture")

	game := games[0]
	assert.Equal(t, game.NodeCount(), 7, "wrong number of game nodes")

	node, err := game.NthNode(7)
	assert.Equal(t, err, nil, "problem getting node")

	assert.Equal(t, node.Point.String(),  "B[hd]", "wrong node")
	assert.Equal(t, len(node.Variations), 3,       "wrong number of variations")
}
