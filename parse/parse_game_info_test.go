package parse

import (
	"fmt"
	"testing"
)

func verify(t *testing.T, sgf *SGFGame, propertyName, expected string) {
	value, _ := sgf.gameInfo[propertyName]

	if value != expected {
		t.Error(fmt.Sprintf("invalid property: '%s' (found: '%s', expected: '%s')", propertyName, value, expected))
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

	verify(t, sgf, BlackPlayerName, "Lee Sedol")
	verify(t, sgf, BlackPlayerRank, "6p")
	verify(t, sgf, BlackPlayerTeam, "South Korea")

	verify(t, sgf, WhitePlayerName, "Gu Li")
	verify(t, sgf, WhitePlayerRank, "9p")
	verify(t, sgf, WhitePlayerTeam, "China")

	verify(t, sgf, Annotator, "bob")
	verify(t, sgf, Copyright, "Copyright")
	verify(t, sgf, Event, "Pewter Cup")
	verify(t, sgf, GameComment, "it was long")
	verify(t, sgf, Date, "2014-12-25,26")
	verify(t, sgf, GameName, "sally")
	verify(t, sgf, Handicap, "4")

	verify(t, sgf, Opening, "low Chinese")
	verify(t, sgf, Overtime, "byo-yomi")
	verify(t, sgf, Place, "Seoul")

	verify(t, sgf, Result, "B+2")
	verify(t, sgf, Round, "03 (final)")
	verify(t, sgf, Rules, "Japanese")
	verify(t, sgf, Source, "book")
	verify(t, sgf, TimeLimits, "1000")
	verify(t, sgf, User, "bill")

	verify(t, sgf, "charset", "UTF-8")

	verify(t, sgf, "ZZ", "zulu zimbabwe")
	verify(t, sgf, "YY", "yello yulambi")
}

func TestParsingFullGameInfo(t *testing.T) {
	sgf, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	if err != nil {
		t.Error(err)
		return
	}

	verify(t, sgf, "BlackPlayerName", "Go Seigen")
	verify(t, sgf, "BlackPlayerRank", "5p")

	verify(t, sgf, "WhitePlayerName", "Honinbo Shusai")
	verify(t, sgf, "WhitePlayerRank", "9p")

	verify(t, sgf, "Charset", "UTF-8")
	verify(t, sgf, "BoardSize", "19")
	verify(t, sgf, "Event", "The Game of the Century")
	verify(t, sgf, "Date", "1933-10-16")
	verify(t, sgf, "Place", "Tokyo, Japan")
	verify(t, sgf, "Result", "W+2")
	verify(t, sgf, "Komi", "0")

	foundComment, _ := sgf.GetProperty(Comment)
	expectedComment := "This match was sponsored by"

	if foundComment[0:27] != expectedComment {
		t.Error(fmt.Sprintf("invalid comment (found: '%s', expected: '%s')", foundComment, expectedComment))
	}
}
