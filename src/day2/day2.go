package day2

import (
	"AdventOfCode2019/src/InputConverter"
	inputReader "AdventOfCode2019/src/InputReader"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadCsv("resources/d2_t1.txt")
	instructions := InputConverter.StringToInt(lines)

	// Prep the input data
	instructions[1] = 12
	instructions[2] = 2

	output := RunIntCodeProgram(instructions)

	fmt.Println("Task 1: Back up state", output[0])
}
