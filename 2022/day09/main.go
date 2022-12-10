package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/data-structures/vec2d"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"k8s.io/gengo/examples/set-gen/sets"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

var Moves = map[string]*vec2d.Vector2D{
	"R": {X: 1, Y: 0},
	"L": {X: -1, Y: 0},
	"U": {X: 0, Y: 1},
	"D": {X: 0, Y: -1},
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

func makeLocation(x int, y int) int {
	return x<<32 + y
}

func part1(input string) int {
	moves := parseInput(input)

	head := &vec2d.Vector2D{0, 0}
	tail := &vec2d.Vector2D{0, 0}

	tailLocations := sets.NewInt()
	tailLocations.Insert(makeLocation(0, 0))
	_ = tail
	for _, move := range moves {

		for i := 0; i < move.Amount; i++ {
			head = head.Add(move.Direction)
			tail2head := head.Sub(tail)
			if mathy.AbsInt(tail2head.X) > 1 || mathy.AbsInt(tail2head.Y) > 1 {
				tail2head = tail2head.Clip(-1, 1)
				tail = tail.Add(tail2head)
				tailLocations.Insert(makeLocation(tail.X, tail.Y))
			}
		}

	}
	return len(tailLocations)
}

func part2(input string) int {
	moves := parseInput(input)

	numKnots := 10

	knots := []*vec2d.Vector2D{}
	for i := 0; i < numKnots; i++ {
		knots = append(knots, &vec2d.Vector2D{0, 0})
	}

	tailLocations := sets.NewInt()
	tailLocations.Insert(makeLocation(0, 0))

	for _, move := range moves {

		for i := 0; i < move.Amount; i++ {

			// move knot 0
			knots[0] = knots[0].Add(move.Direction)

			for k := 1; k < numKnots; k++ {
				// move knot k based on the position of k-1
				k2km1 := knots[k-1].Sub(knots[k])
				if mathy.AbsInt(k2km1.X) > 1 || mathy.AbsInt(k2km1.Y) > 1 {
					k2km1 = k2km1.Clip(-1, 1)
					knots[k] = knots[k].Add(k2km1)
				}
			}

			tail := knots[numKnots-1]
			tailLocations.Insert(makeLocation(tail.X, tail.Y))
		}

	}
	return len(tailLocations)
}

type Move struct {
	Direction *vec2d.Vector2D
	Amount    int
}

func (m *Move) String() string {
	return fmt.Sprintf("Dir: %v, Amt: %v", m.Direction.String(), m.Amount)
}

func parseInput(input string) (ans []*Move) {
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " ")
		dir := Moves[tokens[0]]
		amt, _ := strconv.Atoi(tokens[1])
		move := Move{Direction: dir, Amount: amt}
		ans = append(ans, &move)
	}
	return ans
}
