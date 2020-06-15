package main

//求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

func sumNums(n int) int {
	result := 0
	var sum func(x int) bool
	sum = func(x int) bool {
		result += x
		return x > 0 && sum(x-1)
	}
	sum(n)
	return result
}

func main() {
	println(sumNums(6))
}
