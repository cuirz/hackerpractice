package main

type MyCalendarTwo struct {
	array   []book
	overlap []book
}

type book struct {
	s, e int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.overlap {
		if v.s < end && v.e > start {
			return false
		}
	}

	for _, v := range this.array {
		if v.s < end && v.e > start {
			this.overlap = append(this.overlap, book{s: max(start, v.s), e: min(end, v.e)})
		}
	}

	this.array = append(this.array, book{s: start, e: end})
	return true
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
