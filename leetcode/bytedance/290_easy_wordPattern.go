package main

import "strings"

//290. 单词规律
//给定一种规律 pattern和一个字符串str，判断 str 是否遵循相同的规律。
//
//这里的遵循指完全匹配，例如，pattern里的每个字母和字符串str中的每个非空单词之间存在着双向连接的对应规律。
//
//示例1:
//
//输入: pattern = "abba", str = "dog cat cat dog"
//输出: true
//示例 2:
//
//输入:pattern = "abba", str = "dog cat cat fish"
//输出: false
//示例 3:
//
//输入: pattern = "aaaa", str = "dog cat cat dog"
//输出: false
//示例4:
//
//输入: pattern = "abba", str = "dog dog dog dog"
//输出: false
//说明:
//你可以假设pattern只包含小写字母，str包含了由单个空格分隔的小写字母。

func wordPattern(pattern string, s string) bool {
	dic := make(map[byte]string)
	exist := make(map[string]bool)
	array := strings.Split(s, " ")
	if len(pattern) != len(array) {
		return false
	}
	for i, v := range pattern {
		if str, ok := dic[byte(v)]; ok {
			if str != array[i] {
				return false
			}
		} else {
			if exist[array[i]] {
				return false
			}
			dic[byte(v)] = array[i]
			exist[array[i]] = true
		}
	}
	return true
}
