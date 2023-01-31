package awesomeProject1

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top  *Node
	Size int
}

func (s *Stack) Push(value int) {
	node := &Node{value, nil}
	node.Next = s.Top
	s.Top = node
	s.Size++
}

func (s *Stack) Pop() {
	if s.Top == nil {
		return
	}
	s.Top = s.Top.Next
	s.Size--
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) getSize() int {
	return s.Size
}

func (s *Stack) toString() {
	for s.Size != 0 {
		fmt.Println(s.Top.Value)
		s.Pop()
	}
}
func (s *Stack) RevtoString() {
	st := Stack{}
	for s.Size != 0 {
		st.Push(s.Top.Value)
		s.Pop()
	}
	for st.Size != 0 {
		fmt.Println(st.Top.Value)
		st.Pop()
	}
}
func (s *Stack) Increment(value int) {
	for s.Size != 0 {
		s.Top.Value += value
		s.Pop()
	}
}
