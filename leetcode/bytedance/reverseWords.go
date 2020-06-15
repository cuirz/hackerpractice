package main

import "strings"

func reverseWords(s string) string {
	result := ""
	size := len(s)
	peek := 0
	for i := 0; i < size; i++ {
		if s[i] != ' ' {
			peek++
		} else {
			if peek > 0 {
				result = s[i-peek:i] + " " + result
			}
			peek = 0
		}
	}
	if peek > 0 {
		result = s[size-peek:size] + " " + result
	}
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	return result
}

func reverseWordsSimple(s string) string {
	//双指针法
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		if s[0] != ' ' {
			return s
		} else {
			return ""
		}
	}
	left := len(s) - 1
	right := len(s) - 1
	var res []string
	//先去掉末尾空格
	for right >= 0 && s[right] == ' ' {
		right--
	}
	if right < 0 {
		return ""
	}
	left = right
	sByte := []byte(s)
	for left >= 0 {
		//left向左找到第一个空格
		for left >= 0 && sByte[left] != ' ' {
			left--
		}
		res = append(res, string(sByte[left+1:right+1]))
		//去掉left左边的空格
		for left >= 0 && sByte[left] == ' ' {
			left--
		}
		right = left
	}
	return strings.Join(res, " ")
}

func main() {
	//println(reverseWords("the sky is blue"))
	println(reverseWords("我 你我是 人的 "))

}
