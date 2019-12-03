package day2

import "github.com/sirupsen/logrus"

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
			logrus.Fatalf("Shit, something went wrong we got OpCode %d at position %d", instructions[cur], cur)
		}
	}

	return []int{}
}
