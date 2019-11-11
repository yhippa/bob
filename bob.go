// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isQuestion(remark) && isShouting(remark) {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		return "Sure."
	} else if !isPunctuation(remark) && isShouting(remark) {
		return "Whoa, chill out!"
	} else if !isPunctuation(remark) {
		return "Whatever."
	} else if isShouting(remark) {
		return "Whoa, chill out!"
	}
	return "Whatever."
}

func isQuestion(input string) bool {
	if string(input[len(input)-1]) == "?" {
		return true
	}
	return false
}

func isShouting(input string) bool {
	if strings.ToUpper(input) == input {
		return true
	}
	return false
}

func isPunctuation(input string) bool {
	if string(input[len(input)-1]) == "?" || string(input[len(input)-1]) == "!" || string(input[len(input)-1]) == "." {
		return true
	}
	return false
}
