package DFA

import (
    "fmt"
    "github.com/hhh0pE/REtoNFA/NFA"
)

type DFA struct {
	begin, end *DFA
	nodes      []*Node
}

func isArraysEqual(arr []int, arr2 []int) bool {
    if len(arr) != len(arr2) {
        return false
    }

    for _, val := range arr {
        if !inArray(arr2, val) {
            return false
        }
    }

    return true
}

func inArray(arr []int, num int) bool {
    for _, val := range arr {
        if val == num {
            return true
        }
    }
    return false
}

func walkFrom(node *NFA.Node, ids *[]int) {
    if node == nil {
        return
    }
    if node.Left != nil && node.LeftSymbol == "" {
//        fmt.Println("left if")
        if inArray(*ids, node.Left.Id) {
           return
        }
        *ids = append(*ids, node.Left.Id)
        walkFrom(node.Left, ids)
    }
    if node.Right != nil && node.RightSymbol == "" {
        if inArray(*ids, node.Right.Id) {
            return
        }
        *ids = append(*ids, node.Right.Id)
        walkFrom(node.Right, ids)
    }

//    return ids
}

func NewFromNFA (nfa *NFA.NFA) *DFA {
    type NFATemp struct {
        Symbols map[string][]int
    }
    tmp_nodes :=  make([]NFATemp, nfa.Length())
    for i,_ := range tmp_nodes {
        tmp_nodes[i].Symbols = make(map[string][]int)
    }

    for i, node := range nfa.Nodes() {
        fmt.Printf("%+v\n", node)
        if node.LeftSymbol != "" && node.Left != nil {
            var nodes []int
            nodes = append(nodes, node.Left.Id)
            walkFrom(node.Left, &nodes)
            tmp_nodes[i].Symbols[node.LeftSymbol] = nodes
        }
        if node.RightSymbol != "" && node.Right != nil {
            var nodes []int
            nodes = append(nodes, node.Right.Id)
            walkFrom(node.Right, &nodes)
            tmp_nodes[i].Symbols[node.RightSymbol] = nodes
        }
    }

    for i, tmp := range tmp_nodes {
        fmt.Printf("%d %+v\n", i, tmp)
    }
    return &DFA{}
}

func (dfa *DFA) PrintDFA() {
    if dfa != nil {
        fmt.Println("DFA built success. Printing DFA..")
        if len(dfa.nodes) == 0{
            fmt.Println("DFA is empty.")
            return
        }
        for _, node := range dfa.nodes {
            fmt.Println(node.toString())
        }
    } else {
        fmt.Println("Error when building DFA")
    }
}