package maxslice

import "testing"

func TestFindMax(t *testing.T) {
	tests := map[int][]int{
		//9:  {-1, -4, 4, 4, -9, 10},             // this was just wrong... it WAS 9
		//14: {3, -4, 4, 4, -9, 7, 7, -22},       //.
		-4: {-8, -9, -3, -4, -9, -13, -7, -22}, // fails because -3
	}

	for expected, test := range tests {
		if observed := FindMax(test); observed != expected {
			t.Fatalf("FindMax() = %v, want %v", observed, expected)
		}
		//if observed := FindMaxB(test); observed != expected {
		//	t.Fatalf("FindMaxB() = %v, want %v", observed, expected)
		//}
	}

}
