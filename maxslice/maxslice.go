package maxslice

import "fmt"

// O(n2) /// GAH this actually doesn't quite work either since it can pull a single.
func FindMax(s []int) int {
	l, result := len(s), 0
	for i := 0; i < l; i++ {
		u := 0
		for n := i; n < l; n++ {
			//			fmt.Printf("%d %d %d\n", i, u, result)
			u += s[n]
			// has result be set yet?
			if i < 2 {
				result = u
			} else {
				result = Max(result, u)
			}

		}
	}

	return result
}

// I HATE nested loops.... but I wanted to leave the original attempt (since that was what I was working toward)... I
// think this is O(n)
// this doesn't work yet for all - integers
func FindMaxB(s []int) int {
	var start, mEnd, mS int = 0, 0, 0

	i := 0
	for _, n := range s {
		if i < 2 {
			start = mEnd + n
			fmt.Printf("%d %d %d\n", start, mEnd, n)
		}
		//		fmt.Printf("%d %d %d\n", n, mEnd, mS)
		mEnd = Max(start, mEnd+n)
		mS = Max(mS, mEnd)
		i++
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
