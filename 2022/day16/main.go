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
	"strings"
)

//go:embed input.txt
var input string

type Vertex struct {
	Name     string
	FlowRate int
}

type VertexAndPath struct {
	Vertex *Vertex
	Path   string
}

type VertexState struct {
	Steps      int
	FlowRate   int
	TotalFlow  int
	OpenValves []string
}

type VertexAndState struct {
	Vertex *Vertex
	State  VertexState
}

type StaticData struct {
	Graph     *graph.AdjacencyList[Vertex]
	Distances map[[2]string]int
}

func (vas VertexAndState) canOpen() bool {
	return vas.Vertex.FlowRate > 0 && !contains(vas.State.OpenValves, vas.Vertex.Name)
}

func (v VertexState) Print() {
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

	fmt.Println("Calculating Shortest Paths")
	for _, from := range simpleGraph.GetVertices() {
		fmt.Println("Calculating paths for vertex: ", from.Name)
		// calculate the shorted path from from to all vertices in the original graph
		distances, _ := fullGraph.DjikstraDistances(from)

		for to, d := range distances {
			if to.FlowRate > 0 && d > 0 {
				distanceMx[[2]string{from.Name, to.Name}] = d
				fmt.Println("Distance from ", from.Name, " to ", to.Name, " = ", d)
				// this is another flow node
				simpleGraph.AddEdge(from, to, d)
			}
		}
	}

	initialState := VertexState{
		Steps:      0,
		FlowRate:   0,
		TotalFlow:  0,
		OpenValves: []string{"AA"},
	}

	staticData := StaticData{
		Graph:     simpleGraph,
		Distances: distanceMx,
	}

	ans := solve(vertices["AA"], initialState, staticData)

	x := 0
	for i := range ans {
		x = mathy.Max(x, ans[i].TotalFlow)

	}

	fmt.Println(x)
	return x
}

func solve(current *Vertex, currentState VertexState, staticData StaticData) []VertexState {

	// vertex is the current vertex
	// currentState is the state of the problem that it took to get here
	// assume currentState is valid

	// check for terminal condition
	if currentState.Steps > 30 {
		return []VertexState{}
	} else if currentState.Steps == 30 {
		return []VertexState{currentState}
	} else {

		// open each neighbor
		// create a new state
		// if the state is valid, explore it

		retVal := []VertexState{}
		skipCount := 0
		for _, next := range staticData.Graph.GetNeighbors(current) {

			// check to see if next is already in the open list
			if !contains(currentState.OpenValves, next.Vertex.Name) {
				// get the distance from current to next
				d := staticData.Distances[[2]string{current.Name, next.Vertex.Name}]

				nextState := VertexState{
					Steps:      currentState.Steps + d + 1, // +1 for opening this vertex
					FlowRate:   currentState.FlowRate + next.Vertex.FlowRate,
					TotalFlow:  currentState.TotalFlow + (d+1)*currentState.FlowRate,
					OpenValves: append(currentState.OpenValves, next.Vertex.Name),
				}

				nextSol := solve(next.Vertex, nextState, staticData)
				retVal = append(retVal, nextSol...)
			} else {
				skipCount++
			}
		}

		if skipCount == len(staticData.Graph.GetNeighbors(current)) {

			remaining := 30 - currentState.Steps
			nextState := VertexState{
				Steps:      currentState.Steps + remaining,
				FlowRate:   currentState.FlowRate,
				TotalFlow:  currentState.TotalFlow + (remaining * currentState.FlowRate),
				OpenValves: currentState.OpenValves,
			}

			retVal = append(retVal, nextState)

		}
		return retVal
	}

}

func part2(input string) int {
	return 0
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
			ans.AddEdge(v, vertexNames[e], 1)
		}
	}

	return ans, vertexNames

}

func copyMap(m map[string]bool) map[string]bool {
	retVal := map[string]bool{}
	for k, v := range m {
		retVal[k] = v
	}
	return retVal
}
