package main

//17.11. 单词距离
//有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同
//
//示例：
//
//输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
//输出：1
//提示：
//
//words.length <= 100000

func findClosest(words []string, word1 string, word2 string) int {
	index1, index2, result := -1, -1, len(words)
	for i, w := range words {
		if w == word1 {
			index1 = i
		} else if w == word2 {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 {
			result = min(result, abs(index1-index2))
		}
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	println(findClosest([]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}, "a", "student"))
}
