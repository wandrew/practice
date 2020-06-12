package maxslice

import "fmt"

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

// I HATE nested loops.... but I wanted to leave the original attempt (since that was what I was working toward)... I
// think this is O(n)
func FindMaxB(s []int) int {
	var start, mEnd, mS int = 0, 0, 0

	for _, n := range s {
		//		fmt.Printf("%d %d %d\n", n, mEnd, mS)
		mEnd = Max(start, mEnd+n)
		mS = Max(mS, mEnd)
	}

	return mS
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
