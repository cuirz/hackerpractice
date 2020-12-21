package main

//1370. 上升下降字符串
//给你一个字符串 s ，请你根据下面的算法重新构造字符串：
//
//从 s 中选出 最小 的字符，将它 接在 结果字符串的后面。
//从 s 剩余字符中选出 最小 的字符，且该字符比上一个添加的字符大，将它 接在 结果字符串后面。
//重复步骤 2 ，直到你没法从 s 中选择字符。
//从 s 中选出 最大 的字符，将它 接在 结果字符串的后面。
//从 s 剩余字符中选出 最大 的字符，且该字符比上一个添加的字符小，将它 接在 结果字符串后面。
//重复步骤 5 ，直到你没法从 s 中选择字符。
//重复步骤 1 到 6 ，直到 s 中所有字符都已经被选过。
//在任何一步中，如果最小或者最大字符不止一个 ，你可以选择其中任意一个，并将其添加到结果字符串。
//
//请你返回将 s 中字符重新排序后的 结果字符串 。

func sortString(s string) string {
	dic := make([]int, 26)
	for _, v := range s {
		dic[v-'a'] ++
	}
	n := len(s)
	result := make([]byte, 0, n)
	for len(result) < n {
		for i := 0; i < 26; i++ {
			if dic[i] > 0{
				result = append(result,byte(i+'a'))
				dic[i] --
			}
		}
		for i := 25; i > -1; i-- {
			if dic[i] > 0{
				result = append(result,byte(i+'a'))
				dic[i] --
			}
		}
	}
	return string(result)
}
