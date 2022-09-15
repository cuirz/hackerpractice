package main

//现有一个房间，墙上挂有n只已经打开的灯泡和 4 个按钮。在进行了m次未知操作后，你需要返回这n只灯泡可能有多少种不同的状态。
//
//假设这 n 只灯泡被编号为 [1, 2, 3 ..., n]，这 4 个按钮的功能如下：
//
//将所有灯泡的状态反转（即开变为关，关变为开）
//将编号为偶数的灯泡的状态反转
//将编号为奇数的灯泡的状态反转
//将编号为 3k+1 的灯泡的状态反转（k = 0, 1, 2, ...)
//示例 1:
//
//输入: n = 1, m = 1.
//输出: 2
//说明: 状态为: [开], [关]
//示例 2:
//
//输入: n = 2, m = 1.
//输出: 3
//说明: 状态为: [开, 关], [关, 开], [关, 关]
//示例 3:
//
//输入: n = 3, m = 1.
//输出: 4
//说明: 状态为: [关, 开, 关], [开, 关, 开], [关, 关, 关], [关, 开, 开].

func flipLights(n int, m int) int {
	result := make(map[int]struct{})
	for i := 0; i < 1<<4; i++ {
		pressArr := [4]int{}
		sum := 0
		for j := 0; j < 4; j++ {
			pressArr[j] = i >> j & 1
			sum += pressArr[j]
		}
		if sum%2 == m%2 && sum <= m {
			status := pressArr[0] ^ pressArr[2] ^ pressArr[3]
			if n >= 2 {
				status |= (pressArr[0] ^ pressArr[1]) << 1
			}
			if n >= 3 {
				status |= (pressArr[0] ^ pressArr[2]) << 2
			}
			if n >= 4 {
				status |= (pressArr[0] ^ pressArr[1] ^ pressArr[3]) << 3
			}
			result[status] = struct{}{}
		}
	}
	return len(result)

}
