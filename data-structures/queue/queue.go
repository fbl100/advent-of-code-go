package queue

type QueueNode[T any] struct {
	Data T
	Next *QueueNode[T]
}

type Queue[T any] struct {
	Head  *QueueNode[T]
	Tail  *QueueNode[T]
	Count int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Head:  nil,
		Tail:  nil,
		Count: 0,
	}
}

func (s *Queue[T]) Push(v T) {
	n := &QueueNode[T]{
		Data: v,
		Next: nil,
	}
	if s.Count == 0 {
		s.Head = n
		s.Tail = n
	} else {
		s.Tail.Next = n
		s.Tail = n
	}

	s.Count++
}

func (s *Queue[T]) Pop() T {
	if s.Count == 0 {
		panic("empty queue")
	} else if s.Count == 1 {
		temp := s.Head
		s.Head = nil
		s.Tail = nil
		s.Count = 0
		return temp.Data
	} else {
		temp := s.Head
		s.Head = s.Head.Next
		s.Count--
		return temp.Data
	}

}
