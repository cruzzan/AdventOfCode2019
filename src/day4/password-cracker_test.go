package day4

import "testing"

func TestPasswordAcceptable(t *testing.T) {
	cases := []struct{
		in int
		want bool
	}{
		{in: 122345, want: true},
		{in: 111111, want: true},
		{in: 223450, want: false},
		{in: 123789, want: false},
	}

	for _, c := range cases {
		got := passwordAcceptable(c.in)

		if got != c.want {
			t.Errorf("Failed asserting that the acceptability of %d was %t, expected %t", c.in, got, c.want)
		}
	}
}

func TestPasswordAcceptable2(t *testing.T) {
	cases := []struct{
		in int
		want bool
	}{
		//{in: 112233, want: true},
		{in: 123444, want: false},
		//{in: 111122, want: true},
	}

	for _, c := range cases {
		got := passwordAcceptable2(c.in)

		if got != c.want {
			t.Errorf("Failed asserting that the acceptability of %d was %t, expected %t", c.in, got, c.want)
		}
	}
}
