package maxslice

import "testing"

func TestFindMax(t *testing.T) {
	tests := map[int][]int{
		10: {-1, -4, 4, 4, -9, 10},
		14: {3, -4, 4, 4, -9, 7, 7, -22},
		-4: {-8, -9, -3, -4, -9, -7, -7, -22},
	}

	for expected, test := range tests {
		if observed := FindMax(test); observed != expected {
			t.Fatalf("FindMax() = %v, want %v", observed, expected)
		}
	}

}
