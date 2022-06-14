/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-09 17:32:12
 */
package l

//从内到外必须是<>,(),[],{}
func IsValidWithPriority(s string) bool {

	nums := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '<':
			nums[i] = 1
		case '>':
			nums[i] = 2
		case '(':
			nums[i] = 3
		case ')':
			nums[i] = 4
		case '[':
			nums[i] = 5
		case ']':
			nums[i] = 6
		case '{':
			nums[i] = 7
		case '}':
			nums[i] = 8
		}
	}

	stack := []int{}

	for i := 0; i < len(s); i++ {

		if len(stack) == 0 {
			stack = append(stack, nums[i])
		} else {
			if nums[i]%2 == 1 {
				if stack[len(stack)-1] >= nums[i] {
					stack = append(stack, nums[i])
				} else {
					return false
				}

			} else {
				if nums[i]-1 == stack[len(stack)-1] {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0

}
func IsValidWithPriority2(s string) bool {

	nums := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '<':
			nums[i] = 1
		case '>':
			nums[i] = 2
		case '(':
			nums[i] = 3
		case ')':
			nums[i] = 4
		case '[':
			nums[i] = 5
		case ']':
			nums[i] = 6
		case '{':
			nums[i] = 7
		case '}':
			nums[i] = 8
		}
	}

	stack := []int{}

	for i := 0; i < len(s); i++ {

		if nums[i]%2 == 1 {
			if len(stack) == 0 || stack[len(stack)-1] >= nums[i] {
				stack = append(stack, nums[i])
			} else {
				return false
			}
		} else {
			if len(stack) == 0 {
				return false
			} else if nums[i]-1 == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return true
}
