package problems

import (
	"fmt"
	"math"
)

/*
7. Reverse Integer

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the
signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21

Constraints:

-231 <= x <= 231 - 1
*/
func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	var number int = x
	var reverse int = 0

	for number != 0 {
		reverse = reverse*10 + number%10
		number = number / 10
	}

	return reverse
}

func printReverse(x int) {
	fmt.Println(fmt.Sprintf("Number: %d\tReverse: %d", x, reverse(x)))
}

func Problem7() {
	printReverse(123)
	printReverse(math.MaxInt32 + 1)
	printReverse(-123)
	printReverse(120)
	printReverse(1534236469)

	print(math.MaxInt32)
	print(math.MinInt32)

}
