package main

import "sort"

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}
	//var array svalue = intervals
	//sort.Sort(array)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	max := 1
	count := 1
	for i, v := range intervals {
		count = 1
		for k := i + 1; k < len(intervals); k++ {
			if v[1] > intervals[k][0] {
				count++
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}

type svalue [][]int

func (p svalue) Len() int {
	return len(p)
}
func (p svalue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]

}
func (p svalue) Less(i, j int) bool {
	if p[i][1] < p[j][1] {
		return true
	}
	return false
}
