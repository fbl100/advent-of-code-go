package priorityqueue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap_Int(t *testing.T) {
	h := NewHeap( func(a,b int) bool { return a > b})

	h.Push(5)
	h.Push(3)
	h.Push(7)

	for h.Len() > 0 {
		fmt.Println(h.Pop())
	}

	heap.Init()
}