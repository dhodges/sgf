package main

import (
  "os"
)

func main() {
  if len(os.Args) != 2 {
    errorAndExit("sgf filename missing")
  }

  filename := os.Args[1]
  validateFile(filename)

  dumpFileInfo(filename)
}
