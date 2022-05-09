package main

func romanToInt(s string) int {
	value := []int{1, 5, 10, 50, 100, 500, 1000}
	result := 0
	for i := 0; i < len(s); i++ {
		idx := getPos(s[i])
		if i+1 < len(s) && getPos(s[i+1]) > idx {
			result -= value[idx]
		} else {
			result += value[idx]
		}
	}
	return result
}

func getPos(b byte) int {
	key := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	for idx, value := range key {
		if b == value {
			return idx
		}
	}
	return -1
}
