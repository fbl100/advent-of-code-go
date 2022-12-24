package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/data-structures/graph"
	"github.com/alexchao26/advent-of-code-go/mathy"
	"github.com/alexchao26/advent-of-code-go/util"
	"regexp"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type Vertex struct {
	Name     string
	FlowRate int
}

type ProblemState struct {
	Steps      int
	FlowRate   int
	TotalFlow  int
	OpenValves []string
}

func (p ProblemState) Step(count int) ProblemState {
	return ProblemState{
		Steps:      p.Steps + count,
		FlowRate:   p.FlowRate,
		TotalFlow:  p.TotalFlow + count*p.FlowRate,
		OpenValves: p.OpenValves,
	}
}

func (p ProblemState) StepTo(step int) ProblemState {
	return p.Step(step - p.Steps)
}

func (p ProblemState) OpenValve(valve string, flowRate int) ProblemState {

	// don't open valves with flowRate == 0
	if flowRate == 0 {
		return p
	}
	retVal := ProblemState{
		Steps:      p.Steps + 1,
		FlowRate:   p.FlowRate + flowRate,
		TotalFlow:  p.TotalFlow + p.FlowRate,
		OpenValves: append(p.OpenValves, valve),
	}

	//fmt.Println(strings.Join(retVal.OpenValves, ", "))
	return retVal
}

type StaticData struct {
	Graph     *graph.AdjacencyList[Vertex]
	Distances map[[2]string]int
	MaxTime   int
}

type OpenValveSet struct {
	OpenValves []string
	MaxFlow    int
}

type OpenValveSetArray []OpenValveSet

func (o OpenValveSetArray) Len() int {
	return len(o)
}

func (o OpenValveSetArray) Less(i, j int) bool {
	return o[i].MaxFlow > o[j].MaxFlow
}

func (o OpenValveSetArray) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (v ProblemState) Print() {
	fmt.Println("== Minute ", v.Steps, " ==")
	if len(v.OpenValves) == 0 {
		fmt.Println("No valves are open")
	} else if len(v.OpenValves) == 1 {
		fmt.Println("Valve ", v.OpenValves[0], " is open, releasing ", v.FlowRate, " pressure")
	} else {
		fmt.Println("Valves ", strings.Join(v.OpenValves, " and "), " are open, releasing ", v.FlowRate, " pressure. Total Release is ", v.TotalFlow)
	}
}

func contains(arr []string, item string) bool {
	for _, x := range arr {
		if x == item {
			return true
		}
	}
	return false
}

func string2key(arr []string) string {
	sort.Strings(arr)
	return strings.Join(arr, "")
}

func key2array(s string) []string {
	retVal := []string{}
	for i := 0; i < len(s); i += 2 {
		retVal = append(retVal, s[i:i+2])
	}
	return retVal
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
	fullGraph, vertices := parseInput(input)
	_ = vertices

	// create a graph that contains 'AA", and all vertices
	// with positive flow rates.
	simpleGraph := graph.NewAdjacencyList[Vertex]()

	vertexes := []*Vertex{}

	// add AA and the vertices that have positive flows
	for _, v := range fullGraph.GetVertices() {
		if v.FlowRate > 0 || v.Name == "AA" {
			vertexes = append(vertexes, v)
			simpleGraph.AddVertex(v) // add to the simple graph
		}
	}

	distanceMx := map[[2]string]int{}

	//fmt.Println("Calculating Shortest Paths")
	for _, from := range simpleGraph.GetVertices() {
		//fmt.Println("Calculating paths for vertex: ", from.Name)
		// calculate the shorted path from from to all vertices in the original graph
		distances, _ := fullGraph.DjikstraDistances(from)

		for to, d := range distances {
			if to.FlowRate > 0 && d > 0 {
				distanceMx[[2]string{from.Name, to.Name}] = d
				//fmt.Println("Distance from ", from.Name, " to ", to.Name, " = ", d)
				// this is another flow node
				simpleGraph.AddEdge(from, to, d)
			}
		}
	}

	initialState := ProblemState{
		Steps:      0,
		FlowRate:   0,
		TotalFlow:  0,
		OpenValves: []string{"AA"},
	}

	staticData := StaticData{
		Graph:     simpleGraph,
		Distances: distanceMx,
		MaxTime:   30,
	}

	ans := solve(vertices["AA"], initialState, staticData)

	x := 0
	for i := range ans {
		x = mathy.Max(x, ans[i].TotalFlow)

	}

	return x
}

func solve(current *Vertex, currentState ProblemState, staticData StaticData) []ProblemState {

	// vertex is the current vertex
	// currentState is the state of the problem that it took to get here
	// assume currentState is valid

	// option 1: we do nothing, and we coast until the end
	// option 2: if the flow rate > 0, we turn on the valve and coast until the end
	// option 3: we turn on the valve, but then we explore submodes

	retVal := []ProblemState{}

	if currentState.Steps >= staticData.MaxTime {
		panic("too many steps")
	}
	// option 1: propagate current solution to end
	retVal = append(retVal, currentState.StepTo(staticData.MaxTime))

	// option 2: if we can, turn on the valve
	if current.FlowRate > 0 {
		openedValveState := currentState.OpenValve(current.Name, current.FlowRate)
		retVal = append(retVal, openedValveState.StepTo(staticData.MaxTime))
	}

	// now, explore all submodes
	for _, next := range staticData.Graph.GetNeighbors(current) {

		// check to see if openedValveState is already in the open list
		if !contains(currentState.OpenValves, next.Vertex.Name) {
			// get the distance from current to openedValveState
			d := staticData.Distances[[2]string{current.Name, next.Vertex.Name}]

			nextState := currentState.OpenValve(current.Name, current.FlowRate).Step(d)
			if nextState.Steps < staticData.MaxTime {
				nextSol := solve(next.Vertex, nextState, staticData)
				retVal = append(retVal, nextSol...)
			}
		}
	}

	return retVal

}

