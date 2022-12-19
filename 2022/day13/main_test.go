package main

import (
	"fmt"
	"strings"
	"testing"
)

var example = ``

//
//func Test_part1(t *testing.T) {
//	tests := []struct {
//		name  string
//		input string
//		want  int
//	}{
//		{
//			name:  "example",
//			input: example,
//			want:  0,
//		},
//		// {
//		// 	name:  "actual",
//		// 	input: input,
//		// 	want:  0,
//		// },
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := part1(tt.input); got != tt.want {
//				t.Errorf("part1() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_part2(t *testing.T) {
//	tests := []struct {
//		name  string
//		input string
//		want  int
//	}{
//		{
//			name:  "example",
//			input: example,
//			want:  0,
//		},
//		// {
//		// 	name:  "actual",
//		// 	input: input,
//		// 	want:  0,
//		// },
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := part2(tt.input); got != tt.want {
//				t.Errorf("part2() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_parseNode1(t *testing.T) {
	runTest("[1,1,3,1,1]", "1|1|3|1|1")
}
func Test_parseNode2(t *testing.T) {
	runTest("[1,1,5,1,1]", "1|1|5|1|1")
}
func Test_parseNode3(t *testing.T) {
	runTest("[[1],[2,3,4]]", "[1]|[2,3,4]")
}
func Test_parseNode4(t *testing.T) {
	runTest("[[1],4]", "[1]|4")
}
func Test_parseNode5(t *testing.T) {
	runTest("[[8,7,6]]", "[8,7,6]")
}
func Test_parseNode6(t *testing.T) {
	runTest("[1,[2,[3,[4,[5,6,7]]]],8,9]", "1|[2,[3,[4,[5,6,7]]]]|8|9")
}
func Test_parseNode7(t *testing.T) {
	runTest("[1,[2,[3,[4,[5,6,0]]]],8,9]", "1|[2,[3,[4,[5,6,0]]]]|8|9")
}

func Test_compare1(t *testing.T) {
	runCompare("[1,1,3,1,1]", "[1,1,5,1,1]", CORRECT)
}

func Test_compare2(t *testing.T) {
	runCompare("[[1],[2,3,4]]", "[[1],4]", CORRECT)
}

func Test_compare3(t *testing.T) {
	runCompare("[9]", "[[8,7,6]]", INCORRECT)
}

func Test_compare4(t *testing.T) {
	runCompare("[[4,4],4,4]", "[[4,4],4,4,4]", CORRECT)
}

func Test_compare5(t *testing.T) {
	runCompare("[7,7,7,7]", "[7,7,7]", INCORRECT)
}

func Test_compare6(t *testing.T) {
	runCompare("[]", "[3]", CORRECT)
}

func Test_compare7(t *testing.T) {
	runCompare("[[[]]]", "[[]]", INCORRECT)
}

func Test_compare8(t *testing.T) {
	runCompare("[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]", INCORRECT)
}

func runCompare(left string, right string, expected CompareType) {
	x := compare(left, right)
	if x != expected {
		panic(fmt.Sprintf("Expected %v got %v", x, expected))
	}
}

//

func runTest(input, expected string) {
	s := strings.Join(parseNode(input), "|")
	if s != expected {
		panic("Expected " + expected + " got " + s)
	}
}

//
//	s = fmt.Sprint(parseNode("[1,1,5,1,1]"))
//	if s != "[1 1 5 1 1]" {
//		panic("Expected [1 1 5 1 1], got " + s)
//	}
//
//	s = fmt.Sprint(parseNode("[[1],[2,3,4]]"))
//	if s != "[1 1 5 1 1]" {
//		panic("Expected [1 1 5 1 1], got " + s)
//	}
//
//}
