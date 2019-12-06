package main

import (
	"AdventOfCode2019/src/day1"
	"AdventOfCode2019/src/day2"
	"AdventOfCode2019/src/day4"
	"AdventOfCode2019/src/day6"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to AoC 2019, by cruzzan")

	day1.Task1()
	day1.Task2()

	day2.Task1()
	day2.Task2()

	day4.Task1()
	day4.Task2()

	day6.Task1()
	day6.Task2()

	os.Exit(0)
}
