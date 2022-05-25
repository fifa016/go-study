package l

func IsValid(s string) bool {

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if (s[i] == ')' && (len(stack) == 0 || stack[len(stack)-1] != '(')) ||
			(s[i] == ']' && (len(stack) == 0 || stack[len(stack)-1] != '[')) ||
			(s[i] == '}' && (len(stack) == 0 || stack[len(stack)-1] != '{')) {
			return false
		}
		stack = stack[:len(stack) - 1]
	}

	return len(stack) == 0
}
