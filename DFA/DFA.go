package DFA

import (
    "fmt"
    "github.com/hhh0pE/REtoNFA/NFA"
)

type DFA struct {
	begin, end *Node
	nodes      []*Node
}

func isIntArraysEqual(arr []int, arr2 []int) bool {
    if len(arr) != len(arr2) {
        return false
    }

    for _, val := range arr {
        if !inIntArray(arr2, val) {
            return false
        }
    }

    return true
}

func inIntArray(arr []int, num int) bool {
    for _, val := range arr {
        if val == num {
            return true
        }
    }
    return false
}

func inStringArray(arr []string, str string) bool {
    for _, val := range arr {
        if val == str {
            return true
        }
    }

    return false
}

func walkFrom(node *NFA.Node, ids *[]int, symbol string) {
    if node == nil {
        return
    }
    if node.Left != nil && (node.LeftSymbol == "" || node.LeftSymbol == symbol) {
//        fmt.Println("left if")
        if inIntArray(*ids, node.Left.Id) {
           return
        }
        *ids = append(*ids, node.Left.Id)
        walkFrom(node.Left, ids, symbol)
    }
    if node.Right != nil && (node.RightSymbol == "" || node.RightSymbol == symbol){
        if inIntArray(*ids, node.Right.Id) {
            return
        }
        *ids = append(*ids, node.Right.Id)
        walkFrom(node.Right, ids, symbol)
    }

//    return ids
}

func NewFromNFA (nfa *NFA.NFA) *DFA {
    type NFATemp struct {
        Symbols map[string][]int
    }

    NodeBucketIds := make(map[int][]int)


    var dfa_nodes_count int

    var symbols []string

    for _, node := range nfa.Nodes() {
        if node.LeftSymbol!="" && !inStringArray(symbols, node.LeftSymbol) {
            symbols = append(symbols, node.LeftSymbol)
        }
        if node.RightSymbol!="" && !inStringArray(symbols, node.RightSymbol) {
            symbols = append(symbols, node.RightSymbol)
        }
    }

//    nodes := make(map[int]map[string]int)

    var tmp_nodes []int
    tmp_nodes = append(tmp_nodes, nfa.Nodes()[0].Id)
    walkFrom(nfa.Nodes()[0], &tmp_nodes, "")
    NodeBucketIds[dfa_nodes_count] = tmp_nodes
    dfa_nodes_count++

    fmt.Println(NodeBucketIds)

//    for _, node := range nfa.Nodes() {
//
//    }
//
//    var nodes []int
//    walkFrom(nfa.Nodes()[0], &nodes, "")
//    fmt.Printf("%+v\n", nodes)

//    for _, node := range nfa.Nodes() {
//        var nodes []int
//        nodes = append(nodes, node.Id)
//        walkFrom(node, &nodes, "")
//        fmt.Printf("%+v: %+v\n", node, nodes)
//    }
//    tmp_nodes :=  make([]NFATemp, nfa.Length())
//    for i,_ := range tmp_nodes {
//        tmp_nodes[i].Symbols = make(map[string][]int)
//    }
//
//    for i, node := range nfa.Nodes() {
//        fmt.Printf("NFA Node: %+v\n", node)
//        if node.LeftSymbol != "" && node.Left != nil {
//            var nodes []int
//            nodes = append(nodes, node.Left.Id)
//            walkFrom(node.Left, &nodes)
//            tmp_nodes[i].Symbols[node.LeftSymbol] = nodes
//        }
//        if node.RightSymbol != "" && node.Right != nil {
//            var nodes []int
//            nodes = append(nodes, node.Right.Id)
//            walkFrom(node.Right, &nodes)
//            tmp_nodes[i].Symbols[node.RightSymbol] = nodes
//        }
//    }
//
//    for i, tmp := range tmp_nodes {
//        fmt.Printf("%d %+v\n", i, tmp)
//    }
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