package main

//1371. 每个元音包含偶数次的最长子字符串
//给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。
//eleetminicoworoep
// pos[e] 					pos[0x10] = 位置(0+1)	res = 0
// pos[el] 					pos[0x10] = 位置(1)		res = max(0,1-1+1)
// pos[ele]					pos[0x0]  = 位置(0)		res = max(1,2-0+1)
// pos[elee]				pos[0x10] = 位置(3+1)    res = 3
// pos[eleet]				pos[0x10] = 位置(4)		res = max(3,4-4+1)
// pos[eleetm]				pos[0x10] = 位置(4)		res = max(3,5-4+1)
// pos[eleetmi]				pos[0x110] = 位置(6+1)	res = 3
// pos[eleetmin]			pos[0x110] = 位置(7)		res = max(3,8-7+1)
// pos[eleetmini]			pos[0x10] = 位置(4)		res = max(3,9-4+1)
// pos[eleetminic]			pos[0x10] = 位置(4)		res = max(6,10-4+1)
// pos[eleetminico]			pos[0x1010] = 位置(11+1) res = 7
// pos[eleetminicow]		pos[0x1010] = 位置(12)	res = max(7,12-12+1)
// pos[eleetminicowo]		pos[0x10] = 位置(4)		res = max(7,13-4+1)
// pos[eleetminicowor]		pos[0x10] = 位置(4)		res = max(10,14-4+1)
// pos[eleetminicoworo]		pos[0x1010] = 位置(12)	res = max(11,15-12+1)
// pos[eleetminicoworoe]	pos[0x1000] = 位置(16+1) res = 11
// pos[eleetminicoworoep]	pos[0x1000] = 位置(17)	res = max(11,17-17+1)

func findTheLongestSubstring(s string) int {
	//pattern := "aeiou"
	const n = 1 << 5
	pos := [n]int{}
	for i := 0; i < n; i++ {
		pos[i] = -1
	}
	result := 0
	status := 0
	pos[0] = 0
	for i, v := range s {
		switch v {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] >= 0 {
			result = max(result, i-pos[status]+1)
		} else {
			//元音类型 第一次出现 (32种组合)
			pos[status] = i + 1
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//println(findTheLongestSubstring("eleetminicoworoep"))
	println(findTheLongestSubstring("el"))
}
