package day2

import (
	"github.com/sirupsen/logrus"
)

func RunIntCodeProgram(instructions []int) []int {
	for cur := 0; cur <= len(instructions); cur += 4 {
		switch instructions[cur] {
		case 1:
			in1Cur := instructions[cur+1]
			in2Cur := instructions[cur+2]
			outCur := instructions[cur+3]
			instructions[outCur] = instructions[in1Cur] + instructions[in2Cur]
			break
		case 2:
			in1Cur := instructions[cur+1]
			in2Cur := instructions[cur+2]
			outCur := instructions[cur+3]
			instructions[outCur] = instructions[in1Cur] * instructions[in2Cur]
			break
		case 99:
			return instructions
			break
		default:
			logrus.Fatalf("Shit, something went wrong we got OpCode %d at position %d, %#v", instructions[cur], cur, instructions)
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
			copy(i ,instructions)

			i[1] = in1
			i[2] = in2

			res := RunIntCodeProgram(i)

			if output == res[0] {
				return in1, in2
			}
		}
	}

	return min - 1, min -1
}
