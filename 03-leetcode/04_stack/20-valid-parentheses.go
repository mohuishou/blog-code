package stack

import "strings"

func isValid(s string) bool {
	ss := strings.Split(s, "")
	var stack []string
	for i := len(ss) - 1; i >= 0; i-- {
		switch ss[i] {
		case ")":
			stack = append(stack, "(")
		case "]":
			stack = append(stack, "[")
		case "}":
			stack = append(stack, "{")
		default:
			if len(stack) == 0 {
				return false
			}
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ss[i] != val {
				return false
			}
		}
	}

	return len(stack) == 0
}
