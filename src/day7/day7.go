package day7

import (
	"AdventOfCode2019/src/InputConverter"
	inputReader "AdventOfCode2019/src/InputReader"
	"AdventOfCode2019/src/intcode"
	"fmt"
)

func Task1() {
	lines := inputReader.ReadCsv("resources/d7_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	p := buildPhases(4)
	highestOutput := 0

	for _, phaseSet := range p {
		prevOutput := 0
		for i := 0; i < 5; i++ {
			inputs := []int{phaseSet[i], prevOutput}
			_, outputs := intcode.RunIntCodeProgram(instructions, inputs)

			prevOutput = outputs[len(outputs)-1]
		}

		if prevOutput > highestOutput {
			highestOutput = prevOutput
		}
	}

	fmt.Printf("Task 1: max thrust %d\n", highestOutput)
}

func buildPhases(max int) [][]int {
	p := make([][]int, 0)

	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			if j != i {
				for k := 0; k <= max; k++ {
					if k != j && k != i {
						for l := 0; l <= max; l++ {
							if l != i && l != j && l != k {
								for m := 0; m <= max; m++ {
									if m != i && m != j && m != k && m != l {
										p = append(p, []int{i, j, k, l, m})
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return p
}
