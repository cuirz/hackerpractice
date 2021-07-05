package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

//726. 原子的数量
//给定一个化学式formula（作为字符串），返回每种原子的数量。
//
//原子总是以一个大写字母开始，接着跟随0个或任意个小写字母，表示原子的名字。
//
//如果数量大于 1，原子后会跟着数字表示原子的数量。如果数量等于 1 则不会跟数字。例如，H2O 和 H2O2 是可行的，但 H1O2 这个表达是不可行的。
//
//两个化学式连在一起是新的化学式。例如 H2O2He3Mg4 也是化学式。
//
//一个括号中的化学式和数字（可选择性添加）也是化学式。例如 (H2O2) 和 (H2O2)3 是化学式。
//
//给定一个化学式，输出所有原子的数量。格式为：第一个（按字典序）原子的名子，跟着它的数量（如果数量大于 1），然后是第二个原子的名字（按字典序），跟着它的数量（如果数量大于 1），以此类推。
//
//示例 1:
//
//输入:
//formula = "H2O"
//输出: "H2O"
//解释:
//原子的数量是 {'H': 2, 'O': 1}。
//示例 2:
//
//输入:
//formula = "Mg(OH)2"
//输出: "H2MgO2"
//解释:
//原子的数量是 {'H': 2, 'Mg': 1, 'O': 2}。
//示例 3:
//
//输入:
//formula = "K4(ON(SO3)2)2"
//输出: "K4N2O14S4"
//解释:
//原子的数量是 {'K': 4, 'N': 2, 'O': 14, 'S': 4}。
//注意:
//
//所有原子的第一个字母为大写，剩余字母都是小写。
//formula的长度在[1, 1000]之间。
//formula只包含字母、数字和圆括号，并且题目中给定的是合法的化学式。

//思路 倒序遍历，栈，哈希表
func countOfAtoms(formula string) string {
	dic := make(map[string]int)
	array := make([]string, 0)
	stack := make([]int, 0)
	n := len(formula)
	count, h := 1, 0
	var elm []byte
	for i := n - 1; i > -1; i-- {
		ch := formula[i]
		switch {
		case ch >= '0' && ch <= '9':
			if h == 0 {
				count = 0
			}
			count += int(ch-'0') * int(math.Pow10(h))
			h++
		case ch == ')':
			if len(stack) > 0 {
				count *= stack[len(stack)-1]
			}
			stack = append(stack, count)
			count, h = 1, 0
		case ch == '(':
			stack = stack[:len(stack)-1]
		case ch >= 'a' && ch <= 'z':
			elm = append([]byte{ch}, elm...)
		default:
			elm = append([]byte{ch}, elm...)
			str := string(elm)
			if len(stack) > 0 {
				count *= stack[len(stack)-1]
			}
			if dic[str] == 0{
				array = append(array, str)
			}
			dic[str] += count

			count, h = 1, 0
			elm = []byte{}
		}
	}
	sort.Strings(array)
	var result strings.Builder
	for _, v := range array {
		result.WriteString(v)
		if dic[v] > 1 {
			result.WriteString(strconv.Itoa(dic[v]))
		}
	}
	return result.String()
}

func main() {
	println(countOfAtoms("K4(ON(SO3)2)2"))
}
