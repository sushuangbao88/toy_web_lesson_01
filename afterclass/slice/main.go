package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}

	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Printf("%v \n", s)

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Printf("%v \n", s)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Printf("%v \n", s)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Printf("%v \n", s)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Printf("%v \n", s)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Printf("%v \n", s)

}

func Add(s []int, index int, value int) []int {
	//当指定的位置超过的切片的容量时，矫正index未cap(s)
	if index > cap(s) {
		index = cap(s)
	}

	//根据index位置，将s分成两个字分片s1和s2
	s1 := s[0:index]
	s2 := s[index:cap(s)]
	new := make([]int, 0, cap(s)+1)

	//赋值s1部分
	for _, v := range s1 {
		new = append(new, v)
	}

	//赋值index
	new = append(new, value)

	//赋值s2部分
	for _, v := range s2 {
		new = append(new, v)
	}

	return new
}

func Delete(s []int, index int) []int {
	// 要删除的下标超出切片的容量，直接返回当前切片
	if index < 0 || index >= cap(s) {
		return s
	}

	//根据index位置，将s分成两个字分片s1和s2
	s1 := s[0:index]
	s2 := s[index+1 : cap(s)]
	new := make([]int, 0, cap(s)-1)

	//赋值s1部分
	for _, v := range s1 {
		new = append(new, v)
	}

	//赋值s2部分
	for _, v := range s2 {
		new = append(new, v)
	}

	return new
}
