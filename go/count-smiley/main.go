package main

import "fmt"

// Rules for a smiling face:
//? Each smiley face must contain a valid pair of eyes. Eyes can be marked as : or ;
//? A smiley face can have a nose but it does not have to. Valid characters for a nose are - or ~
//? Every smiling face must have a smiling mouth that should be marked with either ) or D

// No additional characters are allowed except for those mentioned.

//? Valid smiley face examples:  :) :D ;-D :~)
//! Invalid smiley faces:  ;( :> :} :]

func main() {
	fmt.Println("Count Smiley face :D")
	smileys := []string{":)", ";(", ";}", ":-D", ":_D"}

	smileyfaces := detectSmiley(smileys)
	fmt.Println("smileyfaces : ", smileyfaces)
}

func detectSmiley(smileys []string) int {
	count := 0
	for _, val := range smileys {
		if len(val) > 3 || len(val) < 2 {
			fmt.Println("invalid smile")
			continue
		}
		hasEyes := false
		hasMouth := false
		hasNose := false
		for _, r := range val {
			switch r {
			case ';', ':':
				hasEyes = true
			case '-', '~':
				hasNose = true
			case ')', 'D':
				hasMouth = true
			}
		}
		if hasEyes && hasMouth && (len(val) == 2 || (len(val) == 3 && hasNose)) {
			count++
		}
	}
	return count
}
