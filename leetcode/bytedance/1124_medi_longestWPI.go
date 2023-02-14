package main

//1124. 表现良好的最长时间段
//给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
//
//我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
//
//所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
//
//请你返回「表现良好时间段」的最大长度。
//
//
//
//示例 1：
//
//输入：hours = [9,9,6,0,6,6,9]
//输出：3
//解释：最长的表现良好时间段是 [9,9,6]。
//示例 2：
//
//输入：hours = [6,6,6]
//输出：0
//
//
//提示：
//
//1 <= hours.length <= 10^4
//0 <= hours[i] <= 16
func longestWPI(hours []int) int {
	goodDay := 0
	var result int
	position := make(map[int]int)
	for i, h := range hours {
		if h > 8 {
			goodDay++
		} else {
			goodDay--
		}
		if goodDay > 0 {
			result = i + 1
		} else if j, ok := position[goodDay-1]; ok {
			result = max(result, i-j)
		}
		if _, ok := position[goodDay]; !ok {
			position[goodDay] = i
		}
	}
	return result
}
