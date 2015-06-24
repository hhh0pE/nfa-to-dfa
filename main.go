package main

import (
	"fmt"
	"github.com/hhh0pE/NFAtoDFA/DFA"
	"github.com/hhh0pE/REtoNFA/NFA"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) == 1 {
		panic("No arguments passed. Exit.")
	}

	new_nfa := NFA.NewFromFile(os.Args[1])
	dfa := DFA.NewFromNFA(new_nfa)

	//    fmt.Printf("%+v", new_nfa)
	dfa.PrintDFA()

	//	nfa := NFA.BuildNFA("(aab+c)")
	//	nfa.PrintJSON()
}
