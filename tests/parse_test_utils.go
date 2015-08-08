package tests

import (
	"errors"
	"fmt"

	"github.com/dhodges/sgfinfo/sgf"
	"github.com/dhodges/sgfinfo/parse"
	"github.com/dhodges/sgfinfo/fixtures"
)

func parseString(str string) (games []*sgf.Game, err error) {
	games = parse.Parse(str)
	if len(games[0].Errors) > 0 {
		return nil, errors.New(fmt.Sprintf("problems parsing sgf: %q", games[0].Errors[0]))
	}

	return games, nil
}

func parseFixture(fixname string) (games []*sgf.Game, err error) {
	fixture, err := fixtures.Sgf(fixname)
	if err != nil {
		return games, err
	}

	return parseString(fixture)
}
