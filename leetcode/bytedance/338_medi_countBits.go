package main

//338. 比特位计数
//给定一个非负整数num。对于0 ≤ i ≤ num 范围中的每个数字i，计算其二进制数中的 1 的数目并将它们作为数组返回。
//
//示例 1:
//
//输入: 2
//输出: [0,1,1]
//示例2:
//
//输入: 5
//输出: [0,1,1,2,1,2]
//进阶:
//
//给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
//要求算法的空间复杂度为O(n)。
//你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的__builtin_popcount）来执行此操作。

//思路   n &= n-1 计算1的个数
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		count := 0
		n := i
		for n != 0 {
			if result[n] > 0 {
				count += result[n]
				break
			}
			n = n & (n - 1)
			count++
		}
		result[i] = count
	}
	return result
}

//思路  动态规划
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = result[i>>1] + (i & 1)
		// 或者
		//result[i] = result[i&(i-1)] + 1
	}
	return result

}

func main() {
	for i := 0; i < 5; i++ {

		println(i & (i - 1))
	}

}
