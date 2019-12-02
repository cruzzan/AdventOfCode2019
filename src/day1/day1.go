package day1

import (
	"AdventOfCode2019/src/InputConverter"
	"AdventOfCode2019/src/InputReader"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadLines("resources/d1_t1.txt")
	modules := InputConverter.StringToInt64(lines)
	fuel := SumFuelNet(modules)

	fmt.Println("Task 1: Fuel required is", fuel)
}

func Task2() {
	lines := inputReader.ReadLines("resources/d1_t1.txt")
	modules := InputConverter.StringToInt64(lines)
	fuel := SumFuelGross(modules)

	fmt.Println("Task 2: Fuel required is", fuel)
}
