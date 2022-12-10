package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

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

type Tree struct {
	Height int
	Seen   bool
}

func scanRowLR(rowIndex int, grid [][]*Tree) int {

	count := 0
	max := -1
	for c := 0; c < len(grid[rowIndex]); c++ {
		if grid[rowIndex][c].Height > max {
			// count if not seen
			if grid[rowIndex][c].Seen == false {
				count++
				grid[rowIndex][c].Seen = true
			}
			max = grid[rowIndex][c].Height
		}
	}
	return count
}

func scanRowRL(rowIndex int, grid [][]*Tree) int {

	count := 0
	max := -1
	for c := len(grid[rowIndex]) - 1; c >= 0; c-- {
		if grid[rowIndex][c].Height > max {
			// count if not seen
			if grid[rowIndex][c].Seen == false {
				count++
				grid[rowIndex][c].Seen = true
			}
			max = grid[rowIndex][c].Height
		}
	}
	return count
}

func scanRowTB(colIndex int, grid [][]*Tree) int {

	count := 0
	max := -1
	for r := 0; r < len(grid); r++ {
		if grid[r][colIndex].Height > max {
			// count if not seen
			if grid[r][colIndex].Seen == false {
				count++
				grid[r][colIndex].Seen = true
			}
			max = grid[r][colIndex].Height
		}
	}
	return count
}

func scanRowBT(colIndex int, grid [][]*Tree) int {

	count := 0
	max := -1
	for r := len(grid) - 1; r >= 0; r-- {
		if grid[r][colIndex].Height > max {
			// count if not seen
			if grid[r][colIndex].Seen == false {
				count++
				grid[r][colIndex].Seen = true
			}
			max = grid[r][colIndex].Height
		}
	}
	return count
}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	rows := len(parsed)
	cols := len(parsed[0])

	sum := 0
	for r := 0; r < rows; r++ {
		sum += scanRowLR(r, parsed)
		sum += scanRowRL(r, parsed)
	}

	for c := 0; c < cols; c++ {
		sum += scanRowBT(c, parsed)
		sum += scanRowTB(c, parsed)
	}

	return sum
}

func part2LookLeft(row int, col int, grid [][]*Tree) int {
	h := grid[row][col].Height
	r := row
	c := col - 1
	// you always see the tree next to you
	count := 0

	for c >= 0 {
		if grid[r][c].Height >= h {
			return count + 1
		} else {
			count++
			c--
		}
	}

	return count
}

func part2LookRight(row int, col int, grid [][]*Tree) int {
	h := grid[row][col].Height
	r := row
	c := col + 1
	// you always see the tree next to you
	count := 0

	for c < len(grid[row]) {
		if grid[r][c].Height >= h {
			return count + 1
		} else {
			count++
			c++
		}
	}
	// if we got here, then we can see to the edge
	return count
}

func part2LookUp(row int, col int, grid [][]*Tree) int {
	h := grid[row][col].Height
	r := row - 1
	c := col
	// you always see the tree next to you
	count := 0

	for r >= 0 {
		if grid[r][c].Height >= h {
			return count + 1
		} else {
			count++
			r--
		}
	}

	return count
}

func part2LookDown(row int, col int, grid [][]*Tree) int {
	h := grid[row][col].Height
	r := row + 1
	c := col
	// you always see the tree next to you
	count := 0

	for r < len(grid) {
		if grid[r][c].Height >= h {
			return count + 1
		} else {
			count++
			r++
		}
	}

	return count
}

func part2(input string) int {
	grid := parseInput(input)

	rows := len(grid)
	cols := len(grid[0])

	best := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			ans := 1
			ans *= part2LookLeft(r, c, grid)
			ans *= part2LookRight(r, c, grid)
			ans *= part2LookUp(r, c, grid)
			ans *= part2LookDown(r, c, grid)

			if ans > best {
				best = ans
			}
		}
	}

	//
	//col := 98
	//for row := 0; row < len(grid); row++ {
	//	x := part2LookRight(row, col, grid)
	//	println("[", row+1, "] = [", grid[row][col].Height, "] = ", x)
	//}
	//
	//row := 5
	//for col := 0; col < len(grid[0]); col++ {
	//	x := part2LookUp(row, col, grid)
	//	println("col[", col+1, "] = [", grid[row][col].Height, "] = ", x)
	//}
	return best
}

func parseInput(input string) (ans [][]*Tree) {

	for _, line := range strings.Split(input, "\n") {
		row := []*Tree{}
		for i := 0; i < len(line); i++ {
			h, _ := strconv.Atoi(string(line[i]))
			t := Tree{Height: h, Seen: false}
			row = append(row, &t)
		}
		ans = append(ans, row)
	}
	return ans
}
