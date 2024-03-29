package main

//39. 组合总和
//给定一个无重复元素的数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
//
//candidates中的数字可以无限制重复被选取。
//
//说明：
//
//所有数字（包括target）都是正整数。
//解集不能包含重复的组合。
//示例1：
//
//输入：candidates = [2,3,6,7], target = 7,
//所求解集为：
//[
//[7],
//[2,2,3]
//]
//示例2：
//
//输入：candidates = [2,3,5], target = 8,
//所求解集为：
//[
// [2,2,2,2],
// [2,3,3],
// [3,5]
//]
//
//
//提示：
//
//1 <= candidates.length <= 30
//1 <= candidates[i] <= 200
//candidate 中的每个元素都是独一无二的。
//1 <= target <= 500

//思路 回溯+剪枝
// 为了把重复内容排除，当做根节点的数不再使用

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	result := make([][]int, 0)

	var recursive func(t, index int, array []int)
	recursive = func(t, index int, array []int) {
		for i := index; i < n; i++ {
			t -= candidates[i]
			if t > 0 {
				recursive(t, i, append([]int{candidates[i]}, array...))
			} else if t == 0 {
				result = append(result, append([]int{candidates[i]}, array...))
			}
			t += candidates[i]
		}
	}
	recursive(target, 0, []int{})
	return result

}
