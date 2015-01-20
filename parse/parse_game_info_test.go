package parse

import (
	"fmt"
	"testing"
)

func verify(t *testing.T, field, fieldName, expected string) {
	if field != expected {
		t.Error(fmt.Sprintf("invalid field: '%s' (found: '%s', expected: '%s')", fieldName, field, expected))
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
	sgf, err := parseString(gameInfoString)
	if err != nil {
		t.Error(err)
		return
	}
	sgi := &sgf.gameInfo

	verify(t, sgi.black.name, "black player", "Lee Sedol")
	verify(t, sgi.black.rank, "black player rank", "6p")
	verify(t, sgi.black.team, "black player team", "South Korea")

	verify(t, sgi.white.name, "white player", "Gu Li")
	verify(t, sgi.white.rank, "white player rank", "9p")
	verify(t, sgi.white.team, "white player team", "China")

	verify(t, sgi.annotator, "annotator", "bob")
	verify(t, sgi.copyright, "copyright", "Copyright")
	verify(t, sgi.event, "event", "Pewter Cup")
	verify(t, sgi.gameInfo, "game comment", "it was long")
	verify(t, sgi.date, "game date", "2014-12-25,26")
	verify(t, sgi.gameName, "game name", "sally")
	verify(t, sgi.handicap, "handicap", "4")

	verify(t, sgi.opening, "opening", "low Chinese")
	verify(t, sgi.overtime, "overtime", "byo-yomi")
	verify(t, sgi.place, "place", "Seoul")

	verify(t, sgi.result, "result", "B+2")
	verify(t, sgi.round, "round", "03 (final)")
	verify(t, sgi.rules, "rules", "Japanese")
	verify(t, sgi.source, "source", "book")
	verify(t, sgi.timeLimits, "time limits", "1000")
	verify(t, sgi.user, "user", "bill")

	verify(t, sgi.charset, "charset", "UTF-8")

	if len(sgi.unknownProperties) != 2 {
		t.Error("unknown properties were not kept")
	}

	if sgi.unknownProperties[0].String() != "ZZ[zulu zimbabwe]" {
		t.Error("game info 'UnknownProperties' are incorrect")
	}

	if sgi.unknownProperties[1].String() != "YY[yello yulambi]" {
		t.Error("game info 'UnknownProperties' are incorrect")
	}
}

func TestParsingFullGameInfo(t *testing.T) {
	sgf, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	if err != nil {
		t.Error(err)
		return
	}

	sgi := sgf.gameInfo

	verify(t, sgi.black.name, "black player", "Go Seigen")
	verify(t, sgi.black.rank, "black player rank", "5p")

	verify(t, sgi.white.name, "white player", "Honinbo Shusai")
	verify(t, sgi.white.rank, "white player rank", "9p")

	verify(t, sgi.charset, "charset", "UTF-8")
	verify(t, sgi.boardsize, "board size", "19")
	verify(t, sgi.event, "event", "The Game of the Century")
	verify(t, sgi.date, "date", "1933-10-16")
	verify(t, sgi.place, "place", "Tokyo, Japan")
	verify(t, sgi.result, "result", "W+2")
	verify(t, sgi.komi, "komi", "0")

	foundComment := string(sgi.comment[0:27])
	expectedComment := "This match was sponsored by"

	if foundComment != expectedComment {
		t.Error(fmt.Sprintf("invalid comment (found: '%s', expected: '%s')", foundComment, expectedComment))
	}
}
