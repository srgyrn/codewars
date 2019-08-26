package valid_parantheses

import "regexp"

func ValidParentheses(parens string) bool {
	for len(regexp.MustCompile(`\(\)`).FindString(parens)) > 0 {
		parens = regexp.MustCompile(`\(\)`).ReplaceAllString(parens, "")
	}

	if len(parens) > 0 {
		return false
	}

	return true
}
