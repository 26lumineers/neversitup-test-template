package main

import (
	"fmt"
	"reflect"
	"testing"
)

type OddStruct struct {
	Name     string
	Input    []int
	Expected map[int]int
}

func TestFindOddInt(t *testing.T) {
	testCases := []OddStruct{
		{Name: "Test #1",
			Input:    []int{7},
			Expected: map[int]int{7: 1},
		},
		{Name: "Test #2",
			Input:    []int{0},
			Expected: map[int]int{0: 1},
		},
		{Name: "Test #3",
			Input:    []int{1, 1, 2},
			Expected: map[int]int{2: 1},
		},
		{Name: "Test #4",
			Input:    []int{0, 1, 0, 1, 0},
			Expected: map[int]int{0: 3},
		},
		{Name: "Test #5",
			Input:    []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			Expected: map[int]int{4: 1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := findOddInt(tt.Input)
			if !reflect.DeepEqual(result, tt.Expected) {
				t.Errorf("got %v, want %v", result, tt.Expected)
			}
			for key, value := range tt.Expected {
				fmt.Printf("%v should return %d, because it appears %d time(s) (which is odd).\n", tt.Input, key, value)
			}
		})
	}
}
