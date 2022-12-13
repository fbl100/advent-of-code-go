package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"regexp"
	"sort"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input_small.txt
var input string

type Monkey struct {
	Index           int
	Items           [][]int // store values as coefficients of the base
	Equation        string
	Base            int
	IfTrue          int
	IfFalse         int
	InspectionCount int
}

func (m *Monkey) EvaluateTerm(token string, old []int) []int {
	a := []int{}
	if token == "old" {
		a = old
	} else {
		a = mathy.Int2Poly(cast.ToInt(token), m.Base)
	}
	return a
}

func (m *Monkey) EvaluateEquation(item []int) []int {

	tokens := strings.Split(m.Equation, " ")

	a := m.EvaluateTerm(tokens[0], item)
	b := m.EvaluateTerm(tokens[2], item)
	op := tokens[1]

	if op == "+" {
		return mathy.PolyReduce(mathy.PolyAdd(a, b), m.Base)
	} else if op == "*" {
		return mathy.PolyReduce(mathy.PolyMul(a, b), m.Base)
	} else {
		panic("unkmown operator")
	}
}

func (m *Monkey) adjustWorryLevel(a []int, worryLevel int) []int {
	if worryLevel == 1 {
		return a
	} else {
		A := mathy.Poly2Int(a, m.Base)
		return mathy.Int2Poly(A/worryLevel, m.Base)
	}
}

func (m *Monkey) InspectItem(monkeys []*Monkey, worryLevel int, verbose bool) {
	//x, a = a[0], a[1:]
	item := m.Items[0]
	m.Items = m.Items[1:]
	m.InspectionCount++
	if verbose {
		fmt.Println("\tMonkey inspects an item with a worry level of ", mathy.Poly2Int(item, m.Base), ".")
	}
	x := m.EvaluateEquation(item)

	if verbose {
		fmt.Println("\tEquation ", m.Equation, " evaluates to ", mathy.Poly2Int(x, m.Base), ",")
	}

	x = m.adjustWorryLevel(x, worryLevel)

	if verbose {
		fmt.Println("\tMonkey gets bored with item. Worry level is divided by 3 to ", mathy.Poly2Int(x, m.Base), ".")
	}
	tossIndex := m.IfFalse

	if len(x) == 0 || x[0] == 0 {
		if verbose {
			fmt.Println("\tCurrent worry level   IS   divisible by ", m.Base, ".")
			fmt.Println("\tItem with worry level ", mathy.Poly2Int(x, m.Base), " is thrown to monkey ", m.IfTrue)
		}
		tossIndex = m.IfTrue
	} else {
		if verbose {
			fmt.Println("\tCurrent worry level IS NOT divisible by ", m.Base, ".")
			fmt.Println("\tItem with worry level ", mathy.Poly2Int(x, m.Base), " is thrown to monkey ", m.IfFalse)
		}
	}

	if verbose {
		fmt.Println()
	}

	x_new := mathy.ChangeBase(x, m.Base, monkeys[tossIndex].Base)

	monkeys[tossIndex].Items = append(monkeys[tossIndex].Items, x_new)

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

func (m *Monkey) HoldingInts() string {
	x := []int{}
	for _, c := range m.Items {
		x = append(x, mathy.Poly2Int(c, m.Base))
	}
	s, _ := json.Marshal(x)
	return string(s)
}

func part1(input string) int {
	monkeys := parseInput(input)

	fmt.Println("Loaded ", len(monkeys), " Monkeys")
	rounds := 20

	verbose := false

	for r := 0; r < rounds; r++ {
		for i, m := range monkeys {
			_ = i
			if verbose {
				fmt.Println("Monkey: ", i)
			}
			for len(m.Items) > 0 {
				m.InspectItem(monkeys, 3, verbose)
			}
		}

		fmt.Println("After Round ", r+1, ":")
		for i, m := range monkeys {
			fmt.Println("Monkey[", i, "] is holding: ", m.HoldingInts())
		}
		fmt.Println()

	}

	counts := []int{}

	for _, m := range monkeys {
		fmt.Println("Monkey[", m.Index, "] inspected ", m.InspectionCount, " items.")
		counts = append(counts, m.InspectionCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	//sort.Ints(counts)
	return counts[0] * counts[1]

	return 0
}

func part2(input string) int {
	monkeys := parseInput(input)

	fmt.Println("Loaded ", len(monkeys), " Monkeys")
	rounds := 10000

	verbose := false

	for r := 0; r < rounds; r++ {
		for i, m := range monkeys {
			_ = i

			if verbose {
				fmt.Println("Monkey: ", i)
			}
			for len(m.Items) > 0 {
				m.InspectItem(monkeys, 1, verbose)
			}
		}

		if verbose {
			fmt.Println("After Round ", r+1, ":")
			for i, m := range monkeys {
				fmt.Println("Monkey[", i, "] is holding: ", m.HoldingString())
			}
			fmt.Println()
		}

		if (r+1)%1000 == 0 || r+1 == 1 || r+1 == 20 {
			counts := []int{}

			fmt.Println("After Round ", r+1, ":")
			for _, m := range monkeys {
				fmt.Println("Monkey[", m.Index, "] inspected ", m.InspectionCount, " items.")
				counts = append(counts, m.InspectionCount)
			}
		}

	}

	counts := []int{}

	for _, m := range monkeys {
		fmt.Println("Monkey[", m.Index, "] inspected ", m.InspectionCount, " items.")
		counts = append(counts, m.InspectionCount)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
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

	monkeyNum := cast.ToInt(reLine1.FindStringSubmatch(lines[0])[1])
	items := cast.StringListToIntList(strings.Split(reLine2.FindStringSubmatch(lines[1])[1], ", "))
	equation := reLine3.FindStringSubmatch(lines[2])[1]
	divisor := cast.ToInt(reLine4.FindStringSubmatch(lines[3])[1])
	ifTrue := cast.ToInt(reLine5.FindStringSubmatch(lines[4])[1])
	ifFalse := cast.ToInt(reLine6.FindStringSubmatch(lines[5])[1])

	itemPolys := [][]int{}
	for _, i := range items {
		p := mathy.Int2Poly(i, divisor)
		itemPolys = append(itemPolys, p)
	}

	m := Monkey{
		Index:           monkeyNum,
		Items:           itemPolys,
		Equation:        equation,
		Base:            divisor,
		IfTrue:          ifTrue,
		IfFalse:         ifFalse,
		InspectionCount: 0,
	}

	return &m
}
