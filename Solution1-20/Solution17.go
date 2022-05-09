package main

func letterCombinations(digits string) []string {
	var res []string
	dict := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}

	if len(digits) == 0 {
		return res
	}

	var addDigit func(k int, cur []byte)
	addDigit = func(k int, cur []byte) {
		if k == len(digits) {
			res = append(res, string(cur))
		} else {
			value, _ := dict[digits[k]]
			for i := 0; i < len(value); i++ {
				cur = append(cur, value[i])
				addDigit(k+1, cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	addDigit(0, []byte{})
	return res
}
