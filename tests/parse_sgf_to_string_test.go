package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/sgf"
  "github.com/stretchr/testify/assert"
)

func TestGameInfoToString(t *testing.T) {
	game := new(sgf.Game)
	game.GameInfo = make(sgf.GameInfo)
	game.AddInfo(sgf.Property{Name: sgf.BlackPlayerName, Value: "Lee Sedol"})
	game.AddInfo(sgf.Property{Name: sgf.BlackPlayerRank, Value: "9p"})
	game.AddInfo(sgf.Property{Name: sgf.BlackPlayerTeam, Value: "South Korea"})
	game.AddInfo(sgf.Property{Name: sgf.WhitePlayerName, Value: "Gu Li"})
	game.AddInfo(sgf.Property{Name: sgf.WhitePlayerRank, Value: "9p"})
	game.AddInfo(sgf.Property{Name: sgf.WhitePlayerTeam, Value: "China"})
	game.AddInfo(sgf.Property{Name: sgf.Result, Value: "B+2"})
	game.AddInfo(sgf.Property{Name: sgf.Charset, Value: "UTF-8"})
	game.AddInfo(sgf.Property{Name: sgf.Annotator, Value: "bob"})
	game.AddInfo(sgf.Property{Name: sgf.Copyright, Value: "Copyright"})
	game.AddInfo(sgf.Property{Name: sgf.Date, Value: "2014-12-25,26"})
	game.AddInfo(sgf.Property{Name: sgf.Handicap, Value: "4"})
	game.AddInfo(sgf.Property{Name: sgf.Event, Value: "Pewter Cup"})
	game.AddInfo(sgf.Property{Name: sgf.GameName, Value: "sally"})
	game.AddInfo(sgf.Property{Name: sgf.GameComment, Value: "it was long"})
	game.AddInfo(sgf.Property{Name: sgf.Opening, Value: "low Chinese"})
	game.AddInfo(sgf.Property{Name: sgf.Overtime, Value: "byo yomi"})

	expected := "(;AN[bob]BR[9p]BT[South Korea]CA[UTF-8]CP[Copyright]DT[2014-12-25,26]EV[Pewter Cup]GC[it was long]GN[sally]HA[4]ON[low Chinese]OT[byo yomi]PB[Lee Sedol]PW[Gu Li]RE[B+2]WR[9p]WT[China])"
	assert.Equal(t, game.String(), expected, "invalid string")
}

func TestSGFtoString(t *testing.T) {
	variant_sgf_1 := "(;GM[1];B[dp];W[pd];B[cd])"

	games, err := parseString(variant_sgf_1)
	assert.Equal(t, err, nil, "problem parsing game string")

	game := games[0]
	assert.Equal(t, game.String(), variant_sgf_1, "error writing SGF to string")
}

func TestVariationsToString(t *testing.T) {
	gameStr := "(;GM[1];B[dp]" +
		"(;W[ef];B[cf](;W[fc];B[dc](;W[pf];B[qk])(;W[rl];B[qg];W[sl]));W[eh];B[ec])" +
		";W[ee])"

	games, err := parseString(gameStr)
	assert.Equal(t, err, nil, "problem parsing game string")

	game := games[0]
	assert.Equal(t, game.String(), gameStr, "error writing SGF to string")
}
