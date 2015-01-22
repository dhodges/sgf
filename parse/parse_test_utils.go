package parse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

func fixture_dirname() string {
	_, file, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(file)
	parDir := filepath.Dir(dirname)
	return parDir + "/" + "fixtures"
}

func sgf_fixture(fname string) (fixture string, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadFile(fixture_dirname() + "/sgf_files/" + fname)
	if err == nil {
		fixture = string(bytes)
	}
	return fixture, err
}

func listSgfFiles(dirname string) ([]string, error) {
	var fileList []string
	dirname = strings.TrimSpace(dirname)
	if dirname[len(dirname)-1] != '/' {
		dirname = dirname + "/"
	}
	fileInfoList, err := ioutil.ReadDir(dirname)
	if err != nil {
		return fileList, err
	}
	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			fileList = append(fileList, dirname+fileInfo.Name())
		}
	}
	return fileList, err
}

func parseString(str string) (sgf *SGFGame, err error) {
	sgf = new(SGFGame)
	sgf.gameInfo = map[string]string{}
	sgf.Parse(str)
	if len(sgf.errors) > 0 {
		return nil, errors.New(fmt.Sprintf("problems parsing sgf: %q", sgf.errors[0]))
	}

	return sgf, nil
}

func parseFixture(fixname string) (sgf *SGFGame, err error) {
	fixture, err := sgf_fixture(fixname)
	if err != nil {
		return sgf, err
	}

	return parseString(fixture)
}
