package main

import "fmt"

func returnCodes(number int) (int, int, int, int) {
	return number % 100, (number % 1000) / 100, (number % 10000) / 1000, (number % 100000) / 10000
}

func getValue(puzzleInput []int, parameterCode, nextOrTwoPastValue, relativeValue int) int {
	// fmt.Println("in getValues", parameterCode, nextOrTwoPastValue, relativeValue)
	if parameterCode == 0 {
		// 0 is position mode
		return puzzleInput[nextOrTwoPastValue]
	}
	if parameterCode == 1 {
		// 1 is immediate mode (no-else-return)
		return nextOrTwoPastValue
	}
	// 2 is relative mode...
	sum := nextOrTwoPastValue + relativeValue
	if sum >= len(puzzleInput) {
		puzzleInput = resize(puzzleInput, sum)
	}
	return puzzleInput[sum]
}

func runDiagnostics(puzzleInput []int, inputValue int) string {
	// fmt.Println(puzzleInput)

	var relativeValue int // nil value of zero

	for i := 0; i < len(puzzleInput); {
		// find op code (last 2 digits of number), a 1, 2, 3, 4, or 99
		// find param1 and param2 which are the 100's and 1000's digit in the i-th element
		opCode, param1, param2, _ := returnCodes(puzzleInput[i])
		var firstVal, secondVal, thirdVal int

		// set value variables based on what the opCodes are (some of them will throw errors if I attempt to set them...)
		if opCode == 1 || opCode == 2 || opCode == 7 || opCode == 8 {
			firstVal, secondVal = getValue(puzzleInput, param1, puzzleInput[i+1], relativeValue), getValue(puzzleInput, param2, puzzleInput[i+2], relativeValue)
			thirdVal = puzzleInput[i+3]
		} else if opCode == 3 || opCode == 4 {
			firstVal = puzzleInput[i+1]
		} else if opCode == 5 || opCode == 6 {
			firstVal, secondVal = getValue(puzzleInput, param1, puzzleInput[i+1], relativeValue), getValue(puzzleInput, param2, puzzleInput[i+2], relativeValue)
		} else if opCode == 9 {
			// not actually sure what I need to do here...
			firstVal = getValue(puzzleInput, param1, puzzleInput[i+1], relativeValue)
		}

		// fmt.Println("desired specs", opCode, firstVal, thirdVal)

		// switch statement to handle the opcode value
		switch opCode {
		case 99:
			// fmt.Println("99 halted at index: ", i)
			return "halted"
		case 1:
			// add and place (by position)
			// firstToAdd, secondToAdd := getValue(puzzleInput, param1, puzzleInput[i+1]), getValue(puzzleInput, param2, puzzleInput[i+2])
			if thirdVal >= len(puzzleInput) {
				puzzleInput = resize(puzzleInput, thirdVal)
			}
			puzzleInput[thirdVal] = firstVal + secondVal
			i += 4
		case 2:
			// multiply and place (by position)
			if thirdVal >= len(puzzleInput) {
				puzzleInput = resize(puzzleInput, thirdVal)
			}
			puzzleInput[thirdVal] = firstVal * secondVal
			i += 4
		case 3:
			// write inputValue to puzzle input by position
			if firstVal >= len(puzzleInput) {
				puzzleInput = resize(puzzleInput, firstVal)
			}
			puzzleInput[firstVal] = inputValue
			i += 2
		case 4:
			// output the value to the console, always by position
			fmt.Println("output: ", puzzleInput[firstVal])
			i += 2
		case 5:
			// jump-if-true
			// if first param is != 0 set instruction pointer to value from second parameter
			if firstVal != 0 {
				// fmt.Println("5 moves i to", secondVal)
				i = secondVal
			} else {
				// otherwise do nothing except increment i
				// increment by 3 (like opCode 6)
				i += 3
			}
		case 6:
			// jump-if-false
			// if first param is != 0 set instruction pointer to value from second parameter
			if firstVal == 0 {
				i = secondVal
			} else {
				// otherwise increment i to the next instruction
				// this instruction has 3 values, so i increments by 3
				i += 3
			}
		case 7:
			// less than - if firstVal < secondVal, store a 1 @ position of thirdVal, otherwise store a 0
			if firstVal < secondVal {
				puzzleInput[thirdVal] = 1
			} else {
				puzzleInput[thirdVal] = 0
			}
			i += 4
		case 8:
			// equals, same as opCode 7 but conditional is firstVal is equal to secondVal
			if firstVal == secondVal {
				puzzleInput[thirdVal] = 1
			} else {
				puzzleInput[thirdVal] = 0
			}
			i += 4
		case 9:
			// relativeValue += puzzleInput[i+1]
			relativeValue += firstVal
			// fmt.Print("relative value op code")
			i += 2
		default:
			fmt.Println("bad opCode!!!")
			return "badCode"
		}
	}
	return "EOF"
}

