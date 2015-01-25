package parse

import "testing"

func verify(t *testing.T, sgf *SGFGame, propertyName, expected string) {
	value, _ := sgf.gameInfo[propertyName]

	if value != expected {
		t.Errorf("invalid property: '%s' (found: '%s', expected: '%s')", propertyName, value, expected)
	}
}

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
	games, err := parseString(gameInfoString)
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	expected := "(;AN[bob]BR[6p]BT[South Korea]CA[UTF-8]CP[Copyright]DT[2014-12-25,26]EV[Pewter Cup]GC[it was long]GN[sally]HA[4]ON[low Chinese]OT[byo-yomi]PB[Lee Sedol]PC[Seoul]PW[Gu Li]RE[B+2]RO[03 (final)]RU[Japanese]SO[book]TM[1000]US[bill]WR[9p]WT[China]YY[yello yulambi]ZZ[zulu zimbabwe])"
	if game.String() != expected {
		t.Errorf("invalid gameInfo, found: %q, expected: %q", game.String(), expected)
	}
}

func TestParsingFullGameInfo(t *testing.T) {
	games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	if err != nil {
		t.Error(err)
		return
	}
	game := games[0]

	found := game.gameInfo.String()[0:22]
	expected := ";BR[5p]C[This match wa"
	if found != expected {
		t.Errorf("invalid gameInfo, found: %q, expected: %q)", found, expected)
	}

	foundComment, _ := game.GetInfo(Comment)
	expectedComment := "This match was sponsored by"

	if foundComment[0:27] != expectedComment {
		t.Errorf("invalid comment (found: '%s', expected: '%s')", foundComment, expectedComment)
	}

	verify(t, game, Event, "The Game of the Century")
	verify(t, game, BlackPlayerName, "Go Seigen")
	verify(t, game, WhitePlayerName, "Honinbo Shusai")
}
