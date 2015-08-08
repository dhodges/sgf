package parse

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/dhodges/sgf/sgf"
)

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

func ParseZipSGFfile(zippath, fname string) (games []*sgf.SGFGame, err error) {
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

			return Parse(contents), err
		}
	}
	err = errors.New(fmt.Sprintf("file %q not found in zip archive %q", fname, zippath))
	return nil, err
}

func ParseZipAllSGFfiles(zippath string) (games []*sgf.SGFGame, err error) {
	fnames, err := ListZipSGFfiles(zippath)
	if err != nil {
		return nil, err
	}

	for _, f := range fnames {
		zippedGames, err := ParseZipSGFfile(zippath, f)
		if err != nil {
			return nil, err
		}
		games = append(games, zippedGames...)
	}

	return games, nil
}
