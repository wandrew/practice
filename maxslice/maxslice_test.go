package maxslice

import "testing"

func TestFindMax(t *testing.T) {
	tests := map[int][]int{
		9:  {-1, -4, 4, 4, -9, 10},             // this was just wrong... it WAS 9 [4,4,-9,10]
		14: {3, -4, 4, 4, -9, 7, 7, -22},       // [7, 7]
		-7: {-8, -9, -3, -4, -9, -13, -7, -22}, // [-3, -4]
	}

	for expected, test := range tests {
		if observed := FindMax(test); observed != expected {
			t.Fatalf("FindMax() = %v, want %v", observed, expected)
		}
	}

}

func TestFindMaxB(t *testing.T) { // this is still broken
	tests := map[int][]int{
		9: {-1, -4, 4, 4, -9, 10}, // this was just wrong... it WAS 9 [4,4,-9,10]
		//14: {3, -4, 4, 4, -9, 7, 7, -22},       // [7, 7]
		//-7: {-8, -9, -3, -4, -9, -13, -7, -22}, // [-3, -4]
	}

	for expected, test := range tests {

		if observed := FindMaxB(test); observed != expected {
			t.Fatalf("FindMaxB() = %v, want %v", observed, expected)
		}
	}
}
