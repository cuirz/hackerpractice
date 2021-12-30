package main

import "sort"

//846. 一手顺子
//Alice 手中有一把牌，她想要重新排列这些牌，分成若干组，使每一组的牌数都是 groupSize ，并且由 groupSize 张连续的牌组成。
//
//给你一个整数数组 hand 其中 hand[i] 是写在第 i 张牌，和一个整数 groupSize 。如果她可能重新排列这些牌，返回 true ；否则，返回 false 。
//
//
//
//示例 1：
//
//输入：hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
//输出：true
//解释：Alice 手中的牌可以被重新排列为 [1,2,3]，[2,3,4]，[6,7,8]。
//示例 2：
//
//输入：hand = [1,2,3,4,5], groupSize = 4
//输出：false
//解释：Alice 手中的牌无法被重新排列成几个大小为 4 的组。
//
//
//提示：
//
//1 <= hand.length <= 10^4
//0 <= hand[i] <= 10^9
//1 <= groupSize <= hand.length

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	dic := make(map[int]int)
	for _, v := range hand {
		dic[v]++
	}
	sort.Ints(hand)
	for _, v := range hand {
		if dic[v] == 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			index := v + i
			if dic[index] == 0 {
				return false
			}
			dic[index]--
			if dic[index] == 0 {
				delete(dic, index)
			}
		}
	}
	return true
}
