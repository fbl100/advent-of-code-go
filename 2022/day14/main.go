package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/halp"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"math"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type Pair struct {
	Col int
	Row int
}

type Wall struct {
	Pairs []Pair
}

type Grid struct {
	MinC    int
	MaxC    int
	MinR    int
	MaxR    int
	Default string
	Floor   int
	Coords  map[[2]int]string
}

func NewGrid(floor int) *Grid {
	g := Grid{
		MinC:    math.MaxInt,
		MaxC:    math.MinInt,
		MinR:    math.MaxInt,
		MaxR:    math.MinInt,
		Floor:   floor,
		Default: ".",
		Coords:  map[[2]int]string{},
	}
	return &g
}

func (g *Grid) Get(c, r int) (string, error) {

	if r >= g.Floor {
		return "#", AbyssError{}
	}

	// out of bounds, return the default character and an error
	// the user can decide if they care about the error or not
	if c < g.MinC || c > g.MaxC || r < g.MinR || r > g.MaxR {
		return g.Default, AbyssError{}
	}

	// we are within the bounds of the grid, so go ahead and return the value if we have one, or '.' if not
	// in this case, we don't return an error because it's technically in bounds
	retVal, ok := g.Coords[[2]int{c, r}]
	if ok {
		return retVal, nil
	} else {
		return g.Default, nil
	}

}

func (g *Grid) Put(c, r int, s string) {

	g.MinC = mathy.Min(g.MinC, c)
	g.MaxC = mathy.Max(g.MaxC, c)
	g.MinR = mathy.Min(g.MinR, r)
	g.MaxR = mathy.Max(g.MaxR, r)

	g.Coords[[2]int{c, r}] = s
}

type AbyssError struct {
}

func (m AbyssError) Error() string {
	return "Fell off the world"
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
	parsed := parseInput(input)
	maxCol := math.MinInt
	maxRow := math.MinInt

	for _, w := range parsed {
		for _, p := range w.Pairs {
			maxCol = mathy.Max(maxCol, p.Col)
			maxRow = mathy.Max(maxRow, p.Row)
		}
	}

	fmt.Println(maxCol, maxRow)

	g := NewGrid(math.MaxInt)

	for _, w := range parsed {
		drawWall(g, w)
	}

	done := false
	count := 0
	for !done {
		stopped := sand(g, 500, 0)
		if stopped {
			count++
		} else {
			done = true
		}

	}
	println()
	halp.PrintInfiniteGridStrings(g.Coords, " ")
	return count
}

func part2(input string) int {
	parsed := parseInput(input)
	maxCol := math.MinInt
	maxRow := math.MinInt

	for _, w := range parsed {
		for _, p := range w.Pairs {
			maxCol = mathy.Max(maxCol, p.Col)
			maxRow = mathy.Max(maxRow, p.Row)
		}
	}

	g := NewGrid(maxRow + 2)

	for _, w := range parsed {
		drawWall(g, w)
	}

	halp.PrintInfiniteGridStrings(g.Coords, ".")

	done := false
	count := 0
	for !done {
		stopped := sand_part2(g, 500, 0)
		if stopped {
			count++
		} else {
			done = true
		}
	}

	fmt.Println()

	halp.PrintInfiniteGridStrings(g.Coords, ".")

	return count + 1
}

func parseInput(input string) (ans []Wall) {
	for _, line := range strings.Split(input, "\n") {
		ans = append(ans, parseWall(line))
	}
	return ans
}

func parseWall(input string) (ans Wall) {

	for _, pair := range strings.Split(input, " -> ") {
		xy := strings.Split(pair, ",")
		col, _ := strconv.Atoi(xy[0])
		row, _ := strconv.Atoi(xy[1])
		ans.Pairs = append(ans.Pairs, Pair{
			Col: col,
			Row: row,
		})
	}
	return ans
}

func drawWall(grid *Grid, wall Wall) {
	for i := 0; i < len(wall.Pairs)-1; i++ {
		this := wall.Pairs[i]
		next := wall.Pairs[i+1]

		startCol := mathy.Min(this.Col, next.Col)
		endCol := mathy.Max(this.Col, next.Col)

		startRow := mathy.Min(this.Row, next.Row)
		endRow := mathy.Max(this.Row, next.Row)

		for c := startCol; c <= endCol; c++ {
			for r := startRow; r <= endRow; r++ {
				grid.Put(c, r, "#")
			}
		}
	}
}

func sand_part2(grid *Grid, startCol, startRow int) bool {
	c := startCol
	r := startRow

	for {
		// if there is space at r+1, increment r
		x, _ := grid.Get(c, r+1)
		if x == "." {
			r++
			continue
		}

		x, _ = grid.Get(c-1, r+1)
		if x == "." {
			c--
			r++
			continue
		}

		x, _ = grid.Get(c+1, r+1)
		if x == "." {
			c++
			r++
			continue
		}

		if c == startCol && r == startRow {
			// we blocked the hole
			return false
		}

		// we hit something, so place the grain of sand
		grid.Put(c, r, "o")
		return true
	}
	panic("should never get here")
}

func sand(grid *Grid, startCol, startRow int) bool {
	c := startCol
	r := startRow

	maxRow := grid.MaxR

	for {
		// if there is space at r+1, increment r
		x, err := grid.Get(c, r+1)
		if err != nil && r+1 > maxRow {
			return false
		} else if x == "." {
			r++
			continue
		}

		x, err = grid.Get(c-1, r+1)
		if err != nil && r+1 > maxRow {
			return false
		} else if x == "." {
			c--
			r++
			continue
		}

		x, err = grid.Get(c+1, r+1)
		if err != nil && r+1 > maxRow {
			return false
		} else if x == "." {
			c++
			r++
			continue
		}

		// we got here
		grid.Put(c, r, "o")
		return true
	}
	panic("should never get here")
}
