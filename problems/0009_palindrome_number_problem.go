package problems

import "fmt"

/*
9. Palindrome a number

Given an integer x, return true if x is palindrome integer.

An integer is a palindrome when it reads the same backward as forward.

For example, 121 is a palindrome while 123 is not.


Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1

*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reverse int = reverse(x)

	return x == reverse
}

func printPalindrome(x int) {
	fmt.Println(fmt.Sprintf("Number: %d Reverse: %d Palindrome: %v", x, reverse(x), isPalindrome(x)))
}

func Problem9() {
	printPalindrome(0)
	printPalindrome(121)
	printPalindrome(-121)
	printPalindrome(10)
	printPalindrome(123)
}
