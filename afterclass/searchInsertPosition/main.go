package main

import "fmt"

/*
35. 搜索插入位置：
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
*/
func main() {
	res1 := searchInsert([]int{1, 3, 5, 6}, 5)
	fmt.Printf("%v \n", res1)

	res2 := searchInsert([]int{1, 3, 5, 6}, 2)
	fmt.Printf("%v \n", res2)

	res3 := searchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Printf("%v \n", res3)

}

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0 //边界情况：在头部找到，或者target 小于数组所有的值
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1 //边界情况：在尾部找到
	}
	if target > nums[len(nums)-1] {
		return len(nums) //边界情况：target大于数组的所有值
	}

	//剩下的情况，target在nums的最小和最大值之间
	start := 0
	end := len(nums) - 1
	for {
		mid := (start + end) / 2
		if mid == start {
			return start + 1 //取中间下标的时候，等于start，说明：找不到targte，且此时start和end是相邻的整数。
		}

		midVal := nums[mid]
		if target > midVal {
			start = mid //缩小查找范围：start后移
		} else if target < midVal {
			end = mid //所辖查找范围：end后移
		} else {
			return mid //中间下标正中目标
		}
	}
}
