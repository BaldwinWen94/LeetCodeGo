package main

func lengthOfLongestSubstring(s string) int {
	result := 0
	left, right := 0, 0
	for right < len(s) {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				result = max(result, right-left)
				left = i + 1
				break
			}
		}
		right++
	}
	result = max(result, right-left)
	return result
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
