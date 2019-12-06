package day5

import (
	"AdventOfCode2019/src/InputConverter"
	inputReader "AdventOfCode2019/src/InputReader"
	"AdventOfCode2019/src/intcode"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadCsv("resources/d5_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	fmt.Println("Task 1: TEST diagnostic")
	intcode.RunIntCodeProgram(instructions)

}
