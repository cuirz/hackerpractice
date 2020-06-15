package main

import "strings"

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
	v []int
}

func (this *MountainArray) get(index int) int {
	return this.v[index]
}
func (this *MountainArray) length() int {
	return len(this.v)
}
func findInMountainArray(target int, mountainArr *MountainArray) int {
	peek := 0
	s := 0
	e := mountainArr.length() - 1
	m := 0
	mid := 0
	for s < e {
		peek = (s + e + 1) / 2
		if mountainArr.get(peek-1) < mountainArr.get(peek) {
			s = peek
		} else {
			e = peek - 1
		}
	}
	peek = s

	//左查找
	s = 0
	e = peek
	for s < e {
		m = (s + e) / 2
		mid = mountainArr.get(m)
		if mid == target {
			return m
		} else if mid > target {
			e = m
		} else {
			s = m + 1
		}
	}
	if mountainArr.get(e) == target {
		return e
	}
	//右查找
	s = peek
	e = mountainArr.length() - 1
	for s < e {
		m = (s + e) / 2
		mid = mountainArr.get(m)
		if mid == target {
			return m
		} else if mid > target {
			s = m + 1
		} else {
			e = m
		}
	}
	if mountainArr.get(s) == target {
		return s
	}

	return -1

}

func simplifyPath(path string) string {
	ps := strings.Split(path, "/")
	var array []string
	for _, v := range ps {
		switch v {
		case ".", "/", "":
		case "..":
			if len(array) > 0 {
				array = array[:len(array)-1]
			}
		default:
			array = append(array, v)
		}
	}

	return "/" + strings.Join(array, "/")

}

func main() {

	println(findInMountainArray(2, &MountainArray{
		v: []int{1, 5, 2},
	}))

	println(simplifyPath("//home/"))
}
