package sgf

import  (
	"sort"
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
	return json.Marshal(gi)
}

func (gi *GameInfo) FromJson(json_str string) error {
  return json.Unmarshal([]byte(json_str), &gi)
}
