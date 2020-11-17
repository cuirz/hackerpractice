package main

//844. 比较含退格的字符串
//给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
//
//注意：如果对空文本输入退格字符，文本继续为空。

func backspaceCompare(S string, T string) bool {
	i := len(S) - 1
	j := len(T) - 1
	count := 0
	for {
		count = 0
		for i >=0 {
			if S[i] == '#'{
				count ++
				i --
			}else if count > 0{
				count --
				i --
			}else{
				break
			}
		}
		count = 0
		for j >= 0{
			if T[j] == '#'{
				count++
				j --
			}else if count > 0{
				count --
				j --
			}else{
				break
			}
		}
		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 || S[i] != T[j] {
			return false
		}
		i--
		j--
	}
}
