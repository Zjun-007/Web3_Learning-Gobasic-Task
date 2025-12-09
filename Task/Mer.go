// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

package main

import (
	"fmt"
	"sort"
)	
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}		
	// 二维切片按照区间起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})	
	merged := [][]int{intervals[0]}  //初始化第一个区间
	for i := 1; i < len(intervals); i++ {
		// 获取当前区间和合并后区间的最后一个区间	
		current := intervals[i]
		lastMerged := merged[len(merged)-1]	
		// 检查是否有重叠	
		if current[0] <= lastMerged[1] {
			// 有重叠，合并区间
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}		
		} else {
			// 无重叠，添加当前区间到合并后的区间列表
			merged = append(merged, current)
		}
	}
	return merged
}	
func main() {		
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},	
		{{1, 4}, {0, 2}, {3, 5}},
	}		
	for _, intervals := range testCases {
		fmt.Printf("输入: %v\n", intervals)
		result := merge(intervals)
		fmt.Printf("输出: %v\n", result)
	}	
}