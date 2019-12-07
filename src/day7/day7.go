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

	p := buildPhases(0,4)
	highestOutput := 0

	for _, phaseSet := range p {
		prevOutput := 0
		for i := 0; i < 5; i++ {
			in := make(chan int)
			out := make(chan int)
			halt := make(chan intcode.HaltObject)

			go intcode.RunIntCodeProgram(instructions, in, out, halt)

			in<-phaseSet[i]
			in<-prevOutput

			for o := range out {
				prevOutput = o
			}
		}

		if prevOutput > highestOutput {
			highestOutput = prevOutput
		}
	}

	fmt.Printf("Task 1: max thrust %d\n", highestOutput)
}

func Task2() {
	lines := inputReader.ReadCsv("resources/d7_t1.txt", ',')
	instructions := InputConverter.StringToInt(lines[0])

	p := buildPhases(5,9)
	highestOutput := 0

	for _, phaseSet := range p {
		chan1 := make(chan int, 1)
		chan2 := make(chan int, 1)
		chan3 := make(chan int, 1)
		chan4 := make(chan int, 1)
		chan5 := make(chan int, 1)
		out := make(chan int, 1)

		// Init amps
		go intcode.RunIntCodeProgram(instructions, chan1, chan2, make(chan intcode.HaltObject))
		go intcode.RunIntCodeProgram(instructions, chan2, chan3, make(chan intcode.HaltObject))
		go intcode.RunIntCodeProgram(instructions, chan3, chan4, make(chan intcode.HaltObject))
		go intcode.RunIntCodeProgram(instructions, chan4, chan5, make(chan intcode.HaltObject))
		go intcode.RunIntCodeProgram(instructions, chan5, out, make(chan intcode.HaltObject))

		// Prep phase settings
		chan1<-phaseSet[0]
		chan2<-phaseSet[1]
		chan3<-phaseSet[2]
		chan4<-phaseSet[3]
		chan5<-phaseSet[4]

		// Start program
		chan1<-0

		outputs := make([]int, 0)
		for o := range out {
			outputs = append(outputs, o)
			chan1<-o
		}

		if outputs[len(outputs)-1] > highestOutput {
			highestOutput = outputs[len(outputs)-1]
		}
	}

	fmt.Printf("Task 2: max thrust with feedback loop %d\n", highestOutput)
}

func buildPhases(min int, max int) [][]int {
	p := make([][]int, 0)

	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			if j != i {
				for k := min; k <= max; k++ {
					if k != j && k != i {
						for l := min; l <= max; l++ {
							if l != i && l != j && l != k {
								for m := min; m <= max; m++ {
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
