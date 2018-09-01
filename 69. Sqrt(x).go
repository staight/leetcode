/*
69. Sqrt(x)

Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4

Output: 2

Example 2:

Input: 8

Output: 2

Explanation: The square root of 8 is 2.82842..., and since 
             the decimal part is truncated, 2 is returned.
*/

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		middle := (left + right) / 2
		if power(middle) <= x && power(middle+1) > x {
			return middle
		} else if power(middle) > x {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return 0
}

func power(x int) int {
	return x * x
}