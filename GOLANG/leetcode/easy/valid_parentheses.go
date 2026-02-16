package main

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stack = append(stack, ch)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (top == '(' && ch != ')') ||
				(top == '{' && ch != '}') ||
				(top == '[' && ch != ']') {
				return false
			}
		}
	}

	return len(stack) == 0
}
