package grid

import (
	"github.com/alexchao26/advent-of-code-go/mathy"
	"math"
)

type AxisRange struct {
	Min int
	Max int
}

func (a *AxisRange) Update(value int) {
	a.Min = mathy.Min(a.Min, value)
	a.Max = mathy.Max(a.Max, value)
}

func (a *AxisRange) Contains(value int) bool {
	return a.Min <= value && value <= a.Max
}

func NewAxis() *AxisRange {
	return &AxisRange{
		Min: math.MaxInt,
		Max: math.MinInt,
	}
}

type Coord3d [3]int

func (c Coord3d) Neighbors() []Coord3d {
	ans := []Coord3d{
		{c[0] + 1, c[1], c[2]},
		{c[0] - 1, c[1], c[2]},
		{c[0], c[1] + 1, c[2]},
		{c[0], c[1] - 1, c[2]},
		{c[0], c[1], c[2] + 1},
		{c[0], c[1], c[2] - 1},
	}
	return ans
}

type InfiniteCube struct {
	XAxis   *AxisRange
	YAxis   *AxisRange
	ZAxis   *AxisRange
	Default string
	Coords  map[Coord3d]string
}

func NewInfiniteCube(blank string) *InfiniteCube {
	ans := InfiniteCube{
		XAxis:   NewAxis(),
		YAxis:   NewAxis(),
		ZAxis:   NewAxis(),
		Default: blank,
		Coords:  map[Coord3d]string{},
	}
	return &ans
}

func (c *InfiniteCube) Contains(x, y, z int) bool {
	return c.XAxis.Contains(x) && c.YAxis.Contains(y) && c.ZAxis.Contains(z)
}

func (c *InfiniteCube) ContainsCoord(coord Coord3d) bool {
	return c.Contains(coord[0], coord[1], coord[2])
}

func (c *InfiniteCube) GetWithCoord(coord Coord3d) (string, error) {
	return c.Get(coord[0], coord[1], coord[2])
}

func (c *InfiniteCube) Get(x, y, z int) (string, error) {

	// out of bounds, return the default character and an error
	// the user can decide if they care about the error or not
	if !c.Contains(x, y, z) {
		return c.Default, AbyssError{}
	}

	// we are within the bounds of the grid, so go ahead and return the value if we have one, or '.' if not
	// in this case, we don't return an error because it's technically in bounds
	retVal, ok := c.Coords[Coord3d{x, y, z}]
	if ok {
		return retVal, nil
	} else {
		return c.Default, nil
	}

}

func (c *InfiniteCube) Put(x, y, z int, s string) {

	c.XAxis.Update(x)
	c.YAxis.Update(y)
	c.ZAxis.Update(z)

	c.Coords[Coord3d{x, y, z}] = s
}

func (c *InfiniteCube) PutWithCoord(coord Coord3d, s string) {
	c.XAxis.Update(coord[0])
	c.YAxis.Update(coord[1])
	c.ZAxis.Update(coord[2])

	c.Coords[coord] = s
}
