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

	_, out := intcode.RunIntCodeProgram(instructions, []int{1})

	fmt.Println("Task 1: TEST diagnostic", out[len(out)-1])
}

func Task2() {
	lines := inputReader.ReadCsv("resources/d5_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	_, out := intcode.RunIntCodeProgram(instructions, []int{5})

	fmt.Println("Task 2: TEST diagnostic", out[len(out)-1])
}
