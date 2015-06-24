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

func walkFrom(node *NFA.Node, ids *[]int, symbol string, strict bool) {
    if node == nil {
        return
    }
    if node.Left != nil && (node.LeftSymbol == "" || node.LeftSymbol == symbol) {
        if strict && node.LeftSymbol=="" {
            return
        }
        // fmt.Println("left if")
        if inIntArray(*ids, node.Left.Id) {
            return
        }
        *ids = append(*ids, node.Left.Id)
        walkFrom(node.Left, ids, symbol, strict)
    }
    if node.Right != nil && (node.RightSymbol == "" || node.RightSymbol == symbol){
        if strict && node.RightSymbol=="" {
            return
        }
        if inIntArray(*ids, node.Right.Id) {
            return
        }
        *ids = append(*ids, node.Right.Id)
        walkFrom(node.Right, ids, symbol, strict)
    }
    // return ids
}

func NewFromNFA (nfa *NFA.NFA) *DFA {
    type NFATemp struct {
        Symbols map[string][]int
    }
    tmp_nodes :=  make([]NFATemp, nfa.Length())
    for i,_ := range tmp_nodes {
        tmp_nodes[i].Symbols = make(map[string][]int)
    }

    // making big table for help
    for i, node := range nfa.Nodes() {
        if node.LeftSymbol != "" && node.Left != nil {
            var nodes []int
            nodes = append(nodes, node.Left.Id)
            walkFrom(node.Left, &nodes, "", false)
            tmp_nodes[i].Symbols[node.LeftSymbol] = nodes
        }
        if node.RightSymbol != "" && node.Right != nil {
            var nodes []int
            nodes = append(nodes, node.Right.Id)
            walkFrom(node.Right, &nodes, "", false)
            tmp_nodes[i].Symbols[node.RightSymbol] = nodes
        }
    }

    tmp := make(map[string][][]int)

    for _, node := range tmp_nodes {
        for symb, val := range node.Symbols {
            tmp[symb] = append(tmp[symb], val)
        }
    }

    fmt.Printf("%+v\n", tmp)

    var start_nodes []int
    start_nodes = append(start_nodes, nfa.Nodes()[0].Id)
    walkFrom(nfa.Nodes()[0], &start_nodes, "", false)
    fmt.Printf("start nodes: %v\n", start_nodes)

    rules := make(map[string]string)

    var start_a, start_b []int
    for _, id := range start_nodes {

        walkFrom(nfa.Nodes()[id], &start_a, "a", true)
        walkFrom(nfa.Nodes()[id], &start_b, "b", true)
    }


    var a_ids []int
    for _, ids := range tmp["a"] {
        for _, id := range start_a {
            if inIntArray(ids, id) {
                a_ids = append(a_ids, ids...)
            }
        }
    }
    if _, exist := rules[fmt.Sprintf("%v", a_ids)]; !exist {
        rules[fmt.Sprintf("%v", a_ids)] = "A"
    }

    var b_ids []int
    
    for _, ids := range tmp["b"] {
        for _, id := range start_b {
            if inIntArray(ids, id) {
                b_ids = append(b_ids, ids...)
            }
        }
    }

    if _, exist := rules[fmt.Sprintf("%v", b_ids)]; !exist {
        rules[fmt.Sprintf("%v", b_ids)] = "B"
    }


    fmt.Printf("rules: %+v\n", rules)
    fmt.Printf("a: %v\n", a_ids)




    fmt.Printf("b: %v\n", b_ids)

    fmt.Printf("start_a: %v\n", start_a)
    fmt.Printf("start_b: %v\n", start_b)
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