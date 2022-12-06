package util

import "testing"

func TestStack(t *testing.T) {

	s := NewStack()
	s.Push("a")
	s.Push("b")

	if s.Top() != "b" {
		t.Fatalf("s.Top() expected 'b'")
	}

	if s.Pop() != "b" {
		t.Fatalf("s.Pop() expected 'b'")
	}

	if s.Pop() != "a" {
		t.Fatalf("s.Pop() expected 'a'")
	}

	if s.Pop() != "" {
		t.Fatalf("s.Pop() expected empty")
	}

}

func TestBoard(t *testing.T) {
	b := NewBoard(3)
	b.Stacks[0].Push("A")
	b.Stacks[0].Push("B")
	b.Stacks[0].Push("C")

	m := Move{
		Count: 2,
		From:  0,
		To:    2,
	}

	b.Print()

	b.Apply(&m)

	b.Print()

}
