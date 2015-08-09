package tests

import (
	"fmt"
	"errors"
	"sort"
	"strings"
	"encoding/json"

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

func mapFromJson(json_str string) (map[string]string, error) {
	r := strings.NewReader(json_str)
	var json_map map[string]string
	err := json.NewDecoder(r).Decode(&json_map)
	if err != nil {
		return nil, err
	}

	return json_map, nil
}

func keysFromMap(keyMap map[string]string) (keys sort.StringSlice) {
	for k, _ := range keyMap {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}
