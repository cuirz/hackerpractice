package main

func gardenNoAdj(N int, paths [][]int) []int {
	n := len(paths)
	graph := make(map[int][]int, 0)
	for i := 0; i < n; i++ {
		f1 := paths[i][0] - 1
		f2 := paths[i][1] - 1
		graph[f1] = append(graph[f1], f2)
		graph[f2] = append(graph[f2], f1)
	}
	result := make([]int, N)
	// 题中 没有3条以上路径，所以可涂色的最大值是4，包含自己
	for i := 0; i < N; i++ {
		v := graph[i]
		flower := make([]bool, 5)
		for _, g := range v {
			flower[result[g]] = true
		}
		for k := 1; k <= 4; k++ {
			if !flower[k] {
				result[i] = k
				break
			}
		}
	}

	return result

}

func main() {

}
