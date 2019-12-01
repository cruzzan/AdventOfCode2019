package day1

import (
	input_converter "AdventOfCode2019/src/input-converter"
	input_reader "AdventOfCode2019/src/input-reader"
	"fmt"
	"math"
)

func Task1() {
	lines := input_reader.ReadLines("resources/d1_t1.txt")
	modules := input_converter.StringToInt64(lines)
	fuel := SumFuelNet(modules)

	fmt.Println("Task 1: Fuel required is", fuel)
}

func Task2() {
	lines := input_reader.ReadLines("resources/d1_t1.txt")
	modules := input_converter.StringToInt64(lines)
	fuel := SumFuelGross(modules)

	fmt.Println("Task 2: Fuel required is", fuel)
}

func SumFuelNet(modules []int64) int {
	sum := 0

	for _, module := range modules {
		sum = sum + CalcFuelNeed(module)
	}

	return sum
}

func SumFuelGross(modules []int64) int {
	sum := 0

	for _, module := range modules {
		baseFule := CalcFuelNeed(module)
		totalFuel := baseFule

		delta := baseFule

		for delta > 0 {
			delta = CalcFuelNeed(int64(delta))

			if delta <= 0 {
				continue
			}

			totalFuel = totalFuel + delta
		}


		sum = sum + totalFuel
	}



	return sum
}

func CalcFuelNeed(m int64) int {
	return int(math.Floor(float64(m/3)))-2
}
