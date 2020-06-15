package main

//990. 等式方程的可满足性
//给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
//
//只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。
//
//
//
//示例 1：
//
//输入：["a==b","b!=a"]
//输出：false
//解释：如果我们指定，a = 1 且 b = 1，那么可以满足第一个方程，但无法满足第二个方程。没有办法分配变量同时满足这两个方程。
//示例 2：
//
//输出：["b==a","a==b"]
//输入：true
//解释：我们可以指定 a = 1 且 b = 1 以满足满足这两个方程。

//["c==c","f!=a","f==b","b==c"]

// 并查集
func equationsPossible(equations []string) bool {
	var parents [26]byte
	for i, _ := range parents {
		parents[i] = byte(i)
	}
	find := func(index byte) byte {
		for parents[index] != index {
			parents[index] = parents[parents[index]]
			index = parents[index]
		}
		return index
	}
	union := func(i1, i2 byte) {
		parents[find(i1)] = find(i2)
	}

	for _, v := range equations {
		if v[1] == '=' {
			union(v[0]-'a', v[3]-'a')
		}
	}

	for _, v := range equations {
		if v[1] == '!' {
			if find(v[0]-'a') == find(v[3]-'a') {
				return false
			}
		}
	}
	return true
}
