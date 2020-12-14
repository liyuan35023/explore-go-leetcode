package _1_simplify_path

/*
	Example 1:

	Input: "/home/"
	Output: "/home"
	Explanation: Note that there is no trailing slash after the last directory name.

	Example 2:

	Input: "/../"
	Output: "/"
	Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.

	Example 3:

	Input: "/home//foo/"
	Output: "/home/foo"
	Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.

	Example 4:

	Input: "/a/./b/../../c/"
	Output: "/c"

	Example 5:

	Input: "/a/../../b/../c//.//"
	Output: "/c"

	Example 6:

	Input: "/a//b////c/d//././/.."
	Output: "/a/b/c"

	题目大意 #
	给出一个 Unix 的文件路径，要求简化这个路径。
	绝对路径
 */

func simplifyPath(path string) string {

	stack := make([]byte, 0)
	stack = append(stack, path[0])
	for i := 1; i < len(path); {
		if path[i] == '.' && path[i-1] != '.' && i < len(path)-2 && path[i+1] == '.' && path[i+2] == '.' {
			cur := i
			for cur < len(path) && path[cur] != '/' {
				stack = append(stack, path[cur])
				cur++
			}
			i = cur
			continue
		}

		if path[i] == '.' && path[i-1] != '.' && i < len(path)-1 && path[i+1] == '.' && (i+2 == len(path) || path[i+2] == '/') {
			// 出栈直到到上一层
			if len(stack) == 1 {
				i += 2
				continue
			}
			stack = stack[:len(stack)-1]
			for len(stack) > 1 && stack[len(stack)-1] != '/' {
				stack = stack[:len(stack)-1]
			}
			i += 2
			continue
		}
		if path[i] == '/' && stack[len(stack)-1] == '/' {
			i++
			continue
		}
		if path[i] == '.' && (i+1 == len(path) || path[i+1] == '/') {
			i++
			continue
		}
		stack = append(stack, path[i])
		i++
	}
	if len(stack) > 1 && stack[len(stack)-1] == '/' {
		return string(stack[:len(stack)-1])
	} else {
		return string(stack)
	}
}

func simplifyPathII(path string) string {
	stack := make([]byte, 0)
	stack = append(stack, path[0])
	for i := 1; i < len(path); {
		if path[i] == '.' && path[i-1] == '/' && (i == len(path)-1 || path[i+1] == '/') {
			// /./
			i = i + 2
			continue
		}

		if path[i] == '.' && path[i-1] == '/' && path[i+1] == '.' && (i == len(path)-2 || path[i+2] == '/') {
			// /../
			// 退到上一级
			if len(stack) != 1 {
				stack = stack[:len(stack)-1]
				for stack[len(stack)-1] != '/' {
					stack = stack[:len(stack)-1]
				}
			}
			i = i + 3
			continue
		}

		if path[i] == '/' && path[i-1] == '/' {
			i = i + 1
			continue
		}
		stack = append(stack, path[i])
		i++
	}
	if len(stack) > 1 && stack[len(stack)-1] == '/' {
		return string(stack[:len(stack)-1])
	} else {
		return string(stack)
	}
}
