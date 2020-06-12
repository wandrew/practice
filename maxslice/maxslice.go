package maxslice

import "fmt"

type Solution int

// O(n2)
func FindMax(s []int) int {
	var result int
	l := len(s)
	for i := 1; i < l-1; i++ {
		var u int /// I want the number to be negative
		fmt.Printf("Iteration %d\n", i)

		for n := i; n < l; n++ {
			//			fmt.Printf("%d %d %d\n", i, u, result)
			// Check if u is initialized
			if n == 0 {
				fmt.Printf("Var u is empty for index %d, skipping sum and assigning to %d...\n", n, s[n])
				u = s[n]
				continue
			} else {
				fmt.Printf("Summing %d + %d for index %d...\n", u, s[n], n)
				u += s[n]
				fmt.Printf("Sum is: %d\n", u)
			}

			fmt.Printf("Current Sum of U for index %d is: %d\n", n, u)
			// has result be set yet?
			fmt.Printf("Evaluating greater of last result: %d and current computed result %d\n", result, u)
			result = Max(result, u)
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
