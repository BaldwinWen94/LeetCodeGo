package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var result []int
	for x > 0 {
		result = append(result, x%10)
		x = x / 10
	}

	n := len(result)
	for i := 0; 2*i <= n-1; i++ {
		if result[i] != result[n-1-i] {
			return false
		}
	}
	return true
}
