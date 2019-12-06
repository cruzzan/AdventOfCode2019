package day1

import (
	"math"
)

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
	return int(math.Floor(float64(m/3))) - 2
}
