package main

import (
	_ "embed"
	"flag"
	"fmt"
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

type Grid [][]string

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

	g := makeGrid(maxCol, maxRow)

	for _, w := range parsed {
		drawWall(&g, w)
	}

	printGrid(g, 494, 503, 0, 9)
	println()

	done := false
	count := 0
	for !done {
		stopped := sand(&g, 500, 0)
		if stopped {
			count++
		} else {
			done = true
		}
		printGrid(g, 494, 503, 0, 9)
		println()
	}

	return count
}

func part2(input string) int {
	return 0
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

func makeGrid(cols, rows int) Grid {
	mx := make([][]string, cols+1)
	for c := range mx {
		mx[c] = make([]string, rows+1)
		for r := range mx[c] {
			mx[c][r] = "."
		}
	}
	return mx
}

func printGrid(grid Grid, startCol, endCol, startRow, endRow int) {

	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			fmt.Print(grid[c][r])
		}
		fmt.Println()
	}
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
				(*grid)[c][r] = "#"
			}
		}
	}
}

func sand(grid *Grid, startCol, startRow int) bool {
	c := startCol
	r := startRow

	for r < len((*grid)[0]) {
		// if there is space at r+1, increment r
		x, err := getChar(grid, c, r+1)
		if err != nil {
			return false
		} else if x == "." {
			r++
			continue
		}

		x, err = getChar(grid, c-1, r+1)
		if err != nil {
			return false
		} else if x == "." {
			c--
			r++
			continue
		}

		x, err = getChar(grid, c+1, r+1)
		if err != nil {
			return false
		} else if x == "." {
			c++
			r++
			continue
		}

		// we got here
		(*grid)[c][r] = "o"
		return true
	}
	panic("should never get here")
}

func getChar(grid *Grid, c, r int) (string, error) {
	if c < 0 {
		// off the left
		return "*", AbyssError{}
	}

	if c > len(*grid) {
		return "*", AbyssError{}
	}

	if r >= len((*grid)[0]) {
		return "*", AbyssError{}
	}

	return (*grid)[c][r], nil
}
