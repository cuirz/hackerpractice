package main
//391. 完美矩形
//给你一个数组 rectangles ，其中 rectangles[i] = [xi, yi, ai, bi] 表示一个坐标轴平行的矩形。这个矩形的左下顶点是 (xi, yi) ，右上顶点是 (ai, bi) 。
//
//如果所有矩形一起精确覆盖了某个矩形区域，则返回 true ；否则，返回 false 。
//
// 
//示例 1：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
//输出：true
//解释：5 个矩形一起可以精确地覆盖一个矩形区域。
//示例 2：
//
//
//输入：rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
//输出：false
//解释：两个矩形之间有间隔，无法覆盖成一个矩形。
//示例 3：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]
//输出：false
//解释：图形顶端留有空缺，无法覆盖成一个矩形。
//示例 4：
//
//
//输入：rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
//输出：false
//解释：因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。
// 
//
//提示：
//
//1 <= rectangles.length <= 2 * 10^4
//rectangles[i].length == 4
//-10^5 <= xi, yi, ai, bi <= 10^5
func isRectangleCover(rectangles [][]int) bool {
	type point struct {
		x,y int
	}
	area,minx,miny,maxx,maxy := 0,rectangles[0][0],rectangles[0][1],rectangles[0][2],rectangles[0][3]
	cnt := make(map[point]int)
	for _,v := range rectangles{
		x,y,a,b := v[0],v[1],v[2],v[3]
		area += (a-x)*(b-y)
		minx = min(minx,x)
		miny = min(miny,y)
		maxx = max(maxx,a)
		maxy = max(maxy,b)

		cnt[point{x,y}]++
		cnt[point{a,b}]++
		cnt[point{x,b}]++
		cnt[point{a,y}] ++
	}
	if area != (maxx-minx)*(maxy-miny) || cnt[point{minx,miny}] != 1 || cnt[point{minx,maxy}] != 1 ||
		cnt[point{maxx,miny}] != 1 || cnt[point{maxx,maxy}] != 1{
		return false
	}
	delete(cnt,point{minx,miny})
	delete(cnt,point{minx,maxy})
	delete(cnt,point{maxx,maxy})
	delete(cnt,point{maxx,miny})

	for _,c := range cnt{
		if c != 2 && c != 4{
			return false

		}
	}
	return true
}


