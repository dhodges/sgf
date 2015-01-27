package parse

import (
	"testing"

	"github.com/dhodges/sgf"
)

func TestParseMultiple(t *testing.T) {
	fixture, err := sgf_fixture("honinbo.sgf")
	if err != nil {
		t.Error("problem reading fixture 'honinbo.sgf'")
		return
	}

	games := Parse(fixture)
	if len(games) != 259 {
		t.Errorf("wrong number of games, found %d, expected 259", len(games))
	}

	game := games[6]

	foundName, _ := game.GetInfo(sgf.BlackPlayerName)
	if foundName != "Hashimoto Utaro" {
		t.Errorf("wrong black player name, found: %q, expected: %q", foundName, "Hashimoto Utaro")
	}
	foundDate, _ := game.GetInfo(sgf.Date)
	if foundDate != "1943-05-05,06,07" {
		t.Errorf("wrong date, found: %q, expected: %q", foundDate, "1943-05-05,06,07")
	}
}
