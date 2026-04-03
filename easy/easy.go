package easy

// 有效的括号
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	res := make([]rune, 0)

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			res = append(res, c)
		} else {
			if len(res) == 0 {
				return false
			}

			if c == '}' && res[len(res)-1] != '{' {
				return false
			}

			if c == ']' && res[len(res)-1] != '[' {
				return false
			}

			if c == ')' && res[len(res)-1] != '(' {
				return false
			}

			res = res[:len(res)-1]
		}
	}

	return len(res) == 0
}
