package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/data-structures/graph"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input_small.txt
var input string

type Edge struct {
	Weight      int
	TargetIndex int
}

type Vertex struct {
	Name  string
	Index int
	Value rune
}

type VertexNode struct {
	Vertex *Vertex
	Dist   int
}

type Graph struct {
	Vertices map[int]*Vertex
	Start    int
	End      int
}

type Index struct {
	Row int
	Col int
}

type HeapNode struct {
	vertex *Vertex
	dist   *map[int]int
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
	g, start, end := parseInput(input)

	x := g.ShortestPath(start, end)

	fmt.Printf("Best = %v\n", x)
	
	return x
}

func part2(input string) int {
	g, _, end := parseInput(input)

	dist := g.ShortestPathTree(end)

	for _, v := range g.GetVertices() {
		if v.Value == rune('a') {
			d := dist[v]
			fmt.Println(d)
		}

	}
	return 0
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

func parseInput(input string) (*graph.AdjacencyList[Vertex], *Vertex, *Vertex) {

	gg := graph.NewAdjacencyList[Vertex]()

	// parse the runes
	mx := [][]rune{}

	for _, line := range strings.Split(input, "\n") {
		row := []rune{}
		for _, c := range line {
			row = append(row, c)
		}
		mx = append(mx, row)
	}

	rows := len(mx)
	cols := len(mx[0])

	var start *Vertex
	var end *Vertex

	vertices := []*Vertex{}

	// find the start/end times and add the vertices
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			i := rowcol2index(r, c, cols)

			var v *Vertex
			if mx[r][c] == rune('S') {
				mx[r][c] = 'a'
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v (START)", r, c, string(mx[r][c])),
					Index: i,
					Value: 'a',
				}

				start = v
			} else if mx[r][c] == rune('E') {
				mx[r][c] = 'z'
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v (END)", r, c, string(mx[r][c])),
					Index: i,
					Value: 'z',
				}
				end = v
			} else {
				v = &Vertex{
					Name:  fmt.Sprintf("r=%v c=%v char=%v", r, c, string(mx[r][c])),
					Index: i,
					Value: mx[r][c],
				}
			}

			vertices = append(vertices, v)

		}
	}

	for _, v := range vertices {
		gg.AddVertex(v)
	}

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

	return gg, start, end
}

func rowcol2index(row, col, numCols int) int {
	return row*numCols + col
}
