package sgf

import (
	"strings"
	"encoding/json"

	"github.com/dhodges/sgfinfo/util"
)

type GameInfo map[string]string

func (gi GameInfo) String() string {
	str := ""
	for _, k := range util.KeysFromMap(gi) {
		str += k + "[" + gi[k] + "]"
	}
	return ";" + str
}

func (gi GameInfo) ToJson() ([]byte, error) {
	json_map := gi.clone()
	for _, k := range util.KeysFromMap(gi) {
		key := property2key(k)
		if key != "" {
			json_map[key] = json_map[k]
			delete(json_map, k)
		}
	}
	return json.Marshal(json_map)
}

func (gi GameInfo) FromJson(json_str string) (GameInfo, error) {
	json_map, err := util.MapFromJson(json_str)
	if err != nil {
		return nil, err
	}

	gameInfo := make(GameInfo)

	for k, _ := range json_map {
		key := key2property(k)
		if key != "" {
			gameInfo[strings.ToUpper(key)] = json_map[k]
		} else {
			gameInfo[strings.ToUpper(k)] = json_map[k]
		}
	}
	return gameInfo, nil
}

func (gi GameInfo) clone() map[string]string {
	duplicate := make(map[string]string)
	for k, v := range gi {
		duplicate[k] = v
	}
	return duplicate
}

func property2key(propname string) string {
	return propsToKeys[strings.ToUpper(propname)]
}

var propsToKeys = map[string]string{
	"AN": "Annotator",
	"PB": "BlackPlayerName",
	"BR": "BlackPlayerRank",
	"BT": "BlackPlayerTeam",
	"SZ": "Boardsize",
	"CA": "Charset",
	"C":  "Comment",
	"CP": "Copyright",
	"DT": "Date",
	"EV": "Event",
	"GC": "GameComment",
	"GN": "GameName",
	"HA": "Handicap",
	"KM": "Komi",
	"ON": "Opening",
	"OT": "Overtime",
	"PC": "Place",
	"RE": "Result",
	"RO": "Round",
	"RU": "Rules",
	"SO": "Source",
	"TM": "TimeLimits",
	"US": "User",
	"PW": "WhitePlayerName",
	"WR": "WhitePlayerRank",
	"WT": "WhitePlayerTeam",
}

func key2property(key string) string {
	return keysToProps[strings.ToLower(key)]
}

var keysToProps = map[string]string{
	"annotator":       "AN",
	"blackplayername": "PB",
	"blackplayerrank": "BR",
	"blackplayerteam": "BT",
	"boardsize":       "SZ",
	"charset":         "CA",
	"comment":         "C",
	"copyright":       "CP",
	"date":            "DT",
	"event":           "EV",
	"gamecomment":     "GC",
	"gamename":        "GN",
	"handicap":        "HA",
	"komi":            "KM",
	"opening":         "ON",
	"overtime":        "OT",
	"place":           "PC",
	"result":          "RE",
	"round":           "RO",
	"rules":           "RU",
	"source":          "SO",
	"timelimits":      "TM",
	"user":            "US",
	"whiteplayername": "PW",
	"whiteplayerrank": "WR",
	"whiteplayerteam": "WT",
}
