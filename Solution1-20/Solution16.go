package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result, n := nums[0]+nums[1]+nums[2], len(nums)

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			sums := nums[i] + nums[left] + nums[right]
			if sums == target {
				return sums
			} else {
				if abs(target-sums) < abs(target-result) {
					result = sums
				}

				if sums > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
