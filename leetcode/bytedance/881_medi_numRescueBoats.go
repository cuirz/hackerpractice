package main

import "sort"

//881. 救生艇
//第i个人的体重为people[i]，每艘船可以承载的最大重量为limit。
//
//每艘船最多可同时载两人，但条件是这些人的重量之和最多为limit。
//
//返回载到每一个人所需的最小船数。(保证每个人都能被船载)。
//
//
//
//示例 1：
//
//输入：people = [1,2], limit = 3
//输出：1
//解释：1 艘船载 (1, 2)
//示例 2：
//
//输入：people = [3,2,2,1], limit = 3
//输出：3
//解释：3 艘船分别载 (1, 2), (2) 和 (3)
//示例 3：
//
//输入：people = [3,5,3,4], limit = 5
//输出：4
//解释：4 艘船分别载 (3), (3), (4), (5)
//提示：
//
//1 <=people.length <= 50000
//1 <= people[i] <=limit <= 30000

//思路 贪心 排序 双指针
func numRescueBoats(people []int, limit int) int {
	n := len(people)
	result := n
	sort.Ints(people)
	l, r := 0, n-1
	for l < r {
		if people[l] == limit {
			break
		}
		if people[r]+people[l] <= limit {
			result--
			l++
		}
		r--
	}
	return result

}
