package main

import (
  "fmt"

  "github.com/dhodges/sgfinfo/util"
  "github.com/dhodges/sgfinfo/parse"
)

func dumpFileInfo(filename string) {
  str, err := util.File2string(filename)
  if err != nil {
    errorAndExit(fmt.Sprintf("\nproblem reading file: %s", filename))
  }

  games, err := parse.ParseString(str)
  if err != nil {
    errorAndExit(fmt.Sprintf("\nproblem parsing file: %s", filename))
  }

  // for now, assume one game per file
  b, err := games[0].GameInfo.ToJson()
  if err != nil {
    errorAndExit(fmt.Sprintf("\nproblem generating json (file: %s)", filename))
  }

  fmt.Println(string(b))
}
