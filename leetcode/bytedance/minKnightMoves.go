package main

func minKnightMoves(x int, y int) int {
	dp := make(map[pos]int)
	dp[pos{x: 0, y: 0}] = 0
	return findPath(dp, x, y, 0)

}

type pos struct {
	x int
	y int
}

var move = [][]int{{1, 2}, {2, 1}, {1, -2}, {-1, 2}, {-1, -2}, {2, -1}, {-2, 1}, {-2, -1}}

func findPath(path map[pos]int, x, y, count int) int {
	dp := make(map[pos]int)
	for k, _ := range path {
		for _, step := range move {
			_x := k.x + step[0]
			_y := k.y + step[1]
			p := pos{x: _x, y: _y}
			if !checkEdge(p, x, y) {
				continue
			}
			dp[p] = 1
			if _x == x && _y == y {
				count++
				return count
			}
		}
	}
	count++
	return findPath(dp, x, y, count)
}

func checkEdge(p pos, x, y int) bool {
	l, r, t, b := 0, 0, 0, 0
	if x > 0 {
		l = 0 - 1
		r = x + 1
	} else {
		l = x - 1
		r = 0 + 1
	}
	if y > 0 {
		t = y + 1
		b = 0 - 1
	} else {
		t = 0 + 1
		b = y - 1
	}

	if l <= p.x && p.x <= r && b <= p.y && p.y <= t {
		return true
	}
	return false
}

func main() {
	println(minKnightMoves(-2, -2))

}
