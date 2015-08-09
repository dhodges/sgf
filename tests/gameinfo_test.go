package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameInfoJson(t *testing.T) {
	games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	assert.Equal(t, err, nil, "problem loading fixture")

	gi := games[0].GameInfo
	assert.Equal(t, "5p",             gi["BR"], "")
	assert.Equal(t, "UTF-8",          gi["CA"], "")
	assert.Equal(t, "1933-10-16",     gi["DT"], "")
	assert.Equal(t, "0",              gi["KM"], "")
	assert.Equal(t, "Go Seigen",      gi["PB"], "")
	assert.Equal(t, "Honinbo Shusai", gi["PW"], "")
	assert.Equal(t, "W+2",            gi["RE"], "")
	assert.Equal(t, "19",             gi["SZ"], "")
	assert.Equal(t, "9p",             gi["WR"], "")
	assert.Equal(t, "The Game of the Century", gi["EV"], "")

	b, err := gi.ToJson()
	assert.Equal(t, err, nil, "problem generating json")

	gi2, err := gi.FromJson(string(b))
	assert.Equal(t, err, nil, "problem unmarshalling json")

	assert.Equal(t, "5p",             gi2["BR"], "")
	assert.Equal(t, "UTF-8",          gi2["CA"], "")
	assert.Equal(t, "1933-10-16",     gi2["DT"], "")
	assert.Equal(t, "0",              gi2["KM"], "")
	assert.Equal(t, "Go Seigen",      gi2["PB"], "")
	assert.Equal(t, "Honinbo Shusai", gi2["PW"], "")
	assert.Equal(t, "W+2",            gi2["RE"], "")
	assert.Equal(t, "19",             gi2["SZ"], "")
	assert.Equal(t, "9p",             gi2["WR"], "")
	assert.Equal(t, "The Game of the Century", gi2["EV"], "")
}

func TestGameInfoJsonKeys(t *testing.T) {
	games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
	assert.Equal(t, err, nil, "problem loading fixture")

	b, err := games[0].GameInfo.ToJson()
	assert.Equal(t, err, nil, "problem generating json")

	keyMap, err := mapFromJson(string(b))
	keys := keysFromMap(keyMap)

	assert.Equal(t, "BlackPlayerName", keys[0], "")
	assert.Equal(t, "BlackPlayerRank", keys[1], "")
	assert.Equal(t, "Boardsize",       keys[2], "")
	assert.Equal(t, "Charset",         keys[3], "")
	assert.Equal(t, "Comment",         keys[4], "")
	assert.Equal(t, "Date",            keys[5], "")
	assert.Equal(t, "Event",           keys[6], "")
	assert.Equal(t, "Komi",            keys[7], "")
	assert.Equal(t, "Place",           keys[8], "")
	assert.Equal(t, "Result",          keys[9], "")
	assert.Equal(t, "WhitePlayerName", keys[10], "")
	assert.Equal(t, "WhitePlayerRank", keys[11], "")
}
