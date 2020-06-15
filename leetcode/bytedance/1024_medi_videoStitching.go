package main

import "math"

//1024. 视频拼接
//你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。
//
//视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。
//
//我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。
//
//
//
//示例 1：
//
//输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
//输出：3
//解释：
//我们选中 [0,2], [8,10], [1,9] 这三个片段。
//然后，按下面的方案重制比赛片段：
//将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
//现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。

// 方案 贪心算法
// 区间问题 贪心算法

func videoStitching(clips [][]int, T int) int {
	n := len(clips)
	result := make([]int, T+1)
	for i := 0; i < n; i++ {
		l := clips[i][0]
		r := clips[i][1]
		if l <= T {
			result[l] = max(result[l], r)
		}
	}

	seek := 0
	big := math.MinInt32
	count := 0
	for i := 0; i < T; i++ {
		big = max(big, result[i])
		if i == seek {
			if big <= i {
				return -1
			}
			seek = big
			count++
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
