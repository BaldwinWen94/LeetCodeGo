package main

func reverse(x int) int {
	MaxInt := 2147483647 / 10
	result, sign := 0, 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for ; x > 0; x /= 10 {
		if result > MaxInt || (result == MaxInt && x > 7) {
			return 0
		} else {
			result = result*10 + x%10
		}
	}

	return result * sign
}
