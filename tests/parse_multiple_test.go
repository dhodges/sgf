package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/parse"
	"github.com/dhodges/sgfinfo/fixtures"
  "github.com/stretchr/testify/assert"
)

func TestParseMultiple(t *testing.T) {
	fixture, err := fixtures.Sgf("honinbo.sgf")
	assert.Equal(t, err, nil, "problem loading fixture")

	games := parse.Parse(fixture)
	assert.Equal(t, len(games), 259, "wrong number of games")

	game := games[6]
	assert.Equal(t, game.GameInfo[sgf.BlackPlayerName], "Hashimoto Utaro",  "wrong black player name")
	assert.Equal(t, game.GameInfo[sgf.Date],            "1943-05-05,06,07", "wrong date")
}
