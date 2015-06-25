package DFA

import (
    "fmt"
    "github.com/hhh0pE/REtoNFA/NFA"
    "strings"
    "strconv"
)

type DFA struct {
	begin, end *DFA
	nodes      []*Node
}

func (nfa *DFA) addNode(new_node *Node) {
    new_node.Id = len(nfa.nodes)
    new_node.Directions = make(map[string]*Node)
    nfa.nodes = append(nfa.nodes, new_node)
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

func inStringArray(arr []string, item string) bool {
    for _, val := range arr {
        if val == item {
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

func makeKeyFromIntArray(arr []int) string {
    key := fmt.Sprintf("%v", arr)
    key = strings.Replace(key, "[", "", 1)
    key = strings.Replace(key, "]", "", 1)

    return key
}

func NewFromNFA (nfa *NFA.NFA) *DFA {
    type NFATemp struct {
        Symbols map[string][]int
    }
    tmp_nodes :=  make([]NFATemp, nfa.Length())
    for i,_ := range tmp_nodes {
        tmp_nodes[i].Symbols = make(map[string][]int)
    }

    nfa_alphabet := []string{}

    // making big table for help
    for i, node := range nfa.Nodes() {
        if node.LeftSymbol != "" && node.Left != nil {
            var nodes []int
            nodes = append(nodes, node.Left.Id)
            walkFrom(node.Left, &nodes, "", false)
            tmp_nodes[i].Symbols[node.LeftSymbol] = nodes
            if !inStringArray(nfa_alphabet, node.LeftSymbol) {
                nfa_alphabet = append(nfa_alphabet, node.LeftSymbol)
            }
        }
        if node.RightSymbol != "" && node.Right != nil {
            var nodes []int
            nodes = append(nodes, node.Right.Id)
            walkFrom(node.Right, &nodes, "", false)
            tmp_nodes[i].Symbols[node.RightSymbol] = nodes
            if !inStringArray(nfa_alphabet, node.RightSymbol) {
                nfa_alphabet = append(nfa_alphabet, node.RightSymbol)
            }
        }
    }

    fmt.Println(nfa_alphabet)

    tmp := make(map[string][][]int)

    for _, node := range tmp_nodes {
        for symb, val := range node.Symbols {
            tmp[symb] = append(tmp[symb], val)
        }
    }

    fmt.Printf("%+v\n", tmp)

    alphabet := []string{}
    for i:='A'; i<'Z'; i++ {
        alphabet = append(alphabet, string(i))
    }

    var start_nodes []int
    start_nodes = append(start_nodes, nfa.Nodes()[0].Id)
    walkFrom(nfa.Nodes()[0], &start_nodes, "", false)


    rules := make(map[string]string)
    relations := make(map[string][]string)

    fmt.Printf("start_nodes: %#v\n", start_nodes)

    rules[makeKeyFromIntArray(start_nodes)] = alphabet[0]

    alphabet = alphabet[1:]

    is_fill := false
    for !is_fill {
        is_fill = true
        for rule_key, rule_symb := range rules {
            rule_ids := strings.Split(rule_key, " ")
            relations[rule_symb] = make([]string, len(nfa_alphabet))
            for ai, alphabet_symb := range nfa_alphabet {
                var ids, full_ids []int
                for _, sid := range rule_ids {
                    id, _ := strconv.Atoi(sid)
                    walkFrom(nfa.Nodes()[id], &ids, alphabet_symb, true)
                }

                if len(ids)==0 {
                    relations[rule_symb][ai] = "Error"
                    rules[""] = "Error"
                    continue
                }

                for _, id := range ids {
                    for _, id_arr := range tmp[alphabet_symb] {
                        if inIntArray(id_arr, id) {
                            full_ids = append(full_ids, id_arr...)
                        }
                    }
                }

                symb, exist := rules[makeKeyFromIntArray(full_ids)]
                if !exist {
                    rules[makeKeyFromIntArray(full_ids)] = alphabet[0]
                    symb = alphabet[0]
                    alphabet = alphabet[1:]
                    is_fill = false
                }

                relations[rule_symb][ai] = symb
            }
        }
    }

    final_states := make(map[string]bool, len(rules))
    for rule_ids, symb := range rules {
        final_states[symb] = false
        sids := strings.Split(rule_ids, " ")
        for _, sid := range sids {
            id, _ := strconv.Atoi(sid)
            if nfa.Nodes()[id] == nfa.Last() {
                final_states[symb] = true
            }
        }
    }

    dfa := DFA{}
    nodes := make(map[string]*Node)
    for _, symb := range rules {
        new_node := Node{}
        if final_states[symb] {
            new_node.IsFinal = true
        }
        dfa.addNode(&new_node)
        nodes[symb] = &new_node
    }
//
    for symb_from, symbs := range relations {
        for i, symb_to := range symbs {
            nodes[symb_from].Directions[nfa_alphabet[i]] = nodes[symb_to]
        }
    }

    fmt.Printf("relations: %v\n", nodes)

//    for srule_ids, rule_symbol := range rules {
//        rule_ids := strings.Split(srule_ids, " ")
//
//        for _, alphabet_symb := range nfa_alphabet {
//            var ids []int
//            for _, sid := range rule_ids {
//                id, _ := strconv.Atoi(sid)
//                walkFrom(nfa.Nodes()[id], &ids, alphabet_symb, false)
//            }
//
//            key := fmt.Sprintf("%v", ids)
//            key = strings.Replace(key, "[", "", 1)
//            key = strings.Replace(key, "]", "", 1)
//
//            relation_symb, exist := rules[key];
//            if !exist {
//                relation_symb = alphabet[0]
//                rules[key] = relation_symb
//                alphabet = alphabet[1:]
//            }
//
//            relations[rule_symbol] = append(relations[rule_symbol], relation_symb)
//        }
//    }


    fmt.Printf("rules: %#v\n", rules)
    fmt.Printf("relations: %#v\n", relations)

    return &dfa
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