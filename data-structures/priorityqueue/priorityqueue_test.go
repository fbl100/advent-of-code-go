package priorityqueue

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := NewHeap( func(a,b int) bool { return a < b})

	for i := 0; i < 100; i++ {
		h.Push(rand.Int() % 5000)
	}
	//h.Push(5)
	//h.Push(3)
	//h.Push(7)

	// every item should be less than the prior
	x := h.Top()

	for h.Len() > 0 {
		if h.Top() < x {
			panic("que order messed up")
		}
		x = h.Pop()
		fmt.Println(x)
	}

}

type Person struct {
	Name string
	Age int
}



func TestHeap_Push_PeoplePointers(t *testing.T) {

	people := []*Person {
		{"Frank", 47},
		{"Sarah", 40},
		{"Steph", 47},
		{"Ben", 18},
		{"Abby", 5},
		{"Jacob", 2},
	}

	h := NewHeap(func(a, b *Person) bool {
		return a.Age < b.Age
	})

	h.Init(people)
	//for _,p := range(people) {
	//	h.Push(p)
	//}

	// every item should be less than the prior
	x := h.Top()

	for h.Len() > 0 {
		if h.Top().Age < x.Age {
			panic("que order messed up")
		}
		x = h.Pop()
		fmt.Println(x.Name)
	}

}