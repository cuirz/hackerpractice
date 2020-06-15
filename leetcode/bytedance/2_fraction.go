package main

func fraction(cont []int) []int {
	n := len(cont)
	// D 分母
	// M 分子
	D, M := 0, 0
	M = cont[n-1]
	D = 1
	temp := 0
	for i := n - 2; i > -1; i-- {
		D = cont[i]*M + D
		temp = D
		D = M
		M = temp
	}
	return []int{M, D}
}

func main() {

}
