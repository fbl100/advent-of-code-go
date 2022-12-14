package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/data-structures/graph"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"golang.org/x/exp/slices"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type Edge struct {
	Weight      int
	TargetIndex int
}

type VertexType int

const (
	Start VertexType = 0
	End              = 1
	Other            = 2
)

type Vertex struct {
	Name  string
	Index int
	Value rune
	Type  VertexType
}

type Index struct {
	Row int
	Col int
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
	g := parseInputPart1(input)
	vertices := g.GetVertices()
	start := slices.IndexFunc(vertices, func(v *Vertex) bool { return v.Type == Start })
	end := slices.IndexFunc(vertices, func(v *Vertex) bool { return v.Type == End })

	x := g.ShortestPath(vertices[start], vertices[end])

	fmt.Printf("Best = %v\n", x)

	return x

}

func part2(input string) int {

	g := parseInputPart2(input)
	vertices := g.GetVertices()

	start := []*Vertex{}
	end := []*Vertex{}

	for _, v := range vertices {
		if v.Type == Start {
			start = append(start, v)
		} else if v.Type == End {
			end = append(end, v)
		}
	}

	x := g.DjikstraDistances(end[0])

	min := math.MaxInt
	for _, v := range start {
		// there are islands of a's that you cannot use as start points
		if x[v] != math.MaxInt {
			fmt.Printf("%+v - %v\n", v, x[v])
			min = mathy.Min(min, x[v])
		}
	}

	return min
}

func neighborIndices(row int, col int, numRows int, numCols int) (ans []Index) {

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

func parseInputPart1(input string) *graph.AdjacencyList[Vertex] {

	gg := graph.NewAdjacencyList[Vertex]()

	mx, rows, cols := parseGraph(input)

	vertices := createVerticesPart1(mx, rows, cols)

	for _, v := range vertices {
		gg.AddVertex(v)
	}

	addEdgesPart1(gg, rows, cols, mx, vertices)

	return gg
}

func parseInputPart2(input string) *graph.AdjacencyList[Vertex] {

	gg := graph.NewAdjacencyList[Vertex]()

	mx, rows, cols := parseGraph(input)

	vertices := createVerticesPart2(mx, rows, cols)

	for _, v := range vertices {
		gg.AddVertex(v)
	}

	addEdgesPart2(gg, rows, cols, mx, vertices)

	return gg
}

func createVerticesPart1(mx [][]rune, rows, cols int) (vertices []*Vertex) {

	isStart := func(r rune) bool {
		return r == 'S'
	}
	isEnd := func(r rune) bool {
		return r == 'E'
	}

	return createVertices(mx, rows, cols, isStart, isEnd)

}

func createVerticesPart2(mx [][]rune, rows, cols int) (vertices []*Vertex) {

	isStart := func(r rune) bool {
		return r == 'S' || r == 'a'
	}
	isEnd := func(r rune) bool {
		return r == 'E'
	}

	return createVertices(mx, rows, cols, isStart, isEnd)

}

func createVertices(mx [][]rune, rows, cols int, isStart func(r rune) bool, isEnd func(r rune) bool) (vertices []*Vertex) {

	//vertices := []*Vertex{}

	// find the start/end times and add the vertices
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := rowcol2index(r, c, cols)

			var v *Vertex
			if isStart(mx[r][c]) {
				mx[r][c] = 'a'
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v (START)", r, c, string(mx[r][c])),
					Index: i,
					Value: 'a',
					Type:  Start,
				}
			} else if isEnd(mx[r][c]) {
				mx[r][c] = 'z'
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v (END)", r, c, string(mx[r][c])),
					Index: i,
					Value: 'z',
					Type:  End,
				}
			} else {
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v", r, c, string(mx[r][c])),
					Index: i,
					Value: mx[r][c],
					Type:  Other,
				}
			}

			vertices = append(vertices, v)
		}
	}

	return vertices
}

func parseGraph(input string) (mx [][]rune, rows, cols int) {

	for _, line := range strings.Split(input, "\n") {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		mx = append(mx, row)
	}

	rows = len(mx)
	cols = len(mx[0])
	return mx, rows, cols
}

func addEdgesPart1(gg *graph.AdjacencyList[Vertex], rows, cols int, mx [][]rune, vertices []*Vertex) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			i := rowcol2index(r, c, cols)
			for _, index := range neighborIndices(r, c, rows, cols) {

				diff := mx[index.Row][index.Col] - mx[r][c]
				if diff <= 1 {
					targetIndex := rowcol2index(index.Row, index.Col, cols)
					gg.AddEdge(vertices[i], vertices[targetIndex], 1)
				}

			}
		}
	}

}

func addEdgesPart2(gg *graph.AdjacencyList[Vertex], rows, cols int, mx [][]rune, vertices []*Vertex) {
	// for the reverse path, we can step down at most one
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {

			// from is the current square
			from := mx[r][c]

			i := rowcol2index(r, c, cols)
			for _, index := range neighborIndices(r, c, rows, cols) {

				// to is the neighbor
				to := mx[index.Row][index.Col]

				// edge is if the neighbor is at most -1 from 'from'
				diff := to - from
				// if from is 20
				//    to is 19
				// to - from = -1

				if diff >= -1 {
					targetIndex := rowcol2index(index.Row, index.Col, cols)
					gg.AddEdge(vertices[i], vertices[targetIndex], 1)
				}

			}
		}
	}

}

func rowcol2index(row, col, numCols int) int {
	return row*numCols + col
}
