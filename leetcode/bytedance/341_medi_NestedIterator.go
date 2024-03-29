package main

//341. 扁平化嵌套列表迭代器
//给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
//
//列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。
//
//
//
//示例 1:
//
//输入: [[1,1],2,[1,1]]
//输出: [1,1,2,1,1]
//解释: 通过重复调用next 直到hasNext 返回 false，next返回的元素的顺序应该是: [1,1,2,1,1]。
//示例 2:
//
//输入: [1,[4,[6]]]
//输出: [1,4,6]
//解释: 通过重复调用next直到hasNext 返回 false，next返回的元素的顺序应该是: [1,4,6]。

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {}

type NestedIterator struct {
	array []int
	index int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var list []int
	var flatList func(nested []*NestedInteger)
	flatList = func(nested []*NestedInteger) {
		for _, v := range nested {
			if v.IsInteger() {
				list = append(list, v.GetInteger())
			} else {
				flatList(v.GetList())
			}
		}
	}
	flatList(nestedList)
	return &NestedIterator{list, 0}
}

func (this *NestedIterator) Next() int {
	i := this.index
	this.index++
	return this.array[i]
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.array)
}
