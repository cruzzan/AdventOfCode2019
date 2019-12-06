package intcode

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func RunIntCodeProgram(instructions []int) []int {
	cur := 0
	for cur <= len(instructions) {
		opCode := getOpCode(instructions[cur])
		p1Mode := getParam1Mode(instructions[cur])
		p2Mode := getParam2Mode(instructions[cur])
		//p3Mode := getParam3Mode(instructions[cur])

		switch opCode {
		case 1:
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			instructions[instructions[cur+3]] = in1 + in2

			cur += 4
			break
		case 2:
			in1 := getValue(instructions, p1Mode, cur+1)
			in2 := getValue(instructions, p2Mode, cur+2)
			instructions[instructions[cur+3]] = in1 * in2

			cur += 4
			break
		case 3:
			var i int
			fmt.Println("Make an input:")
			fmt.Scan(&i)

			instructions[instructions[cur+1]] = i

			cur += 2
			break
		case 4:
			out := getValue(instructions, p1Mode, cur+1)
			fmt.Println(out)

			cur += 2
			break
		case 99:
			return instructions
			break
		default:
			logrus.Fatalf("Shit, something went wrong we got OpCode %d at position %d, %#v", opCode, cur, instructions)
		}
	}

	return []int{}
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

			res := RunIntCodeProgram(i)

			if output == res[0] {
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
