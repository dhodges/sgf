package parse

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func errorUnlessStrEqual(t *testing.T, found string, expected string, errMsg string) {
	if found != expected {
		t.Error(fmt.Sprintf("%s, found: '%s', expected: '%s'", errMsg, found, expected))
	}
}

var gameTreeFixture = "(" +
	";PB[Lee Sedol]BR[6p]BT[South Korea]" +
	"PW[Gu Li]WR[9p]WT[China]RE[B+2]" +
	";B[qc]C[At that time, Go Seigen was just 20 years' old...]" +
	";W[cd]" +
	";B[dp]" +
	";W[pq]" +
	";B[jj]C[In this opening, Black takes two corners with B1 and B3...]" +
	";W[pd]C[W6 is the most powerful tactic when your opponent has a stone at 3-3.]" +
	")"

func TestParsingGameTree(t *testing.T) {
	gt := new(SGFGame).Parse(gameTreeFixture).gameTree

	if len(gt.nodes) != 6 {
		t.Error("game tree is incorrect")
	}

	errorUnlessStrEqual(t, gt.nodes[0].point.String(), "B[qc]", "1st node move is incorrect")
	errorUnlessStrEqual(t, gt.nodes[1].point.String(), "W[cd]", "2nd node move is incorrect")
	errorUnlessStrEqual(t, gt.nodes[2].point.String(), "B[dp]", "3rd node move is incorrect")
	errorUnlessStrEqual(t, gt.nodes[3].point.String(), "W[pq]", "4th node move is incorrect")
	errorUnlessStrEqual(t, gt.nodes[4].point.String(), "B[jj]", "5th node move is incorrect")
	errorUnlessStrEqual(t, gt.nodes[5].point.String(), "W[pd]", "6th node move is incorrect")
}

func TestParsingFullGameTree(t *testing.T) {
	buf, err := ioutil.ReadFile("../fixtures/sgf_files/19331016-Honinbo_Shusai-Go_Seigen.sgf")
	if err != nil {
		t.Error(err.Error)
		return
	}

	fixture := strings.Replace(string(buf), "\n", "", -1)
	fixture = strings.Replace(string(buf), "\r", "", -1)

	gt := new(SGFGame).Parse(fixture).gameTree

	nodeCount := len(gt.nodes)

	if nodeCount != 252 {
		t.Error(fmt.Printf("wrong number of game nodes (found %d, expected 252)\n", nodeCount))
	}
}
