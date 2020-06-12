package maxslice

import (
	"fmt"
	"math"
)

type Solution int

// O(n2)
func FindMax(s []int) int {
	var result int
	l := len(s)
	for i := 0; i < l; i++ {
		var u int /// I want the number to be negative
		fmt.Printf("Iteration %d\n", i)

		addends := 0
		var addendSlice []int
		for n := i; n < l; n++ {
			addends += 1
			addendSlice = append(addendSlice, s[n])
			//			fmt.Printf("%d %d %d\n", i, u, result)
			// Check if u is initialized
			if n == i {
				fmt.Printf("Var u is empty for index %+d, skipping sum and assigning to %+d...\n", n, s[n])
				u = s[n]
				continue
			} else {
				fmt.Printf("Summing %+d + %+d (total of %v) for index %+d...\n", u, s[n], addendSlice, n)
				u += s[n]
				fmt.Printf("Sum is: %+d\n", u)
			}

			fmt.Printf("Current Sum of U for index %+d is: %+d\n", n, u)
			// has result be set yet?

			// we want to set the result only after we have the sum of 2 indexes and never when only a single index is included
			// set our first result if we have 2 or more addends
			if addends > 1 {

				if i == 0 {
					fmt.Printf("Seeding result with initial value of %+d\n", u)
					result = u
				} else {
					fmt.Printf("Evaluating greater of result: %+d and current computed sum %+d\n", result, u)
					result = Max(result, u)
				}

			}

		}
	}

	return result
}

func FindMaxOptimized(s []int) int {

	l := len(s)
	//if l < 3 {
	//	return -1, fmt.Errorf("slice must have 3 or more nodes")
	//}

	i := 0
	for {
		if i >= l-1 {
			break
		} // current window
		sz := l - i
		fmt.Printf("Current window size is (inclusive) %v\n", sz)

		// move the window

		//if 0+i != l-i {
		//	sz := (l - i) - (0 + i)
		//	fmt.Printf("Current window size is (inclusive) %v\n", sz)
		//	win := []int{0 + i, l - i}
		//	fmt.Printf("Current window is (inclusive) %v\n", win)
		//}
		i += 1

	}

	//	return -1, fmt.Errorf("this function isn't functional")
	return -1
}

// I HATE nested loops.... but I wanted to leave the original attempt (since that was what I was working toward)... I
// think this is O(n)
func FindMaxB(s []int) int {
	var mEnd, mS float64 = math.Inf(-1), math.Inf(-1)
	//9:  {-1, -4, 4, 4, -9, 10},             // this was just wrong... it WAS 9 [4,4,-9,10]

	// first window

	for i, n := range s {
		fmt.Printf("\n\nIteration %d\n", i)
		fmt.Printf("evaluating (mS) %v\n", mS)

		fmt.Printf("evaluating (mEnd) %v\n", mEnd)
		fmt.Printf("evaluating (n) %v\n", n)

		fmt.Printf("evaluating (mEnd+n) %v\n", mEnd+float64(n))
		mEnd = MaxFloat(float64(n), mEnd+float64(n))
		fmt.Printf("calculated (mEnd) %v\n", mEnd)
		if i > 0 {

			mS = MaxFloat(mS, mEnd)
		}

	}

	return int(mS)
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func MaxFloat(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}
