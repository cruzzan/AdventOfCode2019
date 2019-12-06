package day4

import (
	"strconv"
)

func CountPossibleBetween(start int, end int) int {
	count := 0

	for i := start; i <= end; i++ {
		if passwordAcceptable(i) {
			count++
		}
	}

	return count
}

func CountPossibleBetween2(start int, end int) int {
	count := 0

	for i := start; i <= end; i++ {
		if passwordAcceptable2(i) {
			count++
		}
	}

	return count
}

func passwordAcceptable(p int) bool {
	sp := strconv.Itoa(p)
	adjacentFound := false

	if len(sp) != 6 {
		return false
	}

	for key := range sp {
		if key > 0 && sp[key] < sp[key-1] {
			return false
		}

		if key > 0 && sp[key] == sp[key-1] {
			adjacentFound = true
		}
	}

	if !adjacentFound {
		return false
	}

	return true
}

func passwordAcceptable2(p int) bool {
	sp := strconv.Itoa(p)

	if len(sp) != 6 {
		return false
	}

	adjacentCount := make(map[int]int)
	for key := range sp {
		if key > 0 && sp[key] < sp[key-1] {
			return false
		}

		if key > 0 && sp[key] == sp[key-1] {
			adjacentCount[int(sp[key])]++
		}
	}

	twoAdjacentFound := false
	for _, a := range adjacentCount {
		if a == 1 { // The number should only have been found in series once
			twoAdjacentFound = true
		}
	}

	if !twoAdjacentFound {
		return false
	}

	return true
}
