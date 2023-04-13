package snipit

import (
    "encoding/json"
    "fmt"
)

type Snipit struct {
    File    string
    Vars    map[string]string
}


func GenSnipit(json_file string) {
  data, err := ioutil.ReadFile(json_file)
  if err != nil { log.Fatal(err) }

  var snipit Snipit
  json.Unmarshal(data, &snipit)

}

