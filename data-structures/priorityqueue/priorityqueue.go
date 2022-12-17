package priorityqueue

// found on https://levelup.gitconnected.com/a-heap-of-go-generics-cd20f362a76
// required login, but it's the best example of generics I've seen

type Heap[T any] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{comp: comp}
}

func (h *Heap[T]) Len() int { return len(h.data) }

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up(h.Len() - 1)
}

func (h *Heap[T]) Pop() T {
	n := h.Len() - 1
	if n > 0 {
		h.swap(0, n)
		h.down(0, h.Len()-1)
	}
	v := h.data[n]
	h.data = h.data[0:n]
	return v
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) up(jj int) {
	for {
		i := parent(jj)
		if i == jj || !h.comp(h.data[jj], h.data[i]) {
			break
		}
		h.swap(i, jj)
		jj = i
	}
}

func (h *Heap[t]) down(i0, n int) bool {
	i := i0
	for {
		j1 := left(i)
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.comp(h.data[j2], h.data[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.comp(h.data[j], h.data[i]) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}

func (h *Heap[T]) check() {
	//fmt.Println("Checking Heap Property")
	lastParent := parent(h.Len() - 1)
	for i := 0; i <= lastParent; i++ {
		l := left(i)
		r := right(i)
		if l < h.Len() && !h.comp(h.data[i], h.data[l]) {
			panic("Left Child Bad")
		}
		if r < h.Len() && !h.comp(h.data[i], h.data[r]) {
			panic("Right Child Bad")
		}
	}
}

//
//func (h *Heap[T]) down() {
//	//n := h.Len()
//	h.downFromIndex(0)
//	//n := h.Len() - 1
//	//i1 := 0
//	//for {
//	//	j1 := left(i1)
//	//	if j1 >= n || j1 < 0 {
//	//		break
//	//	}
//	//	j := j1
//	//	j2 := right(i1)
//	//	if j2 < n && h.comp(h.data[j2], h.data[j1]) {
//	//		j = j2
//	//	}
//	//	if !h.comp(h.data[j], h.data[i1]) {
//	//		break
//	//	}
//	//	h.swap(i1, j)
//	//	i1 = j
//	//}
//}
//
//func (h *Heap[T]) downFromIndex(i int) {
//	n := h.Len()
//	i1 := i
//	for {
//		j1 := left(i1)
//		if j1 >= n || j1 < 0 {
//			break
//		}
//		j := j1
//		j2 := right(i1)
//		if j2 < n && h.comp(h.data[j2], h.data[j1]) {
//			j = j2
//		}
//		if !h.comp(h.data[j], h.data[i1]) {
//			break
//		}
//		h.swap(i1, j)
//		i1 = j
//	}
//}

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (h *Heap[T]) ReHeapify() {
	// put the elements of h in heap order, in-place
	n := (h.Len()) / 2
	for i := n - 1; i >= 0; i-- {
		h.down(i, h.Len())
	}
}

func (h *Heap[T]) Init(items []T) {
	h.data = append(h.data, items...)
	h.ReHeapify()
}

func (h *Heap[T]) Top() T {
	v := h.data[0]
	return v
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return (i * 2) + 1 }
func right(i int) int  { return left(i) + 1 }
