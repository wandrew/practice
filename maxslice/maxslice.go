package maxslice

import (
	"fmt"
)

type Solution int

// O(n2)
func FindMax(s []int) int {
	var result int
	l := len(s)
	for i := 0; i < l; i++ {
		var u int /// I want the number to be negative
		//fmt.Printf("Iteration %d\n", i)

		addends := 0
		var addendSlice []int
		for n := i; n < l; n++ {
			addends += 1
			addendSlice = append(addendSlice, s[n])
			//			fmt.Printf("%d %d %d\n", i, u, result)
			// Check if u is initialized
			if n == i {
				//fmt.Printf("Var u is empty for index %+d, skipping sum and assigning to %+d...\n", n, s[n])
				u = s[n]
				continue
			} else {
				//fmt.Printf("Summing %+d + %+d (total of %v) for index %+d...\n", u, s[n], addendSlice, n)
				u += s[n]
				//fmt.Printf("Sum is: %+d\n", u)
			}

			//fmt.Printf("Current Sum of U for index %+d is: %+d\n", n, u)
			// has result be set yet?

			// we want to set the result only after we have the sum of 2 indexes and never when only a single index is included
			// set our first result if we have 2 or more addends
			if addends > 1 {

				if i == 0 {
					//fmt.Printf("Seeding result with initial value of %+d\n", u)
					result = u
				} else {
					//fmt.Printf("Evaluating greater of result: %+d and current computed sum %+d\n", result, u)
					result = Max(result, u)
				}

			}

		}
	}

	return result
}

func FindMaxOptimized(s []int) (int, error) { // O(n log)

	l := len(s)
	if l < 3 {
		return -1, fmt.Errorf("slice must have 3 or more nodes")
	}

	i := 0
	boIndex := 0
	addendI := boIndex
	var rSum int
	var lSum int
	for {
		if i >= l-1 {
			break
		} // current window
		sz := l - i
		//fmt.Printf("Current window size is (inclusive) %v\n", sz)

		// move the window
		eoIndex := boIndex + sz - 1
		//fmt.Printf("Window index span is from %d to %d\n", boIndex, eoIndex)
		//fmt.Printf("Current addend is s[%d] and has a value of %d\n", addendI, s[addendI])
		if boIndex == addendI {
			//fmt.Printf("Setting intial lSum to %d\n", s[addendI])
			lSum = s[addendI]
		} else {
			//fmt.Printf("Adding current added %d to lSum %d\n", s[addendI], lSum)

			lSum += s[addendI]
		}

		if addendI >= eoIndex {
			if i == 0 { // this is our first itration, we don't want to evaluate a 0, so we'll just set it
				//fmt.Printf("Setting first rSum to %d\n", lSum)
				rSum = lSum
			} else {
				//fmt.Printf("Evaluating max of (rSum) %d and (lSum) %d\n", rSum, lSum)
				rSum = Max(rSum, lSum)
			}

			if eoIndex >= l-1 {
				i += 1
				boIndex = 0
				//fmt.Printf("\n\n Iteration Ends\n\n")
			} else {
				boIndex += 1
			}
			addendI = boIndex
		} else {
			addendI += 1
		}

	}

	//	return -1, fmt.Errorf("this function isn't functional")
	return rSum, nil
}

// I could never get this one to work with all negative values
func FindMaxB(s []int) int {
	var mEnd, mS int
	//9:  {-1, -4, 4, 4, -9, 10},             // this was just wrong... it WAS 9 [4,4,-9,10]

	for _, n := range s {
		mEnd = Max(0, mEnd+n)
		mS = Max(mS, mEnd)

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
