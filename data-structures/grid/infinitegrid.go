package grid

import (
	"github.com/alexchao26/advent-of-code-go/mathy"
	"math"
)

type InfiniteGrid struct {
	MinC    int
	MaxC    int
	MinR    int
	MaxR    int
	Default string
	Floor   int
	Coords  map[[2]int]string
}

func NewInfiniteGridNoFloor() *InfiniteGrid {
	return NewInfiniteGrid(math.MaxInt)
}

func NewInfiniteGrid(floor int) *InfiniteGrid {
	g := InfiniteGrid{
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

func (g *InfiniteGrid) Get(c, r int) (string, error) {

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

func (g *InfiniteGrid) Put(c, r int, s string) {

	g.MinC = mathy.Min(g.MinC, c)
	g.MaxC = mathy.Max(g.MaxC, c)
	g.MinR = mathy.Min(g.MinR, r)
	g.MaxR = mathy.Max(g.MaxR, r)

	g.Coords[[2]int{c, r}] = s
}
