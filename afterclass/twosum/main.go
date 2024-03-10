package main

import "fmt"

/*
两数之和:
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
*/
func main() {
	nums1 := []int{2, 7, 11, 15}
	var target1 int = 9
	res1 := twoSum(nums1, target1)
	fmt.Printf("%v \n", res1)

	nums2 := []int{3, 2, 4}
	var target2 int = 6
	res2 := twoSum(nums2, target2)
	fmt.Printf("%v \n", res2)

	nums3 := []int{3, 3}
	var target3 int = 6
	res3 := twoSum(nums3, target3)
	fmt.Printf("%v \n", res3)
}

func twoSum(nums []int, target int) []int {
	arr := make(map[int]int)

	for idx, val := range nums {
		another := target - val

		index, ok := arr[another]
		if ok { //在map中找到匹配到值，并找到找到记录的原数组的下标
			return []int{index, idx}
		}

		arr[val] = idx //将数组中元素，记录到map中，
	}

	return []int{}
}
