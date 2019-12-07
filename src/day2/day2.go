package day2

import (
	"AdventOfCode2019/src/InputConverter"
	inputReader "AdventOfCode2019/src/InputReader"
	"AdventOfCode2019/src/intcode"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadCsv("resources/d2_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])
	in := make(chan int)
	out := make(chan int)
	halt := make(chan intcode.HaltObject)

	// Prep the input data
	instructions[1] = 12
	instructions[2] = 2

	go intcode.RunIntCodeProgram(instructions, in, out, halt)

	hobj := <-halt

	fmt.Println("Task 1: Back up state", hobj.Instructions[0])
}

func Task2() {
	lines := inputReader.ReadCsv("resources/d2_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	in1, in2 := intcode.ReverseIntCodeProgram(instructions, 19690720, 0, 99)

	fmt.Println("Task 2: Reverseengineer inputs i1:", in1, "i2:", in2, ",as desired output:", 100*in1+in2)
}
