package DFA

import (
    "fmt"
    "github.com/hhh0pE/REtoNFA/NFA"
)

type DFA struct {
	begin, end *Node
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

func walkFrom(node *NFA.Node, ids *[]int, symbol string) {
    if node == nil {
        return
    }
    if node.Left != nil && (node.LeftSymbol == "" || node.LeftSymbol == symbol) {
//        fmt.Println("left if")
        if inArray(*ids, node.Left.Id) {
           return
        }
        *ids = append(*ids, node.Left.Id)
        walkFrom(node.Left, ids, symbol)
    }
    if node.Right != nil && (node.RightSymbol == "" || node.RightSymbol == symbol){
        if inArray(*ids, node.Right.Id) {
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

    string_to_int := make(map[string][]int)
    string_to_int["1,2,3,4,8,9,11,12"] = []int{1,2,3,4,8,9,11,12}

    symbols := make(map[string]bool)

    dfa_nodes := make(map[string]map[string]string)
    dfa_nodes["1,2,3,4,5"] = make(map[string]string)
    dfa_nodes["1,2,3,4,5"]["a"] = "1,2,3,4"
    fmt.Println(dfa_nodes)

    for _, node := range nfa.Nodes() {
        if node.LeftSymbol != "" {
            symbols[node.LeftSymbol] = true
        }
        if node.RightSymbol != "" {
            symbols[node.RightSymbol] = true
        }
    }

    fmt.Println(symbols)

    for _, node := range nfa.Nodes() {
        var nodes []int
        nodes = append(nodes, node.Id)
        walkFrom(node, &nodes, "")
        fmt.Println(nodes)
    }
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