package sgf

import (
	"sort"
	"strings"
	"encoding/json"
)

type GameInfo map[string]string

func (gi GameInfo) SortedKeys() []string {
	var keys sort.StringSlice
	for k, _ := range gi {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

func (gi GameInfo) String() string {
	str := ""
	for _, k := range gi.SortedKeys() {
		str += k + "[" + gi[k] + "]"
	}
	return ";" + str
}

func (gi GameInfo) ToJson() ([]byte, error) {
	json_map := gi.clone()
	for _, k := range gi.SortedKeys() {
		key := propsToKeys[strings.ToUpper(k)]
		if key != "" {
			json_map[key] = json_map[k]
			delete(json_map, k)
		}
	}
	return json.Marshal(json_map)
}

func (gi GameInfo) FromJson(json_str string) (GameInfo, error) {
	json_map, err := mapFromJson(json_str)
	if err != nil {
		return nil, err
	}

	gameInfo := make(GameInfo)

	for k, _ := range json_map {
		key := keysToProps[k]
		if key != "" {
			gameInfo[strings.ToUpper(key)] = json_map[k]
		} else {
			gameInfo[strings.ToUpper(k)] = json_map[k]
		}
	}
	return gameInfo, nil
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

func (gi GameInfo) clone() map[string]string {
	duplicate := make(map[string]string)
	for k, v := range gi {
		duplicate[k] = v
	}
	return duplicate
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

var keysToProps = map[string]string{
	"Annotator":       "AN",
	"BlackPlayerName": "PB",
	"BlackPlayerRank": "BR",
	"BlackPlayerTeam": "BT",
	"Boardsize":       "SZ",
	"Charset":         "CA",
	"Comment":         "C",
	"Copyright":       "CP",
	"Date":            "DT",
	"Event":           "EV",
	"GameComment":     "GC",
	"GameName":        "GN",
	"Handicap":        "HA",
	"Komi":            "KM",
	"Opening":         "ON",
	"Overtime":        "OT",
	"Place":           "PC",
	"Result":          "RE",
	"Round":           "RO",
	"Rules":           "RU",
	"Source":          "SO",
	"TimeLimits":      "TM",
	"User":            "US",
	"WhitePlayerName": "PW",
	"WhitePlayerRank": "WR",
	"WhitePlayerTeam": "WT",
}
