package main

import (
	"strconv"
	"strings"
)

//
//423. 从英文中重建数字
//给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
//
//
//
//示例 1：
//
//输入：s = "owoztneoer"
//输出："012"
//示例 2：
//
//输入：s = "fviefuro"
//输出："45"
//
//
//提示：
//
//1 <= s.length <= 10^5
//s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
//s 保证是一个符合题目要求的字符串

//zero,one,two,three,four,five,six,seven,eight,nine
//e : (zero),one,(three),[five],[seven],(eight),{nine}
//g : eight
//f : (four),[five]
//i : [five],(six),(eight),{nine}
//h : [three],(eight)
//o : (zero),one,(two),(four)
//n : one,[seven],{nine}
//s : (six),[seven]
//r : (zero),(three),(four)
//u : four
//t : (two),(three),(eight)
//w : two
//v : [five],[seven]
//x : six
//z : zero

//第一步 z: zero,x: six,w: two,u: four,g: eight
//第二步 h: three,f: five, s: seven,
//第三步 e: one
//第四步 i: nine
func originalDigits(s string) string {
	nums := make([]int, 10)
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}
	nums[0] = mp['z']
	nums[2] = mp['w']
	nums[4] = mp['u']
	nums[6] = mp['x']
	nums[8] = mp['g']

	nums[3] = mp['h'] - nums[8]
	nums[5] = mp['f'] - nums[4]
	nums[7] = mp['s'] - nums[6]
	nums[1] = mp['o'] - nums[0] - nums[2] - nums[4]
	nums[9] = mp['i'] - nums[5] - nums[6] - nums[8]

	var result string
	for i, v := range nums {
		result += strings.Repeat(strconv.Itoa(i), v)
	}
	return result
}
