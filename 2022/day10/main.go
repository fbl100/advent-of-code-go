package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/cast"
	"github.com/alexchao26/advent-of-code-go/util"
)

//go:embed input.txt
var input string

type Instruction struct {
	// the name of the op code
	OpCode string

	// how many cycles this takes
	Cycles int

	// the operand (if provided)
	Operand int
	// the function to run on each cycle
	// takes the instruction pointer and the current value
	Func func(*Instruction, int) int
}

func (i *Instruction) String() string {

	operand := ""
	if i.Operand != math.MaxInt {
		operand = fmt.Sprint(i.Operand)
	}
	return fmt.Sprint("[", i.OpCode, "] - ", operand, " (", i.Cycles, ") remain")

}

func makeInstruction(tokens []string) Instruction {
	switch tokens[0] {
	case "noop":
		return Instruction{
			OpCode:  "noop",
			Cycles:  1,
			Operand: math.MaxInt,
			Func: func(instruction *Instruction, current int) int {
				// do nothing
				return current
			},
		}
	case "addx":
		return Instruction{
			OpCode:  "addx",
			Cycles:  2,
			Operand: cast.ToInt(tokens[1]),
			Func: func(instruction *Instruction, current int) int {
				// do nothing
				return current + instruction.Operand
			},
		}
	}
	panic("unable to parse instruction")
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

// return the new value of the register and the number of cycles it took
func executeOp(current int, i *Instruction) (int, int) {
	switch i.OpCode {
	case "noop":
		return current, 1
	case "addx":
		return current + i.Operand, 2
	}
	panic("unknown op code")

}

func part1(input string) int {
	instructions := parseInput(input)

	ip := 0

	cycle := 1

	value := 1

	// values at the start of the cycle
	start_history := []int{0}
	// values at the end of the cycle
	end_history := []int{0}

	for ip < len(instructions) {

		start_history = append(start_history, value)

		//fmt.Println("Start of cycle ", cycle, ", value = ", value, " instruction =", instructions[ip].String())

		if instructions[ip].Cycles > 1 {
			// decrement the instruction count
			instructions[ip].Cycles--
		} else if instructions[ip].Cycles == 1 {
			// execute the instruction
			value = instructions[ip].Func(&instructions[ip], value)
			ip++
		}
		//if ip < len(instructions) {
		//	fmt.Println("End of cycle ", cycle, ", value = ", value, " instruction =", instructions[ip].String())
		//} else {
		//	fmt.Println("Final Value: ", value)
		//}

		cycle++
		// log the value at this cycle
		end_history = append(end_history, value)
	}

	ans := 0
	for i := 20; i < len(start_history); i += 40 {
		log.Println("[", i, "] = ", start_history[i])
		ans += i * start_history[i]
	}

	return ans
}

func part2(input string) int {
	instructions := parseInput(input)

	ip := 0

	cycle := 1

	value := 1

	// values at the start of the cycle
	start_history := []int{0}
	// values at the end of the cycle
	end_history := []int{0}

	for ip < len(instructions) {

		start_history = append(start_history, value)

		//fmt.Println("Start of cycle ", cycle, ", value = ", value, " instruction =", instructions[ip].String())

		if instructions[ip].Cycles > 1 {
			// decrement the instruction count
			instructions[ip].Cycles--
		} else if instructions[ip].Cycles == 1 {
			// execute the instruction
			value = instructions[ip].Func(&instructions[ip], value)
			ip++
		}
		//if ip < len(instructions) {
		//	fmt.Println("End of cycle ", cycle, ", value = ", value, " instruction =", instructions[ip].String())
		//} else {
		//	fmt.Println("Final Value: ", value)
		//}

		cycle++
		// log the value at this cycle
		end_history = append(end_history, value)
	}

	for row := 0; row < 6; row++ {
		for col := 0; col < 40; col++ {

			cycleIndex := 40*row + col + 1
			pos := start_history[cycleIndex]
			if pos >= col-1 && pos <= col+1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	return 5
	//
	//ans := 0
	//for i := 20; i < len(start_history); i += 40 {
	//	log.Println("[", i, "] = ", start_history[i])
	//	ans += i * start_history[i]
	//}
	//
	//return ans
}

func parseInput(input string) (ans []Instruction) {
	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " ")
		ans = append(ans, makeInstruction(tokens))
	}
	return ans
}
