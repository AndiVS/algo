package stack

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mapping[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	//for _, r := range s {
	//	if r == '(' || r == '{' || r == '[' {
	//		stack = append(stack, r)
	//	} else {
	//		if len(stack) == 0 {
	//			return false
	//		}
	//
	//		switch len(stack) - 1 {
	//		case '(':
	//			if r == ')' {
	//				continue
	//			}
	//		case '{':
	//			if r == '}' {
	//				continue
	//			}
	//		case '[':
	//			if r == ']' {
	//				continue
	//			}
	//		default:
	//			return false
	//		}
	//		stack = stack[:len(stack)-1]
	//	}
	//}

	return len(stack) == 0
}
