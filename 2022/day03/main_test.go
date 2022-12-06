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

func Test_commonLetters(t *testing.T) {

	x := commonLetters("abcdAusSSd", "adefS")
	for a := range x {
		println(a)
	}

}

func Test_digits(t *testing.T) {
	x := digits("abc")

	println(x)
}

func Test_priority(t *testing.T) {
	x := priority('a')
	if x != 1 {
		panic("expected 1")
	}

	x = priority('z')
	if x != 26 {
		panic("expected 26")
	}

	x = priority('A')
	if x != 27 {
		panic("expected 27")
	}

	x = priority('Z')
	if x != 52 {
		panic("expected 52")
	}

}
