package graph

import (
	"github.com/alexchao26/advent-of-code-go/data-structures/priorityqueue"
	"math"
)

type VertexValue[V any] struct {
	Vertex *V
	Value  int
}

type WeightedEdge[V any] struct {
	Vertex *V
	Weight int
}

type AdjacencyList[V any] struct {
	Vertices map[*V][]WeightedEdge[V]
}

func NewAdjacencyList[V any]() *AdjacencyList[V] {
	retVal := AdjacencyList[V]{}
	retVal.Vertices = map[*V][]WeightedEdge[V]{}
	return &retVal
}

func (g *AdjacencyList[V]) AddVertex(vertex *V) {
	g.Vertices[vertex] = []WeightedEdge[V]{}
}

func (g *AdjacencyList[V]) AddEdge(from *V, to *V, weight int) {
	g.Vertices[from] = append(g.Vertices[from], WeightedEdge[V]{
		Vertex: to,
		Weight: weight,
	})
}

func (g *AdjacencyList[V]) GetNeighbors(v *V) []WeightedEdge[V] {
	return g.Vertices[v]
}

func (g *AdjacencyList[V]) GetVertices() []*V {
	retVal := []*V{}
	for v := range g.Vertices {
		retVal = append(retVal, v)
	}
	return retVal
}

func (g *AdjacencyList[V]) ShortestPath(from *V, to *V) int {

	dist := map[*V]int{}
	q := priorityqueue.NewHeap(func(a, b *V) bool {
		return dist[a] < dist[b]
	})
	for v := range g.Vertices {
		dist[v] = math.MaxInt
	}
	dist[from] = 0

	q.Init(g.GetVertices())

	// 	h := NewHeap(func(a, b int) bool { return a < b })

	for q.Len() > 0 {

		next := q.Pop()
		//fmt.Printf("Relaxing %+v with dist %v\n", *next, dist[next])
		adj := g.GetNeighbors(next)
		for _, e := range adj {

			d := dist[next] + e.Weight
			if d < dist[e.Vertex] {
				dist[e.Vertex] = d
			}

		}
		q.ReHeapify()
	}

	return dist[to]

}

func (g *AdjacencyList[V]) DjikstraDistances(start *V) map[*V]int {
	dist := map[*V]int{}
	prev := map[*V]*V{}

	q := priorityqueue.NewHeap(func(a, b *V) bool {
		return dist[a] <= dist[b]
	})

	for v := range g.Vertices {
		dist[v] = math.MaxInt
		prev[v] = nil
	}

	dist[start] = 0
	q.Push(start)

	for q.Len() > 0 {

		next := q.Pop()

		//fmt.Printf("Relaxing %+v with dist %v\n", *next, dist[next])
		adj := g.GetNeighbors(next)
		for _, e := range adj {

			d := dist[next] + e.Weight
			if d < 0 {
				panic("overflow?!?")
			}
			if d < dist[e.Vertex] {
				dist[e.Vertex] = d
				prev[e.Vertex] = next
				q.Push(e.Vertex)
			}
		}
	}

	return dist

}
