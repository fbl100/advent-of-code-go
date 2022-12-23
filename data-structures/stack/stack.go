package stack

type StackNode[T any] struct {
	Data T
	Prev *StackNode[T]
}

type Stack[T any] struct {
	Head  *StackNode[T]
	Count int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Head:  nil,
		Count: 0,
	}
}

func (s *Stack[T]) Push(v T) {
	n := &StackNode[T]{
		Data: v,
		Prev: s.Head,
	}
	s.Head = n
	s.Count++
}

func (s *Stack[T]) Pop() T {
	if s.Count == 0 {
		panic("empty stack")
	}
	temp := s.Head
	s.Head = s.Head.Prev
	s.Count--
	return temp.Data

}
