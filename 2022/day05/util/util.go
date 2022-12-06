package util

import (
	"fmt"
	"strings"
)

type Move struct {
	Count int
	From  int
	To    int
}

type Stack struct {
	Values []string
}

func NewStack() *Stack {
	return &Stack{Values: []string{}}
}

func (s *Stack) Top() string {
	if len(s.Values) == 0 {
		return ""
	}
	return s.Values[len(s.Values)-1]
}

func (s *Stack) Push(r string) {
	// trick for prepending
	s.Values = append(s.Values, r)
}

func (s *Stack) Pop() string {
	if len(s.Values) == 0 {
		return ""
	}

	retVal := s.Values[len(s.Values)-1]
	s.Values = s.Values[:len(s.Values)-1]
	return retVal
}

func (s *Stack) String() string {
	retVal := strings.Join(s.Values, " : ")
	return retVal
}

type Board struct {
	Stacks []*Stack
}

func (b *Board) Apply(m *Move) {
	for i := 0; i < m.Count; i++ {
		r := b.Stacks[m.From-1].Pop()
		b.Stacks[m.To-1].Push(r)
	}
}

func (b *Board) Apply2(m *Move) {
	s := []string{}

	for i := 0; i < m.Count; i++ {
		r := b.Stacks[m.From-1].Pop()
		s = append([]string{r}, s...)
	}

	for i := 0; i < m.Count; i++ {
		b.Stacks[m.To-1].Push(s[i])
	}

}

func NewBoard(numStacks int) *Board {
	b := &Board{}
	for i := 0; i < numStacks; i++ {
		b.Stacks = append(b.Stacks, NewStack())
	}
	return b
}

func (b *Board) Strings() []string {
	retVal := []string{}
	for i := 0; i < len(b.Stacks); i++ {
		s := "[" + fmt.Sprint(i) + "] " + b.Stacks[i].String()
		retVal = append(retVal, s)
	}
	return retVal
}

func (b *Board) Print() {
	strings := b.Strings()
	for _, s := range strings {
		println(s)
	}
}

func (b *Board) Top() string {
	retVal := ""
	for i := 0; i < len(b.Stacks); i++ {
		retVal += b.Stacks[i].Top()
	}
	return retVal
}
