package parse

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
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
