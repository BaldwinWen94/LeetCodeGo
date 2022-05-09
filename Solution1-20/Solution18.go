package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

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

	for i := 0; i < n-3; {
		if nums[i] > target/4 {
			break
		}
		for j := i + 1; j < n-2; {
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left = addLeft(left)
					right = addRight(right)
				} else if sum > target {
					right = addRight(right)
				} else {
					left = addLeft(left)
				}
			}

			for j+1 < n && nums[j] == nums[j+1] {
				j++
			}
			j++
		}

		for i+1 < n && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return res
}
