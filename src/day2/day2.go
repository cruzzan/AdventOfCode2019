package day2

import (
	"AdventOfCode2019/src/InputConverter"
	inputReader "AdventOfCode2019/src/InputReader"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadCsv("resources/d2_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	// Prep the input data
	instructions[1] = 12
	instructions[2] = 2

	output := RunIntCodeProgram(instructions)

	fmt.Println("Task 1: Back up state", output[0])
}

func Task2() {
	lines := inputReader.ReadCsv("resources/d2_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	in1, in2 := ReverseIntCodeProgram(instructions, 19690720, 0, 99)

	fmt.Println("Task 2: Reverseengineer inputs i1:", in1, "i2:", in2, ",as desired output:", 100*in1+in2)
}
