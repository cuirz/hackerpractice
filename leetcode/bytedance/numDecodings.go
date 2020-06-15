package main

//定义dp[i]是nums前i个字符可以得到的解码种数，假设之前的字符串是abcx，现在新加入了y，则有以下5种情况：
//如果x=='0'，且y=='0'，无法解码，返回0；
//如果只有x=='0'，则y只能单独放在最后，不能与x合并(不能以0开头)，此时有：
//dp[i] = dp[i-1]
//如果只有y=='0'，则y不能单独放置，必须与x合并，并且如果合并结果大于26，返回0，否则有：
//dp[i] = dp[i-2]
//如果 xy<=26: 则y可以“单独”放在abcx的每个解码结果之后后，并且如果abcx以x单独结尾，此时可以合并xy作为结尾，而这种解码种数就是abc的解码结果，此时有：
//dp[i+1] = dp[i] + dp[i-1]
//如果 xy>26: 此时x又不能与y合并，y只能单独放在dp[i]的每一种情况的最后，此时有：
//dp[i+1] = dp[i]

func numDecodings(s string) int {
	size := len(s)
	if size < 1 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, size+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < size; i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] < '7' { //  i 和 i-1 结合的数 1 <= x <= 26 时
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}
	return dp[size]
}

func main() {
	println(numDecodings("12"))

}
