package main

func myAtoi(s string) int {
	n := len(s)
	result, sign := 0, 1
	already := false
	const MaxInt, MinInt = 2147483647, -2147483648

	for i := 0; i < n; i++ {
		switch s[i] {
		case ' ':
			if already {
				return result * sign
			} else {
				continue
			}
		case '-':
			if already {
				return result * sign
			} else {
				sign = -1
				already = true
			}
		case '+':
			if already {
				return result * sign
			} else {
				sign = 1
				already = true
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			already = true
			tmp := int(s[i] - '0')
			if result > MaxInt/10 {
				if sign > 0 {
					return MaxInt
				} else {
					return MinInt
				}
			} else if result == MaxInt/10 && ((sign > 0 && tmp > 7) || (sign < 0 && tmp > 8)) {
				if sign > 0 {
					return MaxInt
				} else {
					return MinInt
				}
			} else {
				result = result*10 + tmp
			}
		default:
			return result * sign
		}
	}

	return result * sign
}
