package main

import "bytes"

type cMap struct {
	Value int
	Key   string
}

func intToRoman(num int) string {
	m := []cMap{{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}}

	result := bytes.Buffer{}
	for i := 0; i < len(m) && num > 0; i++ {
		for num >= m[i].Value {
			result.WriteString(m[i].Key)
			num -= m[i].Value
		}
	}
	return result.String()
}
