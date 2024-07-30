package main

import (
	"reflect"
	"sort"
	"testing"
)

type PermutationStruct struct {
	Name     string
	Input    string
	Expected []string
}

func TestPermutations(t *testing.T) {
	testCases := []PermutationStruct{
		{Name: "Test #1",
			Input:    "a",
			Expected: []string{"a"},
		},
		{Name: "Test #2",
			Input:    "ab",
			Expected: []string{"ab", "ba"},
		},
		{Name: "Test #3",
			Input:    "abc",
			Expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{Name: "Test #4",
			Input:    "aabb",
			Expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := generatePermutations(tt.Input)
			sort.Strings(result)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("got %v, want %v", result, tt.Expected)
			}
		})
	}
}
