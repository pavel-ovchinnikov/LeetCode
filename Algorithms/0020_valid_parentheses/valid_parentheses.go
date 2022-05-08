package valid_parentheses_0020

func isValid(s string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}

	top := 0
	stack := make([]byte, size)

	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack[top] = byte(r)
			top++
		case ')':
			if top > 0 && stack[top-1] == '(' {
				top--
			} else {
				return false
			}
		case ']':
			if top > 0 && stack[top-1] == '[' {
				top--
			} else {
				return false
			}
		case '}':
			if top > 0 && stack[top-1] == '{' {
				top--
			} else {
				return false
			}

		}
	}
	return top == 0
}
