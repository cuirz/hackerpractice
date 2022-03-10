package main

//925. 长按键入
//你的朋友正在使用键盘输入他的名字name。偶尔，在键入字符c时，按键可能会被长按，而字符可能被输入 1 次或多次。
//
//你将会检查键盘输入的字符typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回True。
//
//
//
//示例 1：
//
//输入：name = "alex", typed = "aaleex"
//输出：true
//解释：'alex' 中的 'a' 和 'e' 被长按。

func isLongPressedName(name string, typed string) bool {
	n := len(name)
	m := len(typed)
	if m < n {
		return false
	}
	j := 0
	var pre byte
	for i := 0; i < m; i++ {
		if j < n && name[j] == typed[i] {
			j++
			pre = typed[i]
		} else if pre != typed[i] {
			return false
		}
	}
	return j == n
}
