package inputReader

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	f := "../../resources/test.txt"
	want := []string{
		"Hey",
		"This",
		"Is",
		"A",
		"Test",
	}

	got := ReadLines(f)

	if len(got) != len(want) {
		t.Fatalf("Faild to assert that lists were equally long, expected %d, got %d", len(want), len(got))
	}

	for key, val := range want {
		if got[key] != val {
			t.Errorf("Failed to assert that %s is equal to expected %s", got[key], val)
		}
	}
}
