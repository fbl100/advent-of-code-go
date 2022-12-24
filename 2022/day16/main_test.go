package main

import (
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  0,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  0,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblemState_Step_(t *testing.T) {
	ps := ProblemState{
		Steps:      1,
		FlowRate:   3,
		TotalFlow:  0,
		OpenValves: []string{"AA"},
	}

	ps = ps.Step(1)
	if ps.Steps != 2 {
		panic("step mismatch")
	}
	if ps.TotalFlow != 3 {
		panic("flow mismatcb")
	}
}

func TestProblemState_OpenValve(t *testing.T) {
	ps := ProblemState{
		Steps:      0,
		FlowRate:   0,
		TotalFlow:  0,
		OpenValves: []string{},
	}

	ps = ps.OpenValve("CC", 3)
	if ps.Steps != 1 {
		panic("step mismatch")
	}
	if ps.TotalFlow != 0 {
		panic("flow mismatcb")
	}
	if ps.FlowRate != 3 {
		panic("flow mismatcb")
	}
	ps = ps.Step(5)
	if ps.Steps != 6 {
		panic("step mismatch")
	}
	if ps.TotalFlow != 15 {
		panic("flow mismatcb")
	}
	if ps.FlowRate != 3 {
		panic("flow mismatcb")
	}

}
