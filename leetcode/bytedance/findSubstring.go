package main

func findSubstring(s string, words []string) []int {
	dic := make(map[string]int)
	n := len(words)
	if n < 1{
		return []int{}
	}
	size := len(words[0])
	for i := 0; i < n; i++ {
		dic[words[i]] += 1
	}

	result := make([]int, 0)
	for i := 0; i < len(s); i++ {
		j := i
		sum := 0
		subdic := make(map[string]int)
		for j <= len(s) {
			w := s[j:j+size]
			if dic[w] > 0 && subdic[w] < dic[w]{
				subdic[s[j:j+size]] += 1
				sum += 1
				j += size
			}else{
				break
			}
		}
		if sum == n{
			result = append(result,i)
		}
	}
	return result

}
