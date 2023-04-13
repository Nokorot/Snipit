package main

import (
  "snipits/index"
)

import (
    "os"
    "log"
    "fmt"
)

func usage() {
    fmt.Println("Usage: %s <subcommans> <args...>")
    fmt.Println("Subcommands:")
    fmt.Println("    genIndex <path>    generate an index of the snipits in the spsifyed path")
    fmt.Println("    help               print this message")
}

func main() {
  if len(os.Args) < 1 { 
    log.Fatal("Not enough args")
    os.Exit(1)
  }

  switch os.Args[1] {
    case "genIndex":
        if len(os.Args) < 2 { 
          fmt.Println("The sub-command, genIndex expects an argument")
          os.Exit(1)
        }
        index.GenIndex(os.Args[2])
    case "help":
      usage()
    default:
      fmt.Println("Unknown subcommand!")
      os.Exit(1)
  }
}


