package main

import (
	"AdventOfCode2019/src/day1"
	"AdventOfCode2019/src/day2"
	"AdventOfCode2019/src/day4"
	"AdventOfCode2019/src/day5"
	"AdventOfCode2019/src/day6"
	"AdventOfCode2019/src/day7"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	day = kingpin.Flag("day", "The day to run (1-25)").Short('d').Int()
)

func main() {
	kingpin.Parse()

	logrus.Info("Welcome to AoC 2019, by cruzzan")
	logrus.Info("Day ", *day)

	switch *day {
	case 1:
		day1.Task1()
		day1.Task2()
		break
	case 2:
		day2.Task1()
		day2.Task2()
		break
	case 3:
		logrus.Info("Yeah.. this is not finished yet...")
	case 4:
		day4.Task1()
		day4.Task2()
		break
	case 5:
		day5.Task1()
		day5.Task2()
		break
	case 6:
		day6.Task1()
		day6.Task2()
		break
	case 7:
		day7.Task1()
		break
	default:
		logrus.Info("Yeah.. i haven't even started this one yet...")
	}

	os.Exit(0)
}
