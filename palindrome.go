// https://leetcode.com/problems/palindrome-number/description/
package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// Handling Negative numbers
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversedHalf := 0
	// Reverse half of the number
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// Check if the original number is equal to the reversed half
	return x == reversedHalf || x == reversedHalf/10
}

func main() {
	// Test cases
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
	fmt.Println(isPalindrome(10))   // Output: false
	fmt.Println(isPalindrome(0))    // Output: true
}
