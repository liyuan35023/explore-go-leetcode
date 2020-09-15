package _0_valid_parentheses
/*
	Given a string containing just the characters ‘(’, ‘)’, ‘{’, ‘}’, ‘[’ and ‘]’, determine if the input string is valid.

	An input string is valid if:

	Open brackets must be closed by the same type of brackets. Open brackets must be closed in the correct order.
	Note that an empty string is also considered valid.

 */
import "github.com/liyuan35023/utils/datastru"

// use stack impl

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	parent := map[byte]byte{
		'(':')',
		'{':'}',
		'[':']',
	}
	stack := datastru.NewStack()
	for i := 0; i < len(s); i++ {
		if _, ok := parent[s[i]]; ok {
			stack.Push(s[i])
		} else {
			if stack.Empty() || parent[stack.Pop().(byte)] != s[i] {
				return false
			}
		}
	}
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}
}