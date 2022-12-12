package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input_small.txt
var input string

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type uint64Slice []uint64

func (x uint64Slice) Len() int           { return len(x) }
func (x uint64Slice) Less(i, j int) bool { return x[i] < x[j] }
func (x uint64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Monkey struct {
	Index           uint64
	Items           []uint64
	Equation        string
	Test            uint64
	IfTrue          uint64
	IfFalse         uint64
	InspectionCount uint64
}

func (m *Monkey) EvaluateEquation(item uint64) uint64 {
	e := strings.ReplaceAll(m.Equation, "old", cast.ToString(item))
	tokens := strings.Split(e, " ")
	a := cast.ToUInt64(tokens[0])
	op := tokens[1]
	b := cast.ToUInt64(tokens[2])

	if op == "+" {
		return a + b
	} else if op == "*" {
		return a * b
	} else {
		panic("unkmown operator")
	}
}

func (m *Monkey) InspectItem(monkeys []*Monkey, worryLevel int, verbose bool) {
	//x, a = a[0], a[1:]
	item := m.Items[0]
	m.Items = m.Items[1:]
	m.InspectionCount++
	if verbose {
		fmt.Println("\tMonkey inspects an item with a worry level of ", item, ".")
	}
	x := m.EvaluateEquation(item)

	if verbose {
		fmt.Println("\tEquation ", m.Equation, " evaluates to ", x, ",")
	}
	x = x / uint64(worryLevel)

	if verbose {
		fmt.Println("\tMonkey gets bored with item. Worry level is divided by 3 to ", x, ".")
	}
	tossIndex := m.IfFalse
	if x%m.Test == 0 {
		if verbose {
			fmt.Println("\tCurrent worry level   IS   divisible by ", m.Test, ".")
			fmt.Println("\tItem with worry level ", x, " is thrown to monkey ", m.IfTrue)
		}
		tossIndex = m.IfTrue
	} else {
		if verbose {
			fmt.Println("\tCurrent worry level IS NOT divisible by ", m.Test, ".")
			fmt.Println("\tItem with worry level ", x, " is thrown to monkey ", m.IfFalse)
		}
	}

	if verbose {
		fmt.Println()
	}
	monkeys[tossIndex].Items = append(monkeys[tossIndex].Items, x)

	//Monkey 0:
	//Monkey inspects an item with a worry level of 79.
	//Worry level is multiplied by 19 to 1501.
	//Monkey gets bored with item. Worry level is divided by 3 to 500.
	//Current worry level is not divisible by 23.
	//Item with worry level 500 is thrown to monkey 3.
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

func (m *Monkey) HoldingString() string {
	s, _ := json.Marshal(m.Items)
	return string(s)
}

func part1(input string) uint64 {
	monkeys := parseInput(input)

	fmt.Println("Loaded ", len(monkeys), " Monkeys")
	rounds := 20

	for r := 0; r < rounds; r++ {
		for i, m := range monkeys {
			_ = i
			//fmt.Println("Monkey: ", i)
			for len(m.Items) > 0 {
				m.InspectItem(monkeys, 3, false)
			}
		}

		fmt.Println("After Round ", r+1, ":")
		for i, m := range monkeys {
			fmt.Println("Monkey[", i, "] is holding: ", m.HoldingString())
		}
		fmt.Println()

	}

	counts := []uint64{}

	for _, m := range monkeys {
		fmt.Println("Monkey[", m.Index, "] inspected ", m.InspectionCount, " items.")
		counts = append(counts, m.InspectionCount)
	}

	sort.Sort(sort.Reverse(uint64Slice(counts)))
	//sort.Ints(counts)
	return counts[0] * counts[1]

	return 0
}

func part2(input string) uint64 {
	monkeys := parseInput(input)

	fmt.Println("Loaded ", len(monkeys), " Monkeys")
	rounds := 10000

	for r := 0; r < rounds; r++ {
		for i, m := range monkeys {
			_ = i
			//fmt.Println("Monkey: ", i)
			for len(m.Items) > 0 {
				m.InspectItem(monkeys, 1, false)
			}
		}

		fmt.Println("After Round ", r+1, ":")
		for i, m := range monkeys {
			fmt.Println("Monkey[", i, "] is holding: ", m.HoldingString())
		}
		fmt.Println()

	}

	counts := []uint64{}

	for _, m := range monkeys {
		fmt.Println("Monkey[", m.Index, "] inspected ", m.InspectionCount, " items.")
		counts = append(counts, m.InspectionCount)
	}

	sort.Sort(sort.Reverse(uint64Slice(counts)))
	//sort.Ints(counts)
	return counts[0] * counts[1]

}

func parseInput(input string) (ans []*Monkey) {
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 7 {
		m := parseMonkey(lines[i : i+6])
		ans = append(ans, m)
	}

	return ans
}

func parseMonkey(lines []string) *Monkey {

	reLine1 := regexp.MustCompile(`^\s*Monkey (\d+):$`)
	reLine2 := regexp.MustCompile(`^\s*Starting items: (.*?)$`)
	reLine3 := regexp.MustCompile(`^\s*Operation: new = (.*?)$`)
	reLine4 := regexp.MustCompile(`^\s*Test: divisible by (.*?)$`)
	reLine5 := regexp.MustCompile(`^\s*If true: throw to monkey (\d|)$`)
	reLine6 := regexp.MustCompile(`^\s*If false: throw to monkey (\d|)$`)

	monkeyNum := cast.ToUInt64(reLine1.FindStringSubmatch(lines[0])[1])
	items := cast.StringListToUInt64List(strings.Split(reLine2.FindStringSubmatch(lines[1])[1], ", "))
	equation := reLine3.FindStringSubmatch(lines[2])[1]
	divisor := cast.ToUInt64(reLine4.FindStringSubmatch(lines[3])[1])
	ifTrue := cast.ToUInt64(reLine5.FindStringSubmatch(lines[4])[1])
	ifFalse := cast.ToUInt64(reLine6.FindStringSubmatch(lines[5])[1])

	m := Monkey{
		Index:           monkeyNum,
		Items:           items,
		Equation:        equation,
		Test:            divisor,
		IfTrue:          ifTrue,
		IfFalse:         ifFalse,
		InspectionCount: 0,
	}

	return &m
}
