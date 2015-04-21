package DFA

type DFA struct {
	begin, end *DFA
	nodes      []*Node
}
