package main

import (
	"fmt"
	"github.com/hhh0pE/REtoNFA/NFA"
    "github.com/hhh0pE/NFAtoDFA/DFA"
    "os"
)

func main() {
    fmt.Println(os.Args)
    if len(os.Args)==1 {
        panic("No arguments passed. Exit.")
    }

    new_nfa := NFA.NewFromFile(os.Args[1])
    new_nfa.PrintNFA()
    dfa := DFA.NewFromNFA(new_nfa)

//    fmt.Printf("%+v", new_nfa)
    dfa.PrintDFA()

//	nfa := NFA.BuildNFA("(aab+c)")
//	nfa.PrintJSON()
}
