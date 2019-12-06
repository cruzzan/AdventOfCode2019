package day6

import (
	inputReader "AdventOfCode2019/src/InputReader"
	"fmt"
)

func Task1() {
	orbitMap := inputReader.ReadCsv("resources/d6_t1.txt", ')')
	c := Count(orbitMap)

	fmt.Println("Task 1: Orbit count", c)
}

func Task2() {
	orbitMap := inputReader.ReadCsv("resources/d6_t1.txt", ')')
	c := TransfersBetween(orbitMap, "YOU", "SAN")

	fmt.Println("Task 2: Orbital transfer count", c)
}
