package tests

import (
	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/parse"
	"github.com/dhodges/sgfinfo/fixtures"
)

func parseFixture(fixname string) (games []*sgf.Game, err error) {
	fixture, err := fixtures.Sgf(fixname)
	if err != nil {
		return games, err
	}

	return parse.ParseString(fixture)
}
