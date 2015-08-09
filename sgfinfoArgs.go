package main

import (
  "os"
  "fmt"
)

var usage = fmt.Sprintf("\nusage: %s <sgf_filename>\n", os.Args[0])

func errorAndExit(message string) {
  fmt.Println(message, "\n")
  fmt.Println(usage)
  os.Exit(1)
}

func validateFile(filename string) {
  fileInfo, err := os.Stat(filename)
  if os.IsNotExist(err) {
    errorAndExit(fmt.Sprintf("\nno such file: %s", filename))
  }
  if fileInfo.IsDir() {
    errorAndExit(fmt.Sprintf("\n%s is a directory", filename))
  }
}
