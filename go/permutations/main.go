package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Permutations..")
	word := "aabb"
	perm := generatePermutations(word)
	sort.Strings(perm)
	fmt.Print("should return : [")
	for _, val := range perm {
		fmt.Printf("%v ", val)
	}
	fmt.Print("]\n")
}

func generatePermutations(s string) []string {
	chars := []rune(s)
	var result []string
	backtrack(chars, 0, &result)
	return result
}

func backtrack(chars []rune, start int, result *[]string) {
	if start == len(chars) {
		*result = append(*result, string(chars))
		return
	}

	seen := make(map[rune]bool)
	for i := start; i < len(chars); i++ {
		if _, ok := seen[chars[i]]; ok {
			continue
		}
		seen[chars[i]] = true

		chars[start], chars[i] = chars[i], chars[start]
		backtrack(chars, start+1, result)
		chars[start], chars[i] = chars[i], chars[start]
	}
}
