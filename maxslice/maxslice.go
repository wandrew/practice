package maxslice

// O(n2)
func FindMax(s []int) int {
	l, result := len(s), 0
	for i := 0; i < l; i++ {
		u := 0
		for n := i; n < l; n++ {
			//			fmt.Printf("%d %d %d\n", i, u, result)
			u += s[n]
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
