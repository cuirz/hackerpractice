package main

//729. 我的日程安排表 I
//实现一个 MyCalendar 类来存放你的日程安排。如果要添加的日程安排不会造成 重复预订 ，则可以存储这个新的日程安排。
//
//当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生 重复预订 。
//
//日程可以用一对整数 start 和 end 表示，这里的时间是半开区间，即 [start, end), 实数x 的范围为， start <= x < end 。
//
//实现 MyCalendar 类：
//
//MyCalendar() 初始化日历对象。
//boolean book(int start, int end) 如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true 。否则，返回 false并且不要将该日程安排添加到日历中。
//
//
//示例：
//
//输入：
//["MyCalendar", "book", "book", "book"]
//[[], [10, 20], [15, 25], [20, 30]]
//输出：
//[null, true, false, true]
//
//解释：
//MyCalendar myCalendar = new MyCalendar();
//myCalendar.book(10, 20); // return True
//myCalendar.book(15, 25); // return False ，这个日程安排不能添加到日历中，因为时间 15 已经被另一个日程安排预订了。
//myCalendar.book(20, 30); // return True ，这个日程安排可以添加到日历中，因为第一个日程安排预订的每个时间都小于 20 ，且不包含时间 20 。
//
//
//提示：
//
//0 <= start < end <= 10^9
//每个测试用例，调用 book 方法的次数最多不超过 1000 次。

// 二分查找，红黑树
// 线段树

type MyCalendar struct {
	tree, lazy map[int]bool
}

func Constructor() MyCalendar {
	return MyCalendar{make(map[int]bool), make(map[int]bool)}

}
func (this MyCalendar) query(s, e, l, r, index int) bool {
	if s > r || e < l {
		return false
	}
	if this.lazy[index] {
		return true
	}
	if s <= l && r <= e {
		return this.tree[index]
	}
	mid := (l + r) >> 1
	return this.query(s, e, l, mid, index*2) || this.query(s, e, mid+1, r, index*2+1)

}

func (this MyCalendar) update(s, e, l, r, index int) {
	if s > r || e < l {
		return
	}
	if s <= l && r <= e {
		this.tree[index] = true
		this.lazy[index] = true
	} else {
		mid := (l + r) >> 1
		this.update(s, e, l, mid, index*2)
		this.update(s, e, mid+1, r, index*2+1)
		this.tree[index] = true
		if this.lazy[index*2] && this.lazy[index*2+1] {
			this.lazy[index] = true
		}
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.query(start, end-1, 0, 1e9, 1) {
		return false
	}
	this.update(start, end-1, 0, 1e9, 1)
	return true

}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
