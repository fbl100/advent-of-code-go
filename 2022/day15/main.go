package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/data-structures/grid"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"regexp"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type Record struct {
	SensorX int
	SensorY int
	BeaconX int
	BeaconY int
}

func (r Record) ManhattanDist() int {
	return mathy.ManhattanDistance(r.SensorX, r.SensorY, r.BeaconX, r.BeaconY)
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
	records := parseInput(input)

	grid := grid.NewInfiniteGridNoFloor()
	_ = grid
	for _, rec := range records {
		grid.Put(rec.SensorX, rec.SensorY, "S")
		grid.Put(rec.BeaconX, rec.BeaconY, "B")
	}

	target_y := 2000000
	println("populating grid")
	count := 0
	for _, rec := range records {
		m := rec.ManhattanDist()

		minY := rec.SensorY - m
		maxY := rec.SensorY + m

		if target_y >= minY && target_y <= maxY {
			for x := rec.SensorX - m; x <= rec.SensorX+m; x++ {
				if mathy.ManhattanDistance(rec.SensorX, rec.SensorY, x, target_y) <= m {
					curr, _ := grid.Get(x, target_y)
					if curr == "." {
						grid.Put(x, target_y, "#")
						count++
					} else {
						//fmt.Println("Skipped ", curr, " at x=", x, ", y=", target_y)
					}
				}
			}
		}
	}
	return count
}

func part2(input string) int {

	records := parseInput(input)

	allPoints := map[[2]int]bool{}

	for _, r := range records {
		points := manhattanRing(r.SensorX, r.SensorY, r.ManhattanDist()+1)

		for _, p := range points {
			if p[0] >= 0 && p[1] >= 0 && p[0] <= 4000000 && p[1] <= 4000000 {
				if visibleCount(p, records) == 0 {
					return p[0]*4000000 + p[1]
				}
			}
		}
	}

	fmt.Println("Checking ", len(allPoints), " points")

	return 0
}

func visibleCount(point [2]int, records []Record) int {
	count := 0
	for _, r := range records {
		d := mathy.ManhattanDistance(point[0], point[1], r.SensorX, r.SensorY)
		if d <= r.ManhattanDist() {
			count++
		}
	}
	return count
}

func parseInput(input string) (ans []Record) {
	regex := regexp.MustCompile(`.*?x=(-?\d+), y=(-?\d+).*?x=(-?\d+), y=(-?\d+)`)
	for _, line := range strings.Split(input, "\n") {
		tokens := regex.FindStringSubmatch(line)
		ans = append(ans, Record{
			SensorX: cast.ToInt(tokens[1]),
			SensorY: cast.ToInt(tokens[2]),
			BeaconX: cast.ToInt(tokens[3]),
			BeaconY: cast.ToInt(tokens[4]),
		})
	}
	return ans
}

func manhattanRing(x, y, dist int) [][2]int {

	ans := [][2]int{}

	for dx := 0; dx <= dist; dx++ {

		if dx == 0 {
			ans = append(ans, [2]int{x, y + (dist - dx)})
			ans = append(ans, [2]int{x, y - (dist - dx)})
		} else if dx == dist {
			ans = append(ans, [2]int{dx, y})
			ans = append(ans, [2]int{-dx, y})
		} else {
			ans = append(ans, [2]int{x + dx, y + (dist - dx)})
			ans = append(ans, [2]int{x - dx, y + (dist - dx)})
			ans = append(ans, [2]int{x - dx, y - (dist - dx)})
			ans = append(ans, [2]int{x + dx, y - (dist - dx)})

		}
	}

	return ans

}
