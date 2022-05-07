package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		value, contains := m[nums[i]]
		if contains {
			return []int{value, i}
		} else {
			m[target-nums[i]] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 13}, 9))
}
