package parse

import (
	"fmt"
	"testing"
)

func verify(t *testing.T, sgi *GameInfo, propertyName, expected string) {
	value, _ := sgi.GetProperty(propertyName)

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
	sgi := &sgf.gameInfo

	verify(t, sgi, BlackPlayerName, "Lee Sedol")
	verify(t, sgi, BlackPlayerRank, "6p")
	verify(t, sgi, BlackPlayerTeam, "South Korea")

	verify(t, sgi, WhitePlayerName, "Gu Li")
	verify(t, sgi, WhitePlayerRank, "9p")
	verify(t, sgi, WhitePlayerTeam, "China")

	verify(t, sgi, Annotator, "bob")
	verify(t, sgi, Copyright, "Copyright")
	verify(t, sgi, Event, "Pewter Cup")
	verify(t, sgi, GameComment, "it was long")
	verify(t, sgi, Date, "2014-12-25,26")
	verify(t, sgi, GameName, "sally")
	verify(t, sgi, Handicap, "4")

	verify(t, sgi, Opening, "low Chinese")
	verify(t, sgi, Overtime, "byo-yomi")
	verify(t, sgi, Place, "Seoul")

	verify(t, sgi, Result, "B+2")
	verify(t, sgi, Round, "03 (final)")
	verify(t, sgi, Rules, "Japanese")
	verify(t, sgi, Source, "book")
	verify(t, sgi, TimeLimits, "1000")
	verify(t, sgi, User, "bill")

	verify(t, sgi, "charset", "UTF-8")

	verify(t, sgi, "ZZ", "zulu zimbabwe")
	verify(t, sgi, "YY", "yello yulambi")
}

func TestParsingFullGameInfo(t *testing.T) {
	sgf, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	if err != nil {
		t.Error(err)
		return
	}

	sgi := &sgf.gameInfo

	verify(t, sgi, "BlackPlayerName", "Go Seigen")
	verify(t, sgi, "BlackPlayerRank", "5p")

	verify(t, sgi, "WhitePlayerName", "Honinbo Shusai")
	verify(t, sgi, "WhitePlayerRank", "9p")

	verify(t, sgi, "Charset", "UTF-8")
	verify(t, sgi, "BoardSize", "19")
	verify(t, sgi, "Event", "The Game of the Century")
	verify(t, sgi, "Date", "1933-10-16")
	verify(t, sgi, "Place", "Tokyo, Japan")
	verify(t, sgi, "Result", "W+2")
	verify(t, sgi, "Komi", "0")

	foundComment, _ := sgi.GetProperty(Comment)
	expectedComment := "This match was sponsored by"

	if foundComment[0:27] != expectedComment {
		t.Error(fmt.Sprintf("invalid comment (found: '%s', expected: '%s')", foundComment, expectedComment))
	}
}
