package main

//一次编辑
//字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
//示例1:
//
//输入:
//first = "pale"
//second = "ple"
//输出: True
//
//
//示例2:
//
//输入:
//first = "pales"
//second = "pal"
//输出: False

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i := 0; i < n; i++ {
		if first[i] != second[i] {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true

}
