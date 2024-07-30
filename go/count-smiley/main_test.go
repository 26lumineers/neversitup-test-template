package main

import (
	"fmt"
	"testing"
)

type SmileyFaceStruct struct {
	Name     string
	Input    []string
	Expected int
}

func TestSmileyFaces(t *testing.T) {
	testCases := []SmileyFaceStruct{
		{Name: "Test #1",
			Input:    []string{":)", ";(", ";}", ":-D"},
			Expected: 2,
		},
		{Name: "Test #2",
			Input:    []string{";D", ":-(", ":-)", ";~)"},
			Expected: 3,
		},
		{Name: "Test #3",
			Input:    []string{";]", ":[", ";*", ":$", ";-D"},
			Expected: 1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			result := detectSmiley(tt.Input)
			if result != tt.Expected {
				t.Errorf("got %d, wanted %d", result, tt.Expected)
			}
			fmt.Printf("countSmileys %s = %d\n", tt.Input, tt.Expected)
		})
	}
}
