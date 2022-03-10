package main

import "sort"

//721. 账户合并
//
//给定一个列表 accounts，每个元素 accounts[i]是一个字符串列表，其中第一个元素 accounts[i][0]是名称 (name)，其余元素是 emails 表示该账户的邮箱地址。
//
//现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。
//
//合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按顺序排列的邮箱地址。账户本身可以以任意顺序返回。
//
//
//
//示例 1：
//
//输入：
//accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
//输出：
//[["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
//解释：
//第一个和第三个 John 是同一个人，因为他们有共同的邮箱地址 "johnsmith@mail.com"。
//第二个 John 和 Mary 是不同的人，因为他们的邮箱地址没有被其他帐户使用。
//可以以任何顺序返回这些列表，例如答案 [['Mary'，'mary@mail.com']，['John'，'johnnybravo@mail.com']，
//['John'，'john00@mail.com'，'john_newyork@mail.com'，'johnsmith@mail.com']] 也是正确的。
//
//
//提示：
//
//accounts的长度将在[1，1000]的范围内。
//accounts[i]的长度将在[1，10]的范围内。
//accounts[i][j]的长度将在[1，30]的范围内。

//思路  并查集 + 哈希表
func accountsMerge(accounts [][]string) [][]string {
	mailName := make(map[string]int)
	mailIndex := make(map[string]int)
	for i, v := range accounts {
		for _, m := range v[1:] {
			if _, has := mailIndex[m]; !has {
				mailIndex[m] = len(mailIndex)
				mailName[m] = i
			}
		}
	}
	parent := make([]int, len(mailIndex))
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		from, to = find(from), find(to)
		parent[from] = to
	}

	for _, v := range accounts {
		first := v[1]
		for _, m := range v[2:] {
			union(mailIndex[m], mailIndex[first])
		}
	}
	dic := make(map[int][]string)
	for k, v := range mailIndex {
		index := find(v)
		dic[index] = append(dic[index], k)
	}

	result := make([][]string, 0)
	for _, mails := range dic {
		sort.Strings(mails)
		user := append([]string{accounts[mailName[mails[0]]][0]}, mails...)
		result = append(result, user)
	}
	return result
}
