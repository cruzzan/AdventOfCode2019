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
	in := make(chan int)
	out := make(chan int)
	halt := make(chan intcode.HaltObject)

	go intcode.RunIntCodeProgram(instructions, in, out, halt)

	in<-1

	outputs := make([]int, 0)
	for o := range out{
		outputs = append(outputs, o)
	}

	fmt.Println("Task 1: TEST diagnostic", outputs[len(outputs)-1])
}

func Task2() {
	lines := inputReader.ReadCsv("resources/d5_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])
	in := make(chan int)
	out := make(chan int)
	halt := make(chan intcode.HaltObject)

	go intcode.RunIntCodeProgram(instructions, in, out, halt)

	in<-5

	outputs := make([]int, 0)
	for o := range out{
		outputs = append(outputs, o)
	}

	fmt.Println("Task 2: TEST diagnostic", outputs[len(outputs)-1])
}
