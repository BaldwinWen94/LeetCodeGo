package main

func longestPalindrome(s string) string {
	var max, n = 0, len(s)
	if n <= 1 {
		return s
	}
	var result = ""
	for i := 0; i < n; i++ {
		length1 := 0
		for i-length1-1 >= 0 && i+length1+1 < n && s[i-length1-1] == s[i+length1+1] {
			length1++
		}
		if 2*length1+1 > max {
			max = 2*length1 + 1
			result = s[i-length1 : i+length1+1]
		}

		if i+1 < n && s[i] == s[i+1] {
			length2 := 0
			for i-length2-1 >= 0 && i+length2+2 < n && s[i-length2-1] == s[i+length1+2] {
				length2++
			}
			if 2*length2+2 > max {
				max = 2*length2 + 2
				result = s[i-length2 : i+length2+2]
			}
		}
	}
	return result
}
