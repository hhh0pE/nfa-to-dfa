package DFA
import "fmt"


type Node struct {
    Id int
	Directions map[string]*Node
    IsFinal bool
}
//
func (n *Node) Name() string {
    return string('A'+n.Id)
}

func (n *Node) toString() string {

    var output string
    output += n.Name()+":\n"
    for symb, relation_node := range n.Directions {
        output += fmt.Sprintf("\t%s -> %s\n", symb, relation_node.Name())
    }
    return output
}
