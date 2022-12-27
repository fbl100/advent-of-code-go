package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/data-structures/grid"
	"github.com/alexchao26/advent-of-code-go/halp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type Piece struct {
	// ReferencePixels are relative to a (0,0) bottom/left
	ReferencePixels *grid.InfiniteGrid
	//####
	//
	//.#.
	//###
	//.#.
	//
	//..#
	//..#
	//###
	//
	//#
	//#
	//#
	//#
	//
	//##
	//##
}

func (p *Piece) Print() {
	halp.PrintInfiniteGridStringsCR_reversedRows(p.ReferencePixels.Coords, ".")
	fmt.Println("Min R: ", p.ReferencePixels.MinR)
	fmt.Println("Min C: ", p.ReferencePixels.MinC)
	fmt.Println("Max R: ", p.ReferencePixels.MaxR)
	fmt.Println("Max C: ", p.ReferencePixels.MaxC)
}

func (p *Piece) GetPositons(left, bottom int) [][2]int {
	ans := [][2]int{}

	for k, _ := range p.ReferencePixels.Coords {
		ans = append(ans, [2]int{k[0] + left, k[1] + bottom})
	}
	return ans

}

func NewPiece(pixels []string) *Piece {
	// pixels are in drawing order
	// so row[0] = top
	//    row[n-1] = bottom
	// cols are cols

	g := grid.NewInfiniteGridNoFloor()

	for r := 0; r < len(pixels); r++ {
		for c := 0; c < len(pixels[r]); c++ {
			if string(pixels[r][c]) == "#" {
				g.Put(c, len(pixels)-1-r, "#")
			}
		}
	}

	p := Piece{ReferencePixels: g}
	return &p
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
	_ = parsed

	//####
	//
	//.#.
	//###
	//.#.
	//
	//..#
	//..#
	//###
	//
	//#
	//#
	//#
	//#
	//
	//##
	//##

	p1 := NewPiece([]string{"####"})
	p2 := NewPiece([]string{
		" # ",
		"###",
		" # "})

	p3 := NewPiece([]string{
		"  #",
		"  #",
		"###"})
	p4 := NewPiece([]string{
		"#",
		"#",
		"#",
		"#"})

	p5 := NewPiece([]string{
		"##",
		"##"})

	pieces := []*Piece{p1, p2, p3, p4, p5}

	grid := grid.NewInfiniteGridNoFloor()

	// add the floor
	for c := 0; c < 7; c++ {
		grid.Put(c, 0, "-")
	}

	pieceIndex := 0
	moveIndex := 0
	bottom := grid.MaxR + 4
	left := 2
	stopCount := 0
	for stopCount < 2022 {
		// apply a move
		move := string(input[moveIndex])
		moveIndex = (moveIndex + 1) % len(input)

		coords := pieces[pieceIndex].GetPositons(left, bottom)
		switch move {
		case "<":
			left, bottom = Move(-1, 0, left, bottom, coords, grid)
			coords = pieces[pieceIndex].GetPositons(left, bottom)
		case ">":
			left, bottom = Move(1, 0, left, bottom, coords, grid)
			coords = pieces[pieceIndex].GetPositons(left, bottom)
		}

		// now try the horizontal move
		_, newBott := Move(0, -1, left, bottom, coords, grid)

		if newBott != bottom {
			// we were able to move
			bottom = newBott
		} else {
			// unable to move, add the piece to the grid
			Move(0, -1, left, bottom, coords, grid)
			coords = pieces[pieceIndex].GetPositons(left, bottom)
			for _, coord := range coords {
				grid.Put(coord[0], coord[1], "#")
			}

			stopCount++ // increment stop count
			//fmt.Println("Stop Count: ", stopCount)
			//halp.PrintInfiniteGridStringsCR_reversedRows(grid.Coords, ".")
			//fmt.Println(stopCount)

			pieceIndex = (pieceIndex + 1) % len(pieces)
			bottom = grid.MaxR + 4
			left = 2

		}
	}

	return grid.MaxR
}

func Move(horizontal, vertical, left, bottom int, coords [][2]int, g *grid.InfiniteGrid) (int, int) {
	for i := range coords {
		c := coords[i][0] + horizontal
		r := coords[i][1] + vertical
		px, _ := g.Get(c, r)

		if c < 0 || c >= 7 || px != "." {
			// cannot move
			return left, bottom
		}
	}
	return left + horizontal, bottom + vertical
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans string) {
	return input
}
