package main

import (
	"container/heap"
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input_small.txt
var input string

type Edge struct {
	Weight int
	TargetIndex int
}

type Vertex struct {
	Index int
	Value rune
	Edges []*Edge
}





type Graph struct {
	Vertices map[int]*Vertex
	Start int
	End int
}

type Index struct {
	Row int
	Col int
}

type HeapNode struct {
	vertex *Vertex
	dist *map[int]int

}



func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	g := parseInput(input)

	dist := map[int]int{}
	prev := map[int]int{}

	Q := []HeapNode{}

	for _, v := range(g.Vertices) {
		dist[v.Index] = math.MaxInt
		Q = append(Q, HeapNode{
			vertex: v,
			dist:   &dist,
		})
	}

	heap.Init(Q)

	dist[g.Start] = 0

	for Q.Length() > 0 {
	}

	//function Dijkstra(Graph, source):
	//2
	//3      for each vertex v in Graph.Vertices:
	//4          dist[v] ← INFINITY
	//5          prev[v] ← UNDEFINED
	//6          add v to Q
	//7      dist[source] ← 0
	//8
	//9      while Q is not empty:
	//10          u ← vertex in Q with min dist[u]
	//11          remove u from Q
	//12
	//13          for each neighbor v of u still in Q:
	//14              alt ← dist[u] + Graph.Edges(u, v)
	//15              if alt < dist[v]:
	//16                  dist[v] ← alt
	//17                  prev[v] ← u
	//18
	//19      return dist[], prev[]
	// find 'S' and 'E'

	_ = parsed

	return 0
}


func part2(input string) int {
	return 0
}

func neighborIndices(row int, col int, numRows int, numCols int) (ans []Index){

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if mathy.AbsInt(dr) == mathy.AbsInt(dc) {
				continue
			}
			if dr == 0 && dc == 0 {
				continue
			}
			r := row + dr
			c := col + dc

			if r < 0 || r >= numRows {
				continue
			}

			if c < 0 || c >= numCols {
				continue
			}

			ans = append(ans, Index{
				Row: r,
				Col: c,
			})
		}
	}
	return ans

}


func parseInput(input string) (g *Graph) {

	g = &Graph{
		Vertices: map[int]*Vertex{},
	}

	// parse the runes
	mx := [][]rune{}

	for _, line := range strings.Split(input, "\n") {
		row := []rune{}
		for _,c := range(line) {
			row = append(row, c)
		}
		mx = append(mx, row)
	}

	rows := len(mx)
	cols := len(mx[0])

	start := -1
	end := -1

	// find the start/end times
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := r*cols + c
			if mx[r][c] == rune('S') {
				start = i
				mx[r][c] = 'a'
			}

			if mx[r][c] == rune('E') {
				end = i
				mx[r][c] = 'z'
			}

		}
	}

	g.Start = start
	g.End = end

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			i := r*cols + c

			v := Vertex{
				Index: i,
				Value: mx[r][c],
				Edges: nil,
			}


			for _, index := range neighborIndices(r,c, rows, cols) {

				weight := mx[index.Row][index.Col] - mx[r][c]
				targetIndex := index.Row * cols + index.Col

				e := Edge{
					Weight: int(weight),
					TargetIndex: targetIndex,
				}

				v.Edges = append(v.Edges, &e)
			}

			g.Vertices[i] = &v

		}
	}

	return g
}
