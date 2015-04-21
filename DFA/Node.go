package DFA

import "github.com/hhh0pE/REtoNFA/NFA"

type Node struct {
	NFA.Node
	is_final bool
}
