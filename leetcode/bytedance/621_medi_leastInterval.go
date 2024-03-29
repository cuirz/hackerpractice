package main

import "sort"

//621. 任务调度器
//给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
//
//然而，两个相同种类的任务之间必须有长度为n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//
//你需要计算完成所有任务所需要的最短时间。
//
//
//
//示例 ：
//
//输入：tasks = ["A","A","A","B","B","B"], n = 2
//输出：8
//解释：A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
//在本示例中，两个相同类型任务之间必须间隔长度为 n = 2 的冷却时间，而执行一个任务只需要一个单位时间，所以中间出现了（待命）状态。
//
//
//提示：
//
//任务的总个数为[1, 10000]。
//n 的取值范围为 [0, 100]。

//思路  排序

func leastInterval(tasks []byte, n int) int {
	arrays := make([]int, 0)
	dic := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		if dic[tasks[i]] == 0 {
			arrays = append(arrays, 1)
			dic[tasks[i]] = len(arrays)
		} else {
			arrays[dic[tasks[i]]-1]++
		}
	}
	sort.Ints(arrays)
	count := 0
	size := len(arrays)
	for arrays[size-1] > 0 {
		i := 0
		for i <= n {
			if arrays[size-1] == 0 {
				break
			}
			if i < size && arrays[size-i-1] > 0 {
				arrays[size-i-1]--
			}
			i++
			count++
		}
		sort.Ints(arrays)
	}
	return count

}

func main() {

	println(leastInterval([]byte{
		//'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
		'A', 'A', 'A', 'B', 'B', 'B',
	}, 2))
}