func main() {
	// puzzle input as a slice
	puzzleInput := []int{1102, 34463338, 34463338, 63, 1007, 63, 34463338, 63, 1005, 63, 53, 1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 209, 6, 209, 3, 203, 0, 1008, 1000, 1, 63, 1005, 63, 65, 1008, 1000, 2, 63, 1005, 63, 904, 1008, 1000, 0, 63, 1005, 63, 58, 4, 25, 104, 0, 99, 4, 0, 104, 0, 99, 4, 17, 104, 0, 99, 0, 0, 1101, 27, 0, 1014, 1101, 286, 0, 1023, 1102, 1, 35, 1018, 1102, 20, 1, 1000, 1101, 26, 0, 1010, 1101, 0, 289, 1022, 1102, 1, 30, 1019, 1102, 734, 1, 1025, 1102, 1, 31, 1012, 1101, 25, 0, 1001, 1102, 1, 1, 1021, 1101, 0, 36, 1002, 1101, 0, 527, 1028, 1101, 895, 0, 1026, 1102, 1, 23, 1016, 1101, 21, 0, 1003, 1102, 22, 1, 1011, 1102, 1, 522, 1029, 1102, 1, 892, 1027, 1102, 1, 0, 1020, 1102, 1, 28, 1015, 1102, 38, 1, 1006, 1101, 0, 32, 1008, 1101, 743, 0, 1024, 1101, 0, 37, 1007, 1102, 1, 24, 1013, 1102, 1, 33, 1009, 1102, 39, 1, 1004, 1102, 1, 34, 1005, 1102, 1, 29, 1017, 109, 19, 21102, 40, 1, -3, 1008, 1016, 40, 63, 1005, 63, 203, 4, 187, 1106, 0, 207, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -7, 2101, 0, -7, 63, 1008, 63, 32, 63, 1005, 63, 227, 1106, 0, 233, 4, 213, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -3, 2108, 37, -2, 63, 1005, 63, 255, 4, 239, 1001, 64, 1, 64, 1105, 1, 255, 1002, 64, 2, 64, 109, 11, 21108, 41, 40, -6, 1005, 1014, 275, 1001, 64, 1, 64, 1106, 0, 277, 4, 261, 1002, 64, 2, 64, 109, 10, 2105, 1, -7, 1105, 1, 295, 4, 283, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -27, 1201, -2, 0, 63, 1008, 63, 25, 63, 1005, 63, 321, 4, 301, 1001, 64, 1, 64, 1105, 1, 321, 1002, 64, 2, 64, 109, 15, 21107, 42, 41, 0, 1005, 1018, 341, 1001, 64, 1, 64, 1106, 0, 343, 4, 327, 1002, 64, 2, 64, 109, -25, 2108, 20, 10, 63, 1005, 63, 359, 1105, 1, 365, 4, 349, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 12, 2107, 35, 0, 63, 1005, 63, 385, 1001, 64, 1, 64, 1106, 0, 387, 4, 371, 1002, 64, 2, 64, 109, 4, 21101, 43, 0, 6, 1008, 1015, 43, 63, 1005, 63, 409, 4, 393, 1106, 0, 413, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 9, 21101, 44, 0, -8, 1008, 1010, 46, 63, 1005, 63, 437, 1001, 64, 1, 64, 1106, 0, 439, 4, 419, 1002, 64, 2, 64, 109, 5, 21108, 45, 45, -4, 1005, 1019, 457, 4, 445, 1106, 0, 461, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -22, 2102, 1, 7, 63, 1008, 63, 33, 63, 1005, 63, 481, 1106, 0, 487, 4, 467, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 14, 21102, 46, 1, -1, 1008, 1014, 43, 63, 1005, 63, 507, 1106, 0, 513, 4, 493, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 12, 2106, 0, 1, 4, 519, 1106, 0, 531, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -17, 1205, 10, 547, 1001, 64, 1, 64, 1106, 0, 549, 4, 537, 1002, 64, 2, 64, 109, -8, 1202, -2, 1, 63, 1008, 63, 17, 63, 1005, 63, 569, 1105, 1, 575, 4, 555, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 23, 1206, -5, 593, 4, 581, 1001, 64, 1, 64, 1105, 1, 593, 1002, 64, 2, 64, 109, -14, 1208, -8, 24, 63, 1005, 63, 613, 1001, 64, 1, 64, 1105, 1, 615, 4, 599, 1002, 64, 2, 64, 109, -2, 1207, -1, 33, 63, 1005, 63, 633, 4, 621, 1105, 1, 637, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 2, 21107, 47, 48, 5, 1005, 1016, 659, 4, 643, 1001, 64, 1, 64, 1105, 1, 659, 1002, 64, 2, 64, 109, -11, 1208, 8, 32, 63, 1005, 63, 681, 4, 665, 1001, 64, 1, 64, 1106, 0, 681, 1002, 64, 2, 64, 109, 2, 2101, 0, 0, 63, 1008, 63, 36, 63, 1005, 63, 703, 4, 687, 1106, 0, 707, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 12, 1206, 7, 719, 1106, 0, 725, 4, 713, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 2, 2105, 1, 8, 4, 731, 1001, 64, 1, 64, 1106, 0, 743, 1002, 64, 2, 64, 109, -21, 2102, 1, 9, 63, 1008, 63, 39, 63, 1005, 63, 769, 4, 749, 1001, 64, 1, 64, 1105, 1, 769, 1002, 64, 2, 64, 109, 11, 1201, -3, 0, 63, 1008, 63, 24, 63, 1005, 63, 793, 1001, 64, 1, 64, 1105, 1, 795, 4, 775, 1002, 64, 2, 64, 109, 20, 1205, -5, 809, 4, 801, 1105, 1, 813, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, -23, 1207, 4, 36, 63, 1005, 63, 833, 1001, 64, 1, 64, 1105, 1, 835, 4, 819, 1002, 64, 2, 64, 109, -3, 2107, 33, 5, 63, 1005, 63, 853, 4, 841, 1106, 0, 857, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 16, 1202, -9, 1, 63, 1008, 63, 37, 63, 1005, 63, 879, 4, 863, 1105, 1, 883, 1001, 64, 1, 64, 1002, 64, 2, 64, 109, 12, 2106, 0, -1, 1105, 1, 901, 4, 889, 1001, 64, 1, 64, 4, 64, 99, 21101, 0, 27, 1, 21101, 0, 915, 0, 1106, 0, 922, 21201, 1, 48476, 1, 204, 1, 99, 109, 3, 1207, -2, 3, 63, 1005, 63, 964, 21201, -2, -1, 1, 21101, 0, 942, 0, 1105, 1, 922, 21202, 1, 1, -1, 21201, -2, -3, 1, 21101, 0, 957, 0, 1105, 1, 922, 22201, 1, -1, -2, 1106, 0, 968, 21202, -2, 1, -2, 109, -3, 2106, 0, 0}

	// tests
	// some 16 digit number
	// puzzleInput := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}

	// the middle number
	// puzzleInput := []int{104, 1125899906842624, 99}

	// Part 1: ID of System is 1 (use as the input)
	runDiagnostics(puzzleInput, 1)

	// Part 2: ID of System is 5 (use as the input)
	// runDiagnostics(puzzleInput, 5)
}

// reize function
func resize(puzzle []int, newSize int) []int {
	result := make([]int, newSize+1)
	copy(result, puzzle)
	// fmt.Println(len(puzzle), len(result))
	return result
}
