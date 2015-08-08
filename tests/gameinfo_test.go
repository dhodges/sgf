package tests

import (
  "testing"

  "github.com/dhodges/sgfinfo/sgf"
  "github.com/stretchr/testify/assert"
)

func TestGameInfoJson(t *testing.T) {
  games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
  assert.Equal(t, err, nil, "problem loading fixture")

  game := games[0]
  assert.Equal(t, game.GameInfo["BR"], "5p", "")
  assert.Equal(t, game.GameInfo["CA"], "UTF-8", "")
  assert.Equal(t, game.GameInfo["DT"], "1933-10-16", "")
  assert.Equal(t, game.GameInfo["EV"], "The Game of the Century", "")
  assert.Equal(t, game.GameInfo["KM"], "0", "")
  assert.Equal(t, game.GameInfo["PB"], "Go Seigen", "")
  assert.Equal(t, game.GameInfo["PW"], "Honinbo Shusai", "")
  assert.Equal(t, game.GameInfo["RE"], "W+2", "")
  assert.Equal(t, game.GameInfo["SZ"], "19", "")
  assert.Equal(t, game.GameInfo["WR"], "9p", "")

  b, err := game.GameInfo.ToJson()
  assert.Equal(t, err, nil, "problem generating json")

  var gi sgf.GameInfo
  err = gi.FromJson(string(b))
  assert.Equal(t, err, nil, "problem unmarshalling json")

}




  }
}

