package main

/*
* 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
* 有效字符串需满足：
* 1. 左括号必须用相同类型的右括号闭合。
* 2. 左括号必须以正确的顺序闭合。
* 3. 每个右括号都有一个对应的相同类型的左括号。
*
* 方法：
* 判断括号的有效性可以使用「栈」这一数据结构来解决。
 */
func isValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
