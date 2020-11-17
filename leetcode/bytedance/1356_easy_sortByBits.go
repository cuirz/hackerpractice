package main

import "sort"

//1356. 根据数字二进制下 1 的数目排序
//给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。
//
//如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。
//
//请你返回排序后的数组。

func sortByBits(arr []int) []int {
	dic := make([][2]int,len(arr))
	for i,v := range arr{
		count := 0
		val := v
		for v > 0{
			count ++
			v = v & (v-1)
		}
		dic[i][0] = val
		dic[i][1] = count
	}
	sort.Slice(dic,func(i,j int)bool{
		if dic[i][1] < dic[j][1] || (dic[i][1] == dic[j][1] && dic[i][0] < dic[j][0]){
			return true
		}
		return false
	})
	for i,_ := range arr{
		arr[i] = dic[i][0]
	}
	return arr
}


func main(){
	println(sortByBits([]int{0,1,2,3,4,5,6,7,8}))
}