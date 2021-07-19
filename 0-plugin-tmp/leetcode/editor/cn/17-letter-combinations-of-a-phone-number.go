package cn

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 示例 1： 
//
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2： 
//输入：digits = ""
//输出：[]
// 
// 示例 3：
//
//输入：digits = "2"
//输出：["a","b","c"]
//
// 提示： 
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。 
// 


//leetcode submit region begin(Prohibit modification and deletion)
var _map = map[byte]string{
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	ans := make([]string, 0)
	var dfs func(i int, cur string)
	dfs = func(i int, cur string) {
		if i == len(digits) {
			ans = append(ans, cur)
			return
		}
		for _, r := range _map[digits[i]] {
			dfs(i+1, cur+string(r))
		}
	}
	dfs(0, "")
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
