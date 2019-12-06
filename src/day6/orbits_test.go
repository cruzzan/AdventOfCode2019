package day6

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		in   [][]string
		want int
	}{
		{
			in:   [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}},
			want: 6,
		},
		{
			in:   [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"}, {"E", "J"}, {"J", "K"}, {"K", "L"}},
			want: 42,
		},
	}

	for _, c := range cases {
		got := Count(c.in)

		if got != c.want {
			t.Errorf("Failed asserting that %d is equal to expected %d", got, c.want)
		}
	}
}

func TestTransfersBetween(t *testing.T) {
	cases := []struct {
		in     [][]string
		inOrg  string
		inDest string
		want   int
	}{
		{
			in:     [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"}, {"E", "J"}, {"J", "K"}, {"K", "L"}},
			inOrg:  "F",
			inDest: "L",
			want:   2,
		},
		{
			in:     [][]string{{"COM", "B"}, {"B", "C"}, {"C", "D"}, {"D", "E"}, {"E", "F"}, {"B", "G"}, {"G", "H"}, {"D", "I"}, {"E", "J"}, {"J", "K"}, {"K", "L"}, {"K", "YOU"}, {"I", "SAN"}},
			inOrg:  "YOU",
			inDest: "SAN",
			want:   4,
		},
	}

	for _, c := range cases {
		got := TransfersBetween(c.in, c.inOrg, c.inDest)

		if got != c.want {
			t.Errorf("Failed asserting that %d is equal to expected %d", got, c.want)
		}
	}
}
