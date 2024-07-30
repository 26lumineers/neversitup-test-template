package main

import "fmt"

func main() {
	arrayInt := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1, 5, 5, 5}

	result := findOddInt(arrayInt)
	for k, val := range result {
		fmt.Printf("should return %v, because it occurs %v times (which is odd)\n", k, val)
	}
}

func findOddInt(items []int) map[int]int {
	intMap := make(map[int]int)
	for _, item := range items {
		intMap[item]++
	}
	oddMap := make(map[int]int)
	for k, val := range intMap {
		if intMap[k]%2 != 0 {
			oddMap[k] = val
		}
	}
	return oddMap
}
