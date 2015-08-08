package tests

import (
  "testing"

  "github.com/dhodges/sgfinfo/sgf"
)

func TestGameInfoJson(t *testing.T) {
  games, err := parseFixture("19331016-Honinbo_Shusai-Go_Seigen.sgf")
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
  verifyProperty(t, gi, "BR", "5p")
  verifyProperty(t, gi, "CA", "UTF-8")
  verifyProperty(t, gi, "DT", "1933-10-16")
  verifyProperty(t, gi, "EV", "The Game of the Century")
  verifyProperty(t, gi, "KM", "0")
  verifyProperty(t, gi, "PB", "Go Seigen")
  verifyProperty(t, gi, "PW", "Honinbo Shusai")
  verifyProperty(t, gi, "RE", "W+2")
  verifyProperty(t, gi, "SZ", "19")
  verifyProperty(t, gi, "WR", "9p")
}

func verifyProperty(t *testing.T, gi sgf.GameInfo, propertyName, expected string) {
  value, _ := gi[propertyName]

  if value != expected {
    t.Errorf("invalid property: '%s' (found: '%s', expected: '%s')", propertyName, value, expected)
  }
}

