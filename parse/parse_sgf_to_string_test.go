package parse

import (
	"fmt"
	"testing"
)

func TestGameInfoToString(t *testing.T) {
	sgf := new(SGFGame)
	sgf.gameInfo = map[string]string{}
	sgf.AddProperty(Property{BlackPlayerName, "Lee Sedol"})
	sgf.AddProperty(Property{BlackPlayerRank, "9p"})
	sgf.AddProperty(Property{BlackPlayerTeam, "South Korea"})
	sgf.AddProperty(Property{WhitePlayerName, "Gu Li"})
	sgf.AddProperty(Property{WhitePlayerRank, "9p"})
	sgf.AddProperty(Property{WhitePlayerTeam, "China"})
	sgf.AddProperty(Property{Result, "B+2"})
	sgf.AddProperty(Property{Charset, "UTF-8"})
	sgf.AddProperty(Property{Annotator, "bob"})
	sgf.AddProperty(Property{Copyright, "Copyright"})
	sgf.AddProperty(Property{Date, "2014-12-25,26"})
	sgf.AddProperty(Property{Handicap, "4"})
	sgf.AddProperty(Property{Event, "Pewter Cup"})
	sgf.AddProperty(Property{GameName, "sally"})
	sgf.AddProperty(Property{GameComment, "it was long"})
	sgf.AddProperty(Property{Opening, "low Chinese"})
	sgf.AddProperty(Property{Overtime, "byo yomi"})

	expected := "(;AN[bob]BR[9p]BT[South Korea]CA[UTF-8]CP[Copyright]DT[2014-12-25,26]EV[Pewter Cup]GC[it was long]GN[sally]HA[4]ON[low Chinese]OT[byo yomi]PB[Lee Sedol]PW[Gu Li]RE[B+2]WR[9p]WT[China])"
	if sgf.String() != expected {
		t.Error(fmt.Sprintf("invalid string. Found: '%s', expected: '%s'", sgf.String(), expected))
	}
}

func TestSGFtoString(t *testing.T) {
	variant_sgf_1 := "(;GM[1];B[dp];W[pd];B[cd])"

	sgf, err := parseString(variant_sgf_1)
	if err != nil {
		t.Error(err)
		return
	}

	if sgf.String() != variant_sgf_1 {
		t.Error(fmt.Sprintf("error writing SGF to string. \n   found: %q \nexpected: %q", sgf.String(), variant_sgf_1))
	}
}
