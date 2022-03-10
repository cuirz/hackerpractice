package main

//171. Excel表列序号
//给你一个字符串columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
//
//
//
//例如，
//
//A -> 1
//B -> 2
//C -> 3
//...
//Z -> 26
//AA -> 27
//AB -> 28
//...
//
//
//示例 1:
//
//输入: columnTitle = "A"
//输出: 1
//示例2:
//
//输入: columnTitle = "AB"
//输出: 28
//示例3:
//
//输入: columnTitle = "ZY"
//输出: 701
//示例 4:
//
//输入: columnTitle = "FXSHRXW"
//输出: 2147483647
//
//
//提示：
//
//1 <= columnTitle.length <= 7
//columnTitle 仅由大写英文组成
//columnTitle 在范围 ["A", "FXSHRXW"] 内

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	var pow func(m, n int) int
	pow = func(m, n int) int {
		if n == 0 {
			return 1
		}
		half := pow(m, n/2)
		if n%2 == 0 {
			return half * half
		} else {
			return half * half * m
		}
	}
	var result int
	for i := 0; i < n; i++ {
		result += int(columnTitle[n-1-i]-'A'+1) * pow(26, i)
	}
	return result
}

func main() {
	println(titleToNumber("AB"))
}
