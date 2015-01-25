package parse

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Property struct {
	name  string
	value string
}

func (p Property) String() string {
	return fmt.Sprintf("%s[%s]", p.name, p.value)
}

type Point struct {
	x rune
	y rune
}

func (point Point) String() string {
	return fmt.Sprintf("[%c%c]", point.x, point.y)
}

type Node struct {
	point      Property
	properties []Property
	variations []*Node
	next       *Node
}

func (node Node) variationString() string {
	if len(node.variations) == 0 {
		return ""
	}

	str := ""
	for _, nodevar := range node.variations {
		nodestr := ""
		for nptr := nodevar; nptr != nil; nptr = nptr.next {
			nodestr += nptr.String()
		}
		str += "(" + nodestr + ")"
	}
	return str
}

func (node Node) propertiesString() string {
	str := ""
	for _, prop := range node.properties {
		str += prop.String()
	}
	return str
}

func (node Node) String() string {
	return ";" +
		node.point.String() +
		node.propertiesString() +
		node.variationString()
}

func (node *Node) AddProperty(prop Property) {
	switch prop.name {
	case "B", "W":
		node.point = prop
	default:
		node.properties = append(node.properties, prop)
	}
}

func (node *Node) NewNode() *Node {
	node.next = new(Node)
	return node.next
}

func (n *Node) NewVariation() *Node {
	node := new(Node)
	n.variations = append(n.variations, node)
	return node
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

type SGFGame struct {
	gameInfo GameInfo
	gameTree *Node
	errors   []error
}

func (sgf *SGFGame) AddInfo(prop Property) {
	sgf.gameInfo[strings.ToUpper(prop.name)] = prop.value
}

func (sgf *SGFGame) GetInfo(name string) (value string, ok bool) {
	value, ok = sgf.gameInfo[strings.ToUpper(name)]
	return value, ok
}

func (sgf SGFGame) GameTreeString() string {
	treeString := ""
	for node := sgf.gameTree; node != nil; node = node.next {
		treeString += node.String()
	}
	return treeString
}

func (sgf SGFGame) String() string {
	return "(" + sgf.gameInfo.String() + sgf.GameTreeString() + ")"
}

func (sgf SGFGame) NodeCount() int {
	count := 0
	for node := sgf.gameTree; node != nil; node = node.next {
		count += 1
	}
	return count
}

func (sgf SGFGame) NthNode(n int) (node *Node, err error) {
	if n < 1 {
		return nil, errors.New("n less than 1")
	}
	nodeCount := sgf.NodeCount()

	if n > nodeCount {
		return nil, errors.New(fmt.Sprintf("n greater than node count (%d)", nodeCount))
	}
	for node = sgf.gameTree; n > 1; n -= 1 {
		node = node.next
	}
	return node, nil
}

func (sgf *SGFGame) AddError(msg string) {
	sgf.errors = append(sgf.errors, errors.New(msg))
}

func (sgf SGFGame) showAnyErrors() {
	if len(sgf.errors) == 0 {
		return
	}

	fmt.Println("Parsing errors:")
	for _, err := range sgf.errors {
		fmt.Println(err)
	}
}

func ListZipSGFfiles(fpath string) (fnames sort.StringSlice, err error) {
	r, err := zip.OpenReader(fpath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		fname := trim(f.Name)
		if isSGFfileName(fname) {
			fnames = append(fnames, fname)
		}
	}
	sort.Sort(fnames)
	return fnames, err
}

func trim(fname string) string {
	return strings.Trim(fname, " ")
}

func isSGFfileName(fname string) bool {
	fname = strings.ToLower(fname)
	suffix := fname[len(fname)-3:]
	return suffix == "sgf"
}

func zipSGFfileContents(f *zip.File) (contents string, err error) {
	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()

	bytes, err := ioutil.ReadAll(rc)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ParseZipSGFfile(zippath, fname string) (sgf *SGFGame, err error) {
	r, err := zip.OpenReader(zippath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	fname = strings.ToLower(fname)

	for _, f := range r.File {
		source := strings.ToLower(trim(f.Name))
		if source == fname {
			contents, err := zipSGFfileContents(f)
			if err != nil {
				return nil, err
			}
			sgf = new(SGFGame)
			sgf.gameInfo = make(GameInfo)
			sgf.Parse(contents)

			return sgf, err
		}
	}
	err = errors.New(fmt.Sprintf("file %q not found in zip archive %q", fname, zippath))
	return nil, err
}

func ParseZipAllSGFfiles(zippath string) (games []*SGFGame, err error) {
	fnames, err := ListZipSGFfiles(zippath)
	if err != nil {
		return nil, err
	}

	for _, f := range fnames {
		sgf, err := ParseZipSGFfile(zippath, f)
		if err != nil {
			return nil, err
		}
		games = append(games, sgf)
	}

	return games, nil
}

func (sgf *SGFGame) Parse(input string) *SGFGame {
	var currentNode *Node
	l := lex(input)
	prop := Property{}
	parsingSetup := false
	nodeStack := new(Stack)

Loop:
	for {
		i := l.nextItem()
		switch i.typ {
		case itemLeftParen:
			if parsingSetup {
				sgf.AddError(l.QuoteErrorContext("unexpected left parenthesis"))
				break Loop
			} else {
				if len(sgf.gameInfo) == 0 {
					parsingSetup = true
				} else {
					nodeStack.Push(currentNode)
					currentNode = currentNode.NewVariation()
					if l.nextItem().typ != itemSemiColon {
						sgf.AddError(l.QuoteErrorContext("semi-colon expected here"))
						break Loop
					}
				}
			}
		case itemRightParen:
			node := nodeStack.Pop()
			if node != nil {
				currentNode = node.(*Node)
			}
		case itemSemiColon:
			if parsingSetup {
				if len(sgf.gameInfo) > 0 {
					parsingSetup = false
					sgf.gameTree = new(Node)
					currentNode = sgf.gameTree
				}
			} else {
				currentNode = currentNode.NewNode()
			}
		case itemPropertyName:
			prop = Property{i.val, ""}
		case itemPropertyValue:
			prop.value = i.val
			if parsingSetup {
				sgf.AddInfo(prop)
			} else {
				currentNode.AddProperty(prop)
			}
		case itemError:
			sgf.AddError(i.val)
			break Loop
		case itemEOF:
			break Loop
		}
	}
	sgf.showAnyErrors()
	return sgf
}
