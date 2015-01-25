package parse

import "testing"

func TestGameInfoToString(t *testing.T) {
	sgf := new(SGFGame)
	sgf.gameInfo = make(GameInfo)
	sgf.AddInfo(Property{BlackPlayerName, "Lee Sedol"})
	sgf.AddInfo(Property{BlackPlayerRank, "9p"})
	sgf.AddInfo(Property{BlackPlayerTeam, "South Korea"})
	sgf.AddInfo(Property{WhitePlayerName, "Gu Li"})
	sgf.AddInfo(Property{WhitePlayerRank, "9p"})
	sgf.AddInfo(Property{WhitePlayerTeam, "China"})
	sgf.AddInfo(Property{Result, "B+2"})
	sgf.AddInfo(Property{Charset, "UTF-8"})
	sgf.AddInfo(Property{Annotator, "bob"})
	sgf.AddInfo(Property{Copyright, "Copyright"})
	sgf.AddInfo(Property{Date, "2014-12-25,26"})
	sgf.AddInfo(Property{Handicap, "4"})
	sgf.AddInfo(Property{Event, "Pewter Cup"})
	sgf.AddInfo(Property{GameName, "sally"})
	sgf.AddInfo(Property{GameComment, "it was long"})
	sgf.AddInfo(Property{Opening, "low Chinese"})
	sgf.AddInfo(Property{Overtime, "byo yomi"})

	expected := "(;AN[bob]BR[9p]BT[South Korea]CA[UTF-8]CP[Copyright]DT[2014-12-25,26]EV[Pewter Cup]GC[it was long]GN[sally]HA[4]ON[low Chinese]OT[byo yomi]PB[Lee Sedol]PW[Gu Li]RE[B+2]WR[9p]WT[China])"
	if sgf.String() != expected {
		t.Errorf("invalid string. Found: '%s', expected: '%s'", sgf.String(), expected)
	}
}

func TestSGFtoString(t *testing.T) {
	variant_sgf_1 := "(;GM[1];B[dp];W[pd];B[cd])"

	games, err := parseString(variant_sgf_1)
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	if game.String() != variant_sgf_1 {
		t.Errorf("error writing SGF to string. \n   found: %q \nexpected: %q", game.String(), variant_sgf_1)
	}
}

func TestVariationsToString(t *testing.T) {
	gameStr := "(;GM[1];B[dp]" +
		"(;W[ef];B[cf](;W[fc];B[dc](;W[pf];B[qk])(;W[rl];B[qg];W[sl]));W[eh];B[ec])" +
		";W[ee])"

	games, err := parseString(gameStr)
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	sgfStr := game.String()
	if sgfStr != gameStr {
		t.Errorf("error writing SGF to string. \n   Found: %q, \nExpected: %q", sgfStr, gameStr)
	}
}
