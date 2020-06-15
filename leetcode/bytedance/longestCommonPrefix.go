package main

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = compare(prefix, strs[i])
	}
	return prefix
}

func compare(str1, str2 string) string {
	l1 := len(str1)
	l2 := len(str2)
	l := l1
	if l1 > l2 {
		l = l2
	}

	for i := 0; i < l; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:l]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	index := 0
	prefix := strs[0]
	if len(prefix) < 1 {
		return ""
	}
	for {
		for i := 1; i < len(strs); i++ {
			if index == len(strs[i]) || prefix[index] != strs[i][index] {
				goto End
			}
		}
		index += 1
		if index == len(prefix) {
			goto End
		}
	}
End:
	return prefix[:index]
}
func main() {
	println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
}
