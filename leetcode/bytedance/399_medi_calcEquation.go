package main

//399. 除法求值
//给出方程式A / B = k, 其中A 和B 均为用字符串表示的变量，k 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回-1.0。
//
//示例 :
//给定a / b = 2.0, b / c = 3.0
//问题: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
//返回[6.0, 0.5, -1.0, 1.0, -1.0 ]
//
//输入为: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries(方程式，方程式结果，问题方程式)，
//其中equations.size() == values.size()，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。返回vector<double>类型。
//
//基于上述例子，输入如下：
//
//equations(方程式) = [ ["a", "b"], ["b", "c"] ],
//values(方程式结果) = [2.0, 3.0],
//queries(问题方程式) = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].
//输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。

//思路 并查集
// union & find
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	type node struct {
		parent string
		value  float64
	}
	dic := make(map[string]*node)
	nums := make(map[string]int)

	var find func(x string) string
	find = func(x string) string {
		if x != dic[x].parent {
			parent := dic[x].parent
			dic[x].parent = find(parent)
			dic[x].value *= dic[parent].value
			return dic[x].parent
		}
		return x
	}
	union := func(v []string, value float64) {
		p1 := find(v[0])
		p2 := find(v[1])
		if p1 != p2 {
			dic[p1].parent = p2
			dic[p1].value = value * dic[v[1]].value / dic[v[0]].value
		}
	}

	for i, item := range equations {
		for _, v := range item {
			if nums[v] == 0 {
				nums[v] = 1
				dic[v] = &node{parent: v, value: 1}
			}
		}
		union(item, values[i])
	}

	result := make([]float64, 0)
	for _, v := range queries {
		a, b := v[0], v[1]
		if nums[a] == 0 || nums[b] == 0 {
			result = append(result, -1.0)
			continue
		}
		p1, p2 := find(a), find(b)
		if p1 != p2 {
			result = append(result, -1.0)
			continue
		}
		result = append(result, dic[a].value/dic[b].value)
	}
	return result
}
