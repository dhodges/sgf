package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/parse"
  "github.com/stretchr/testify/assert"
)

var gameInfoString = "(;" +
	"PB[Lee Sedol]BR[6p]BT[South Korea]" +
	"PW[Gu Li]WR[9p]WT[China]RE[B+2]" +
	"CA[UTF-8]" +
	"AN[bob]CP[Copyright]DT[2014-12-25,26]HA[4]" +
	"EV[Pewter Cup]GN[sally]GC[it was long]" +
	"ON[low Chinese]OT[byo-yomi]PC[Seoul]RO[03 (final)]" +
	"RU[Japanese]SO[book]TM[1000]US[bill]" +
	"ZZ[zulu zimbabwe]YY[yello yulambi]" +
	")"

func TestParsingGameInfo(t *testing.T) {
	games, err := parse.ParseString(gameInfoString)
	assert.Equal(t, err, nil, "problem loading fixture")

	game := games[0]
	expected := "(;AN[bob]BR[6p]BT[South Korea]CA[UTF-8]CP[Copyright]DT[2014-12-25,26]EV[Pewter Cup]GC[it was long]GN[sally]HA[4]ON[low Chinese]OT[byo-yomi]PB[Lee Sedol]PC[Seoul]PW[Gu Li]RE[B+2]RO[03 (final)]RU[Japanese]SO[book]TM[1000]US[bill]WR[9p]WT[China]YY[yello yulambi]ZZ[zulu zimbabwe])"
	assert.Equal(t, game.String(), expected, "invalid gameInfo")
}

func TestParsingFullGameInfo(t *testing.T) {
	games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	assert.Equal(t, err, nil, "problem loading fixture")

	game := games[0]
	assert.Equal(t, game.GameInfo.String()[0:22],     ";BR[5p]C[This match wa", "invalid gameInfo")
	assert.Equal(t, game.GameInfo[sgf.Comment][0:27], "This match was sponsored by", "invalid comment")

	assert.Equal(t, game.GameInfo[sgf.Event], "The Game of the Century",  "invalid property")
	assert.Equal(t, game.GameInfo[sgf.BlackPlayerName], "Go Seigen",      "invalid property")
	assert.Equal(t, game.GameInfo[sgf.WhitePlayerName], "Honinbo Shusai", "invalid property")
}
