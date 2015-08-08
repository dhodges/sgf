package tests

import (
  "testing"

  "github.com/dhodges/sgfinfo/sgf"
  "github.com/dhodges/sgfinfo/parse"
)

func TestGameInfoJson(t *testing.T) {
  games, err := parse.ParseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
  if err != nil {
    t.Error(err)
    return
  }
  game := games[0]

  verifyGameInfo(t, game.GameInfo)

  b, err := game.GameInfo.ToJson()
  if err != nil {
    t.Error(err)
    return
  }

  var gi sgf.GameInfo
  err = gi.FromJson(string(b))
  if err != nil {
    t.Error(err)
    return
  }

  verifyGameInfo(t, gi)
}


func verifyGameInfo(t *testing.T, gi sgf.GameInfo) {
  verify(t, gi, "BR", "5p")
  verify(t, gi, "CA", "UTF-8")
  verify(t, gi, "DT", "1933-10-16")
  verify(t, gi, "EV", "The Game of the Century")
  verify(t, gi, "KM", "0")
  verify(t, gi, "PB", "Go Seigen")
  verify(t, gi, "PW", "Honinbo Shusai")
  verify(t, gi, "RE", "W+2")
  verify(t, gi, "SZ", "19")
  verify(t, gi, "WR", "9p")
}

func verify(t *testing.T, gi sgf.GameInfo, propertyName, expected string) {
  value, _ := gi[propertyName]

  if value != expected {
    t.Errorf("invalid property: '%s' (found: '%s', expected: '%s')", propertyName, value, expected)
  }
}

