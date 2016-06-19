package stack

// node for linked list
type Node struct {
	value interface{}
	prev  *Node
	next  *Node
}

func NewNode(v interface{}) *Node {
	return &Node{
		value: v,
		prev:  nil,
		next:  nil,
	}
}

// set prev
func (n *Node) Prev(node *Node) {
	n.prev = node
}

// set next
func (n *Node) Next(node *Node) {
	n.next = node
}

type Stack struct {
	size  int
	first *Node
	last  *Node
}

// make a new stack given the first value
func New(v interface{}) *Stack {
	// root doesn't have a value
	n := NewNode(nil)
	return &Stack{
		size:  0,
		first: n,
		last:  n,
	}
}

// push stack
func (s *Stack) Push(v interface{}) {
	nn := NewNode(v)
	nn.Prev(s.last)
	nn.Next(s.first)
	s.first.Prev(nn)
	s.last.Next(nn)
	s.last = nn
	s.size++
}

// pop stack
func (s *Stack) Pop() {
	if s.size == 0 {
		return
	}
	ln := s.last.prev
	ln.Next(s.first)
	s.first.Prev(ln)
	s.last = ln
	s.size--
}
