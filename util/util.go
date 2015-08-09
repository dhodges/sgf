package util

import (
  "sort"
  "strings"
  "io/ioutil"
  "encoding/json"
)


func File2string(fpath string) (str string, err error) {
  var bytes []byte
  bytes, err = ioutil.ReadFile(fpath)
  if err == nil {
    str = string(bytes)
  }
  return str, err
}

func MapFromJson(json_str string) (map[string]string, error) {
  r := strings.NewReader(json_str)
  var json_map map[string]string
  err := json.NewDecoder(r).Decode(&json_map)
  if err != nil {
    return nil, err
  }

  return json_map, nil
}

func KeysFromMap(stringmap map[string]string) []string {
  var keys sort.StringSlice
  for k, _ := range stringmap {
    keys = append(keys, k)
  }
  sort.Sort(keys)
  return keys
}
