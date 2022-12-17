package priorityqueue

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := NewHeap(func(a, b int) bool { return a < b })

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
	Age  int
}

func TestHeap_Push_PeoplePointers(t *testing.T) {

	people := []*Person{
		{"Frank", 47},
		{"Sarah", 40},
		{"Steph", 47},
		{"Ben", 18},
		{"Abby", 5},
		{"Jacob", 2},
	}

	h := NewHeap(func(a, b *Person) bool {
		//if a.Age <= b.Age {
		//	fmt.Printf("%v (age %v) <= %v (age (%v)\n", a.Name, a.Age, b.Name, b.Age)
		//} else {
		//	fmt.Printf("%v (age %v) > %v (age (%v)\n", a.Name, a.Age, b.Name, b.Age)
		//}
		return a.Age <= b.Age
	})

	//           [0]
	//       [1]     [2]
	//     [3] [4] [5] [6]
	h.Init(people)
	h.check()
	//
	//for _, p := range people {
	//	p.Age = rand.Int()
	//}
	//
	//h.ReHeapify()
	//h.check()

	// every item should be less than the prior
	x := h.Top()

	for h.Len() > 0 {
		if h.Top().Age < x.Age {
			panic("que order messed up")
		}
		x = h.Pop()
		h.check()
		fmt.Println(x.Name)
	}

}
