package main

func numTrees2(n int) int {
	return tree(0, n-1)
}

func tree(s, e int) int {
	if s > e {
		return 1
	}

	result := 0
	for i := s; i < e; i++ {
		l := tree(s, i-1)
		r := tree(i+1, e)
		result += l * r
	}
	return result
}

func numTrees(n int) int {
	//int fun[] = new int[n + 1];
	sum := make([]int, n+1)
	sum[0] = 1
	sum[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			sum[i] += sum[j] * sum[i-1-j]
		}
	}
	return sum[n]
}

func main() {

}
