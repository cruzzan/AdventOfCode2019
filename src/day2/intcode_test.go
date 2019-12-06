package day2

import "testing"

func TestRunIntCodeProgram(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{in: []int{1, 0, 0, 0, 99}, want: []int{2, 0, 0, 0, 99}},
		{in: []int{2, 3, 0, 3, 99}, want: []int{2, 3, 0, 6, 99}},
		{in: []int{2, 4, 4, 5, 99, 0}, want: []int{2, 4, 4, 5, 99, 9801}},
		{in: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, want: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}

	for _, c := range cases {
		got := RunIntCodeProgram(c.in)

		if len(got) != len(c.want) {
			t.Fatalf("Failed asserting that the slices were the same length, got %d, expected %d", len(got), len(c.want))
		}

		for key, val := range got {
			if val != c.want[key] {
				t.Errorf("Failed asserting that the slices were equal, got %v, expected %v", got, c.want)
			}
		}
	}
}

func TestReverseIntCodeProgram(t *testing.T) {
	cases := []struct {
		inInsctructions []int
		inOutput        int
		inMin           int
		inMax           int
		wantIn1         int
		wantIn2         int
	}{
		{inInsctructions: []int{1, 0, 0, 0, 99}, inOutput: 100, inMin: 0, inMax: 100, wantIn1: 0, wantIn2: 4},
		{inInsctructions: []int{2, 3, 0, 0, 99}, inOutput: 0, inMin: 0, inMax: 100, wantIn1: 0, wantIn2: 1},
	}

	for key, c := range cases {
		println(key)
		got1, got2 := ReverseIntCodeProgram(c.inInsctructions, c.inOutput, c.inMin, c.inMax)

		if got1 != c.wantIn1 {
			t.Fatalf("Failed asserting that input 1 was correct, got %d, expected %d", got1, c.wantIn1)
		}

		if got2 != c.wantIn2 {
			t.Fatalf("Failed asserting that input 2 was correct, got %d, expected %d", got2, c.wantIn2)
		}
	}
}
