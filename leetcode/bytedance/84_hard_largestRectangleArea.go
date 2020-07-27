package main

//84. 柱状图中最大的矩形

func largestRectangleArea(heights []int) int {
	n := len(heights)
	result := 0
	for i := 0; i < n; i++ {
		h := heights[i]
		for j := i; j < n; j++ {
			h = min(h, heights[j])
			result = max(result, (j-i+1)*h)
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
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
