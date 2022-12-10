package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

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
	parsed := parseInput(input)

	nodes := parsed.Flatten()

	ans := 0
	for _, n := range nodes {
		s := n.TotalSize()
		if s <= 100000 {
			ans += s
		}
	}
	return ans
}

func part2(input string) int {
	parsed := parseInput(input)

	capacity := 70000000
	required := 30000000

	free := capacity - parsed.TotalSize()
	needed := required - free
	println("Need ", needed)

	nodes := parsed.Flatten()

	bestExcess := capacity
	var bestNode *Node = nil

	for _, n := range nodes {
		s := n.TotalSize()
		spaceLeftIfDeleted := free + s
		excess := spaceLeftIfDeleted - required

		if excess > 0 && excess < bestExcess {
			bestExcess = excess
			bestNode = n
			println(bestExcess)
		}

		// needed + s - required
	}
	return bestNode.TotalSize()
}

type File struct {
	Name string
	Size int
}

type Node struct {
	Parent        *Node
	DirectoryName string
	Files         []*File
	SubDirs       map[string]*Node
}

func (n *Node) Flatten() []*Node {
	retVal := []*Node{n}

	for _, d := range n.SubDirs {
		retVal = append(retVal, d.Flatten()...)
	}

	return retVal
}

func (n *Node) TotalSize() int {
	// sum of files + sum of subdirs
	size := 0
	for _, f := range n.Files {
		size += f.Size
	}

	for _, d := range n.SubDirs {
		size += d.TotalSize()
	}
	return size
}

func (n *Node) PrintRecursive(level int) {
	buff := strings.Repeat("  ", level)
	println(buff, "- ", n.DirectoryName, " (dir)")
	for _, f := range n.Files {
		println(buff, "- ", f.Name, " (file, size =", f.Size, ")")
	}
	for _, f := range n.SubDirs {
		f.PrintRecursive(level + 1)
	}
}

func makeNode(name string, parent *Node) *Node {
	retVal := Node{
		DirectoryName: name,
		Parent:        parent,
		Files:         []*File{},
		SubDirs:       make(map[string]*Node),
	}
	return &retVal
}

func parseInput(input string) *Node {
	root := makeNode("Top", nil)

	current := root
	_ = current
	//
	//regex := *regexp.MustCompile(`move (\d+) from (\d) to (\d)`)
	//for i := firstBlank + 1; i < len(lines); i++ {
	//	res := regex.FindStringSubmatch(lines[i])

	cd := regexp.MustCompile(`\$ cd (.+?)$`)
	ls := regexp.MustCompile(`\$ ls`)
	file := regexp.MustCompile(`(\d+) (.*)$`)
	dir := regexp.MustCompile(`dir (.*)$`)

	lines := strings.Split(input, "\n")
	i := 0
	for i < len(lines) {
		if cd.MatchString(lines[i]) {
			dir := cd.FindStringSubmatch(lines[i])[1]
			if dir == ".." {
				current = current.Parent
			} else {
				if current.SubDirs[dir] != nil {
					current = current.SubDirs[dir]
				} else {
					// new diretory
					next := makeNode(dir, current)
					current.SubDirs[dir] = next
					current = next
				}
			}
			i++
		} else if ls.MatchString(lines[i]) {
			i++
			for i < len(lines) && (file.MatchString(lines[i]) || dir.MatchString(lines[i])) {
				if file.MatchString(lines[i]) {
					x := file.FindStringSubmatch(lines[i])
					size, _ := strconv.Atoi(x[1])
					name := x[2]

					current.Files = append(current.Files, &File{Name: name, Size: size})
				} else if dir.MatchString(lines[i]) {
					x := dir.FindStringSubmatch(lines[i])
					name := x[1]
					current.SubDirs[name] = makeNode(name, current)
				} else {
					panic("?")
				}
				i++
			}

		}

	}

	return root
}
