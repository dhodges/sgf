package sgf

import "fmt"

type Property struct {
	Name  string
	Value string
}

func (p Property) String() string {
	return fmt.Sprintf("%s[%s]", p.Name, p.Value)
}

const BlackPlayerName = "PB"
const BlackPlayerRank = "BR"
const BlackPlayerTeam = "BT"
const WhitePlayerName = "PW"
const WhitePlayerRank = "WR"
const WhitePlayerTeam = "WT"
const Annotator = "AN"
const Copyright = "CP"
const Date = "DT"
const Event = "EV"
const GameComment = "GC"
const Comment = "C"
const GameName = "GN"
const Handicap = "HA"
const Opening = "ON"
const Overtime = "OT"
const Place = "PC"
const Result = "RE"
const Round = "RO"
const Rules = "RU"
const Source = "SO"
const TimeLimits = "TM"
const User = "US"
const Charset = "CA"
const Boardsize = "SZ"
const Komi = "KM"
