package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	if nums == nil || len(nums) < 3 {
		return result
	}

	addLeft := func(left int) int {
		for left < n-1 && nums[left] == nums[left+1] {
			left++
		}
		left++
		return left
	}

	addRight := func(right int) int {
		for right > 0 && nums[right] == nums[right-1] {
			right--
		}
		right--
		return right
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sums := nums[i] + nums[left] + nums[right]
			if sums == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left = addLeft(left)
				right = addRight(right)
			} else if sums > 0 {
				right = addRight(right)
			} else {
				left = addLeft(left)
			}
		}
	}
	return result
}

//Slow
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var result [][]int

	if nums == nil || len(nums) < 3 {
		return result
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		m := make(map[int]int)
		n1 := nums[i]
		for j := i + 1; j < n; j++ {
			tmp := nums[j]
			value, contains := m[tmp]
			if contains {
				if value != -1 {
					result = append(result, []int{n1, nums[value], tmp})
					m[tmp] = -1
				}
			} else {
				m[-tmp-n1] = j
			}
		}
	}
	return result
}
