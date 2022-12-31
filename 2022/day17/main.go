package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/data-structures/grid"
	"github.com/alexchao26/advent-of-code-go/halp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type State struct {
	StackHeight int
	NumBlocks   int
}

func (b State) Diff(a State) StateDiff {
	return StateDiff{
		DeltaStackHeight: b.StackHeight - a.StackHeight,
		DeltaBlockCount:  b.NumBlocks - a.NumBlocks,
	}
}

func (a StateDiff) Equals(b StateDiff) bool {
	return a.DeltaStackHeight == b.DeltaStackHeight && a.DeltaBlockCount == b.DeltaBlockCount
}

type StateDiff struct {
	DeltaStackHeight int
	DeltaBlockCount  int
}

type IndexPair [2]int

type StateMap map[IndexPair][]State

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

	pieces := makePieces()

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

func makePieces() []*Piece {

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
	return pieces

}

func makeGrid() *grid.InfiniteGrid {

	grid := grid.NewInfiniteGridNoFloor()

	// add the floor
	for c := 0; c < 7; c++ {
		grid.Put(c, 0, "-")
	}

	return grid
}

func part2(input string) int {

	pieces := makePieces()

	grid := makeGrid()

	idx := runFixedStopCount(grid, pieces, 100000, 0, 0)

	var toPieceIndex int
	var toMoveIndex int
	var stateDiff *StateDiff

	for k, v := range idx {

		if len(v) > 1 {
			dx, err := determineCycle(v)
			if err != nil {
				fmt.Println(k[0], ",", k[1], " --> ", err.Error())
			} else {
				fmt.Println(k[0], ",", k[1], " --> ", dx.DeltaBlockCount, "/", dx.DeltaStackHeight)
				if stateDiff == nil || dx.DeltaBlockCount > stateDiff.DeltaBlockCount {
					toPieceIndex = k[0]
					toMoveIndex = k[1]
					stateDiff = dx
					fmt.Println(k[0], ",", k[1], " --> ", dx)
				}
			}
		}

	}

	// rebuild
	grid = makeGrid()
	count := runToIndex(grid, pieces, 0, 0, toPieceIndex, toMoveIndex)

	fmt.Println("Dropped ", count, " blocks")
	goal := 1000000000000
	remaining := goal - count
	fmt.Println("Need ", remaining, " more")
	skipCycles := remaining / stateDiff.DeltaBlockCount
	skipHeight := skipCycles * stateDiff.DeltaStackHeight
	skipBlocks := skipCycles * stateDiff.DeltaBlockCount
	fmt.Println("Skipping ", skipCycles, " cycles and adding ", skipHeight, " to the stack")
	remainingBlocks := remaining % stateDiff.DeltaBlockCount
	fmt.Println("There are ", count+skipBlocks, " drops blocked, ", remainingBlocks, " remain")

	//currentHeight := grid.MaxR

	h1 := grid.MaxR
	runFixedStopCount(grid, pieces, remainingBlocks, toPieceIndex, toMoveIndex)

	h2 := grid.MaxR
	dh := h2 - h1
	fmt.Println("Added ", dh, " to the stack")
	ans := h1 + dh + skipHeight

	return ans
}

func determineCycle(v []State) (*StateDiff, error) {
	if len(v) > 2 {
		delta := v[1].Diff(v[0])
		for i := 2; i < len(v); i++ {
			dv := v[i].Diff(v[i-1])
			if !dv.Equals(delta) {
				return nil, errors.New("not a cycle")
			}
		}

		// check to see if it's an even skip to 1000000000000
		lastH := v[len(v)-1].NumBlocks
		remaining := 1000000000000 - lastH
		if remaining%delta.DeltaBlockCount == 0 {
			return &delta, nil
		} else {
			return nil, errors.New("Not a good candidate")
		}

	} else {
		return nil, errors.New("not enough data")
	}

}

func runFixedStopCount(grid *grid.InfiniteGrid, pieces []*Piece, count int, startPieceIndex int, startMoveIndex int) StateMap {

	// [piece index][move index] -> stack height after the piece stops
	retVal := StateMap{}

	pieceIndex := startPieceIndex
	moveIndex := startMoveIndex
	bottom := grid.MaxR + 4
	left := 2
	stopCount := 0
	for stopCount < count {
		// apply a move
		move := string(input[moveIndex])

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

			indexPair := IndexPair{pieceIndex, moveIndex}
			retVal[indexPair] = append(retVal[indexPair], State{
				StackHeight: grid.MaxR,
				NumBlocks:   stopCount,
			})

			pieceIndex = (pieceIndex + 1) % len(pieces)
			bottom = grid.MaxR + 4
			left = 2

		}
		moveIndex = (moveIndex + 1) % len(input)
	}

	return retVal
}

func runToIndex(grid *grid.InfiniteGrid, pieces []*Piece, fromPieceIndex, fromMoveIndex, toPieceIndex, toMoveIndex int) int {

	// [piece index][move index] -> stack height after the piece stops
	//heightIndex := map[[2]int][]int{}

	pieceIndex := fromPieceIndex
	moveIndex := fromMoveIndex
	bottom := grid.MaxR + 4
	left := 2
	stopCount := 0
	cycleCount := 0
	for {
		// apply a move
		move := string(input[moveIndex])

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

			if pieceIndex == toPieceIndex && moveIndex == toMoveIndex {
				cycleCount++
				if cycleCount == 100 {
					return stopCount
				}
			}

			pieceIndex = (pieceIndex + 1) % len(pieces)
			bottom = grid.MaxR + 4
			left = 2

		}
		moveIndex = (moveIndex + 1) % len(input)
	}

}

func parseInput(input string) (ans string) {
	return input
}
