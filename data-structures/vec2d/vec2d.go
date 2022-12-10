package vec2d

import (
	"fmt"
	"github.com/alexchao26/advent-of-code-go/mathy"
)

type Vector2D struct {
	X int
	Y int
}

func (a *Vector2D) Add(b *Vector2D) *Vector2D {
	return &Vector2D{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a *Vector2D) Sub(b *Vector2D) *Vector2D {
	return &Vector2D{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a *Vector2D) Negate() *Vector2D {
	return &Vector2D{
		X: -a.X,
		Y: -a.Y,
	}
}

// clips all values between lo and hi
func (a *Vector2D) Clip(lo int, hi int) *Vector2D {
	return &Vector2D{
		X: mathy.ClipInt(a.X, lo, hi),
		Y: mathy.ClipInt(a.Y, lo, hi),
	}
}

func (v *Vector2D) String() string {
	return fmt.Sprintf("{ X: %v, Y: %v }", v.X, v.Y)
}
