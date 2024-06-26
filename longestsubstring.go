// Q : https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	start, maxLength := 0, 0

	for i, char := range s {
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndexMap[char] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	// Test cases
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring(" "))        // Output: 1
	fmt.Println(lengthOfLongestSubstring("au"))       // Output: 2
	fmt.Println(lengthOfLongestSubstring("dvdf"))     // Output: 3
}
