package main

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}

	var result []byte

	for i := 1; i <= numRows; i++ {
		for tmp := i - 1; tmp < len(s); tmp += 2*numRows - 2 {
			result = append(result, s[tmp])

			if i > 1 && i < numRows {
				tmpRight := tmp + 2*(numRows-i)
				if tmpRight >= len(s) {
					break
				} else {
					result = append(result, s[tmpRight])
				}
			}
		}
	}
	return string(result)
}
