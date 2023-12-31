package leedcodeproblm

/*
Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.

A substring is a contiguous sequence of characters within a string.

Example 1:

Input: s = "aa"
Output: 0
Explanation: The optimal substring here is an empty substring between the two 'a's.
Example 2:

Input: s = "abca"
Output: 2
Explanation: The optimal substring here is "bc".
Example 3:

Input: s = "cbzxy"
Output: -1
Explanation: There are no characters that appear twice in s.
*/
func NonOptimalSolution(s string) int {
	if len(s) == 2 {
		return 0
	}
	maxi := -1
	for right, ch := range []byte(s) {
		for j := len(s) - 1; j > right; j-- {
			if ch == s[j] {
				maxi = max(maxi, j-right)
				break
			}
		}
	}
	return maxi
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func OptimalSolution(s string) int {
	firstOccurence := make(map[rune]int)
	maxi := -1
	for i, ch := range s {
		if value, ok := firstOccurence[ch]; ok {
			newValue := i - value - 1
			if newValue > maxi {
				maxi = newValue
			}
		}
	}
	return maxi
}
