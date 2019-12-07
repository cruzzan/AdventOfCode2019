package intcode

import (
	"fmt"
)

type HaltObject struct {
	Instructions []int
}

func RunIntCodeProgram(instructions []int, inChan chan int, outChan chan int, haltChan chan HaltObject) {
	cur := 0
	for cur <= len(instructions) {
		opCode := getOpCode(instructions[cur])
		p1Mode := getParam1Mode(instructions[cur])
		p2Mode := getParam2Mode(instructions[cur])
		//p3Mode := getParam3Mode(instructions[cur])

		switch opCode {
		case 1: // Addition
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			instructions[instructions[cur+3]] = in1 + in2

			cur += 4
			break
		case 2: // Multiplication
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			instructions[instructions[cur+3]] = in1 * in2

			cur += 4
			break
		case 3: // Input
			i := <-inChan

			instructions[instructions[cur+1]] = i

			cur += 2
			break
		case 4: // Output
			out := getValue(instructions, p1Mode, cur+1)

			outChan<-out

			cur += 2
			break
		case 5: // Jump if non-zero
			val := getValue(instructions, p1Mode, cur+1)
			jmp := getValue(instructions, p2Mode, cur+2)

			if val != 0 {
				cur = jmp
				break
			}

			cur += 3
			break
		case 6: // Jump if zero
			val := getValue(instructions, p1Mode, cur+1)
			jmp := getValue(instructions, p2Mode, cur+2)

			if val == 0 {
				cur = jmp
				break
			}

			cur += 3
			break
		case 7: // Less than
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			output := 0

			if in1 < in2 {
				output = 1
			}

			instructions[instructions[cur+3]] = output
			cur += 4
			break
		case 8: // Equal
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			output := 0

			if in1 == in2 {
				output = 1
			}

			instructions[instructions[cur+3]] = output
			cur += 4
			break
		case 99: // Halt
			close(outChan)
			haltChan<-HaltObject{
				Instructions:instructions,
			}
			return
		default:
			fmt.Printf("Shit, something went wrong we got OpCode %d at position %d, %#v\n", opCode, cur, instructions)
		}
	}
}

func ReverseIntCodeProgram(instructions []int, output int, min int, max int) (int, int) {
	in1 := min
	for ; in1 <= max && in1 < len(instructions); in1++ {
		in2 := min
		for ; in2 <= max && in2 < len(instructions); in2++ {
			i := make([]int, len(instructions))
			copy(i, instructions)

			i[1] = in1
			i[2] = in2

			in := make(chan int)
			out := make(chan int)
			halt := make(chan HaltObject)
			go RunIntCodeProgram(i, in, out, halt)

			hObj := <-halt


			if output == hObj.Instructions[0] {
				return in1, in2
			}
		}
	}

	return min - 1, min - 1
}

func getOpCode(i int) int {
	return i % 100
}

func getParam1Mode(i int) int {
	if ((i % 1000) / 100) >= 1 {
		return 1
	}

	return 0
}

func getParam2Mode(i int) int {
	if ((i % 10000) / 1000) >= 1 {
		return 1
	}

	return 0
}

func getParam3Mode(i int) int {
	if ((i % 100000) / 10000) >= 1 {
		return 1
	}

	return 0
}

func getValue(i []int, m int, cur int) int {
	if m == 1 {
		return i[cur]
	}

	return i[i[cur]]
}
