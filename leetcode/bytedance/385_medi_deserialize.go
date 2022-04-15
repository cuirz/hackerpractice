package main

import (
	"unicode"
)

//385. 迷你语法分析器
//给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果NestedInteger 。
//
//列表中的每个元素只可能是整数或整数嵌套列表
//
//
//
//示例 1：
//
//输入：s = "324",
//输出：324
//解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
//示例 2：
//
//输入：s = "[123,[456,[789]]]",
//输出：[123,[456,[789]]]
//解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
//1. 一个 integer 包含值 123
//2. 一个包含两个元素的嵌套列表：
//i.  一个 integer 包含值 456
//ii. 一个包含一个元素的嵌套列表
//a. 一个 integer 包含值 789
//
//
//提示：
//
//1 <= s.length <= 5 * 10^4
//s 由数字、方括号"[]"、负号'-'、逗号','组成
//用例保证s 是可解析的NestedInteger
//输入中的所有值的范围是[-10^6, 10^6]

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	var index int
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[index] == '[' {
			index++
			for s[index] != ']' {
				ni.Add(*dfs())
				if s[index] == ',' {
					index++
				}
			}
			index++
			return ni
		}
		sign := 1
		if s[index] == '-' {
			index++
			sign = -1
		}
		num := 0
		for index < len(s) && unicode.IsDigit(rune(s[index])) {
			num = num*10 + int(s[index]-'0')
			index++
		}
		ni.SetInteger(num * sign)
		return ni
	}
	return dfs()
}
