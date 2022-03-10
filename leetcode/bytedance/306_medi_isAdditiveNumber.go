package main

//306. 累加数
//累加数 是一个字符串，组成它的数字可以形成累加序列。
//
//一个有效的 累加序列 必须 至少 包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
//
//给你一个只包含数字'0'-'9'的字符串，编写一个算法来判断给定输入是否是 累加数 。如果是，返回 true ；否则，返回 false 。
//
//说明：累加序列里的数 不会 以 0 开头，所以不会出现1, 2, 03 或者1, 02, 3的情况。
//
//
//
//示例 1：
//
//输入："112358"
//输出：true
//解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
//示例2：
//
//输入："199100199"
//输出：true
//解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
//
//
//提示：
//
//1 <= num.length <= 35
//num 仅由数字（0 - 9）组成
//
//
//进阶：你计划如何处理由过大的整数输入导致的溢出?

func isAdditiveNumber(num string) bool {
	n := len(num)
	add := func(a, b string) string {
		var res []byte
		var carry, cur int
		for a != "" || b != "" || carry != 0 {
			cur = carry
			if a != "" {
				cur += int(a[len(a)-1] - '0')
				a = a[:len(a)-1]
			}
			if b != "" {
				cur += int(b[len(b)-1] - '0')
				b = b[:len(b)-1]
			}
			carry = cur / 10
			cur %= 10
			res = append(res, byte(cur)+'0')
		}
		for i, m := 0, len(res); i < m/2; i++ {
			res[i], res[m-1-i] = res[m-1-i], res[i]
		}
		return string(res)

	}

	isValid := func(s, e int) bool {
		fs, fe := 0, s-1
		for e <= n-1 {
			third := add(num[fs:fe+1], num[s:e+1])
			ts := e + 1
			te := e + len(third)
			if te >= n || num[ts:te+1] != third {
				break
			}
			if te == n-1 {
				return true
			}
			fs, fe = s, e
			s, e = ts, te
		}
		return false
	}
	for s := 1; s < n-1; s++ {
		if num[0] == '0' && s != 1 {
			break
		}
		for e := s; e < n-1; e++ {
			if num[s] == '0' && s != e {
				break
			}
			if isValid(s, e) {
				return true
			}

		}
	}
	return false
}

func main() {
	println(isAdditiveNumber("112358"))
}
