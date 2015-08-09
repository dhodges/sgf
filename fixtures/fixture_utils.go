package fixtures

import (
	"path/filepath"
	"runtime"

	"github.com/dhodges/sgfinfo/util"
)

func fixture_dirname() string {
	_, file, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(file)
	parDir := filepath.Dir(dirname)
	return parDir + "/" + "fixtures"
}

func Sgf(fname string) (string, error) {
	return util.File2string(fixture_dirname() + "/sgf_files/" + fname)
}

func Zip_fpath(fname string) string {
	return fixture_dirname() + "/zip_files/" + fname
}
