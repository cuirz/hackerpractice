package main

import (
	"bytes"
	"strings"
)

//6. Z 字形变换
//将一个给定字符串根据给定的行数，以从上往下、从左到右进行Z 字形排列。
//
//比如输入字符串为 "LEETCODEISHIRING"行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例2:
//
//输入: s = "LEETCODEISHIRING", numRows =4
//输出:"LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G

func convert(s string, numRows int) string {
	rows := make([]bytes.Buffer, numRows)
	var index, sign int

	for _, v := range s {
		rows[index].WriteRune(v)
		if numRows == 1 {
			sign = 0
		} else if index == 0 {
			sign = 1
		} else if index == numRows-1 {
			sign = -1
		}
		index = index + sign
	}
	var result strings.Builder
	for _, v := range rows {
		result.Write(v.Bytes())
	}
	return result.String()
}

func main() {
	println(convert("LEETCODEISHIRING", 4))
}
