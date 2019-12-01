package day1

import "testing"

func TestCalcModuleFuelNeed(t *testing.T) {
	tests := []struct{
		in int64
		want int
	}{
		{in: 12, want: 2},
		{in: 14, want: 2},
		{in: 1969, want: 654},
		{in: 100756, want: 33583},
	}

	for _, c := range tests {
		got := CalcFuelNeed(c.in)

		if got != c.want {
			t.Errorf("Faild asserting that %d is equal to expected %d", got, c.want)
		}
	}
}

func TestSumFuelNet(t *testing.T) {
	in := []int64{
		12,
		14,
		1969,
		100756,
	}

	want := 34241

	got := SumFuelNet(in)

	if got != want {
		t.Errorf("Faild asserting that %d is equal to expected %d", got, want)
	}
}

func TestSumFuelGross(t *testing.T) {
	in := []int64{
		100756,
	}

	want := 50346

	got := SumFuelGross(in)

	if got != want {
		t.Errorf("Faild asserting that %d is equal to expected %d", got, want)
	}
}
