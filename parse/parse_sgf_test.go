package parse

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParsingSGF(t *testing.T) {
	dirname := "/Users/david/Google Drive/SGF/AWAGC-2014"
	fileList, err := listSgfFiles(dirname)
	if err != nil {
		t.Errorf("error reading sgf file list: %s", err.Error())
		return
	}
Loop:
	for _, fname := range fileList {
		buf, err := ioutil.ReadFile(fname)
		if err != nil {
			t.Errorf("Error reading file: %q, %q", fname, err.Error())
			return
		}

		game := Parse(string(buf))[0]
		if len(game.errors) > 0 {
			fmt.Printf("problems parsing file: %q\n", fname)
			fmt.Println(game.errors[0].Error())
			break Loop
		}
	}
}

func TestSGFErrors(t *testing.T) {
	sgf := new(SGFGame)
	sgf.gameInfo = make(GameInfo)
	sgf.AddError("hells belles")
	sgf.AddError("hades' ladies")
	if len(sgf.errors) != 2 {
		t.Errorf("SGFGame error count is wrong (found %d, expected 2)", len(sgf.errors))
	}
}
