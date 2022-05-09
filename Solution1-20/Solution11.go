package main

func maxArea(height []int) int {
	n := len(height)
	cursorA, cursorB := 0, n-1
	result := 0
	for cursorA < cursorB {
		tmp := min(height[cursorA], height[cursorB]) * (cursorB - cursorA)
		if tmp > result {
			result = tmp
		}
		if height[cursorA] > height[cursorB] {
			cursorB--
		} else {
			cursorA++
		}
	}
	return result
}
