package main

import (
	"fmt"
)

func main() {
	testBasicImport()
	testSlice()
	testChildSlice()
	testMap()
	testTypeConversion()
	num := 8
	testIfA(&num)
	fmt.Println(testFunction(num, num))
}

func testBasicImport() {
	fmt.Println("#####Test Basic Import#####")
}

func testSlice() {
	fmt.Println("#####Test Slice#####")
	var nums0 []int
	nums0 = append(nums0, 1, 2, 3, 4)
	var nums1 = make([]int, 15)
	copy(nums1, nums0)
	fmt.Println(nums0)
	fmt.Println(nums1)
	var nums2 = make([]int, 5, 10)
	fmt.Println(nums2)
	nums2 = append(nums2, 100, 101)
	fmt.Println(nums2)
	nums2[0] = 1
	fmt.Println(nums2)

}

func testChildSlice() {
	fmt.Println("#####Test Child Slice#####")
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	nums1 := nums[3:]
	fmt.Println(nums1)
	var nums2 = nums[:3]
	fmt.Println(nums2)
	var nums3 = nums[2:3]
	fmt.Println(nums3)
}

func testMap() {
	fmt.Println("#####Test Map#####")
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m["a"])
	fmt.Println(m["c"])
}

func testTypeConversion() {
	fmt.Println("#####Test Type Conversion#####")
	a := 1.1
	b := int(a)
	fmt.Println(a)
	fmt.Println(b)
}

func testIfA(input *int) {
	fmt.Println("#####Test If#####")
	if *input < 10 {
		*input++
		fmt.Println(*input)
	} else {
		fmt.Println(*input)
	}
}

func testFunction(a int, b int) int {
	fmt.Println("#####Test Function#####")
	return a + b
}
