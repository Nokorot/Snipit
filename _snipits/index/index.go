package index

import (
    "snipits/snipit"
)

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "os"

)

type TemplateEntry struct {
    Key     string
    Vars    map[string]string
}


type TemplateData struct {
    File    string
    Data    []TemplateEntry
}

type Node struct {
    // Name   string
    Snipits map[string]snipit.Snipit
}


func read_templatefile(node Node, path string, name string) {
  snipits := node.Snipits

  content, err := ioutil.ReadFile(path + "/" + name)
  if err != nil { fmt.Println("Err") }

  var data []TemplateData
  json.Unmarshal(content, &data)
    
  for _, td := range data {
    for _, entry := range td.Data {
      // TODO: check for overvride
      snipits[entry.Key] = snipit.Snipit{
          // TODO: _templates not hardcoded
          File: path + "/_templates/" + td.File,
          Vars: entry.Vars,
      }
    }
  }
    
}

func gen_dir(node Node, path string) {
    snipits := node.Snipits

    lfiles, err := ioutil.ReadDir(path)
    if err != nil {
      log.Fatal(err)
    }

    _template_file := false
    _template_dir := false

    for _, entry := range lfiles {
      if !entry.IsDir() {
        if entry.Name()[0] != '_' {
          snipit := snipit.Snipit{
              // Name: entry.Name(),
              File: path + "/" + entry.Name(),
              Vars: make(map[string]string) }
          snipits[entry.Name()] = snipit
        } else if entry.Name() == "_templates.json" {
          _template_file = true
          read_templatefile(node, path, entry.Name())
        }
      } else if entry.Name()[0] != '_' {
        gen_dir(node, path + "/" + entry.Name())
      } else if entry.Name() == "_templates" {
        _template_dir = true
      }
    }

    if _template_file != _template_dir {
      if _template_file {
        fmt.Println("Worning! There is a '_templates.json' file but no '_templates' folder")
      } else {
        fmt.Println("Worning! There is a '_templates' folder, but no '_templates.json' file")
      }
    }
}

func GenIndex(path string) {
  // if len(os.Args) < 1 { 
  //   log.Fatal("Not enough args")
  //   os.Exit(1)
  // }
  // path := os.Args[1] 

  os.Chdir(path)
  // path := "./test"

  entries, err := ioutil.ReadDir(".")
  if err != nil {
      log.Fatal(err)
  }

  var nodes = make(map[string]Node) // []*Node
  for _, entry := range entries {
    if entry.IsDir() {
      node := Node{
          Snipits: make(map[string]snipit.Snipit),
      }

      gen_dir(node, entry.Name())

      nodes[entry.Name()] = node
    }
  }

  res, _ := json.Marshal(nodes)
  fmt.Println(string(res))
}

