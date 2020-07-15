package main

import "fmt"

//17. 电话号码的字母组合
//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//示例:
//
//输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

//思路 回溯
func letterCombinations(digits string) []string {
	dic := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	result := make([]string, 0)
	n := len(digits)
	if n < 1 {
		return []string{}
	}
	var recursive func(pre string, index int)
	recursive = func(pre string, index int) {
		if index == n {
			result = append(result, pre)
			return
		}
		num := string(digits[index])
		for i := 0; i < len(dic[num]); i++ {
			newpre := fmt.Sprintf("%s%s", pre, dic[num][i])
			recursive(newpre, index+1)
		}
	}
	recursive("", 0)
	return result

}

func main() {
	letterCombinations("23")
}