func part2(input string) int {
	fullGraph, vertices := parseInput(input)
	_ = vertices

	// create a graph that contains 'AA", and all vertices
	// with positive flow rates.
	simpleGraph := graph.NewAdjacencyList[Vertex]()

	vertexes := []*Vertex{}

	// add AA and the vertices that have positive flows
	for _, v := range fullGraph.GetVertices() {
		if v.FlowRate > 0 || v.Name == "AA" {
			vertexes = append(vertexes, v)
			simpleGraph.AddVertex(v) // add to the simple graph
		}
	}

	distanceMx := map[[2]string]int{}

	//fmt.Println("Calculating Shortest Paths")
	for _, from := range simpleGraph.GetVertices() {
		//fmt.Println("Calculating paths for vertex: ", from.Name)
		// calculate the shorted path from from to all vertices in the original graph
		distances, _ := fullGraph.DjikstraDistances(from)

		for to, d := range distances {
			if to.FlowRate > 0 && d > 0 {
				distanceMx[[2]string{from.Name, to.Name}] = d
				//fmt.Println("Distance from ", from.Name, " to ", to.Name, " = ", d)
				// this is another flow node
				simpleGraph.AddEdge(from, to, d)
			}
		}
	}

	//dot := simpleGraph.ToDot(
	//	func(v *Vertex) string {
	//		return v.Name
	//	},
	//	func(v *Vertex) string {
	//		return fmt.Sprint(`label="`, v.Name, " (", v.FlowRate, `)"`)
	//	},
	//	func(e graph.WeightedEdge[Vertex]) string {
	//		return fmt.Sprint(`label="`, e.Weight, `"`)
	//	})
	//
	//fmt.Println(dot)

	initialState := ProblemState{
		Steps:      0,
		FlowRate:   0,
		TotalFlow:  0,
		OpenValves: []string{},
	}

	staticData := StaticData{
		Graph:     simpleGraph,
		Distances: distanceMx,
		MaxTime:   26,
	}

	ans := solve(vertices["AA"], initialState, staticData)

	maxSets := map[string]int{}

	x := 0
	for i := range ans {
		if ans[i].TotalFlow > 0 {
			x = mathy.Max(x, ans[i].TotalFlow)
			key := string2key(ans[i].OpenValves)
			_, ok := maxSets[key]
			if ok {
				// already in
				maxSets[key] = mathy.Max(maxSets[key], ans[i].TotalFlow)
			} else {
				maxSets[key] = ans[i].TotalFlow
			}
		}
	}

	//for k, v := range maxSets {
	//	fmt.Println(k, " - ", v)
	//}

	openSets := OpenValveSetArray{}
	for k, v := range maxSets {
		openSets = append(openSets, OpenValveSet{
			OpenValves: key2array(k),
			MaxFlow:    v,
		})
	}

	sort.Sort(openSets)

	//for _, os := range openSets {
	//	fmt.Println(os.MaxFlow, " - ", strings.Join(os.OpenValves, ", "))
	//}

	best := 0
	for i := 0; i < len(openSets); i++ {
		for j := 0; j < len(openSets); j++ {
			if disjoint(openSets[i].OpenValves, openSets[j].OpenValves) {
				total := openSets[i].MaxFlow + openSets[j].MaxFlow
				best = mathy.MaxInt(best, total)
			}
		}
	}

	return best

}

func parseInput(input string) (*graph.AdjacencyList[Vertex], map[string]*Vertex) {
	re := regexp.MustCompile(`Valve (.*?) has flow rate=(\d+); tunnels? leads? to valves? (.*?)$`)

	ans := graph.NewAdjacencyList[Vertex]()
	edges := map[*Vertex][]string{}
	vertexNames := map[string]*Vertex{}

	for _, line := range strings.Split(input, "\n") {
		m := re.FindStringSubmatch(line)
		v := Vertex{
			Name:     m[1],
			FlowRate: cast.ToInt(m[2]),
		}
		ans.AddVertex(&v)
		edges[&v] = strings.Split(m[3], ", ")
		vertexNames[v.Name] = &v
	}

	for v := range edges {
		for _, e := range edges[v] {

			_, ok := vertexNames[e]
			if !ok {
				// vertex does not exist
				v := Vertex{
					Name:     e,
					FlowRate: 0,
				}
				ans.AddVertex(&v)
				vertexNames[e] = &v
			}

			ans.AddEdge(v, vertexNames[e], 1)
		}
	}

	return ans, vertexNames

}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func disjoint(a, b []string) bool {
	for i := range a {
		if contains(b, a[i]) {
			return false
		}
	}

	for i := range b {
		if contains(a, b[i]) {
			return false
		}
	}

	return true

}
