package main

//1109. 航班预订统计
//这里有n个航班，它们分别从 1 到 n 进行编号。
//
//有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。
//
//请你返回一个长度为 n 的数组answer，其中 answer[i] 是航班 i 上预订的座位总数。
//
//
//
//示例 1：
//
//输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
//输出：[10,55,45,25,25]
//解释：
//航班编号        1   2   3   4   5
//预订记录 1 ：   10  10
//预订记录 2 ：       20  20
//预订记录 3 ：       25  25  25  25
//总座位数：      10  55  45  25  25
//因此，answer = [10,55,45,25,25]
//示例 2：
//
//输入：bookings = [[1,2,10],[2,2,15]], n = 2
//输出：[10,25]
//解释：
//航班编号        1   2
//预订记录 1 ：   10  10
//预订记录 2 ：       15
//总座位数：      10  25
//因此，answer = [10,25]
//
//
//提示：
//
//1 <= n <= 2 * 10^4
//1 <= bookings.length <= 2 * 10^4
//bookings[i].length == 3
//1 <= firsti <= lasti <= n
//1 <= seatsi <= 10^4

//思路 差分
func corpFlightBookings(bookings [][]int, n int) []int {
	ans := make([]int, n)
	for _, bs := range bookings {
		ans[bs[0]-1] += bs[2]
		if bs[1] < n {
			ans[bs[1]] -= bs[2]
		}
	}
	for i := 1; i < n; i++ {
		ans[i] += ans[i-1]
	}
	return ans
}
