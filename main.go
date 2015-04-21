package main

import (
	"fmt"
//	"github.com/hhh0pE/REtoNFA/NFA"
    "os"
)

func main() {
    fmt.Println(os.Args)
    if len(os.Args)==1 {
        panic("No arguments passed. Exit.")
    }
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error when opening file: "+err.Error())
    }
    fmt.Println(file)
//	nfa := NFA.BuildNFA("(aab+c)")
//	nfa.PrintJSON()
}
