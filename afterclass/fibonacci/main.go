package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(`实现计算“斐波那契数列”`)
	fmt.Println("请输入计算方式（f1/f2）以及位数n：")
	var way string
	var n int
	fmt.Scanf("%s %d", &way, &n)

	start := time.Now()
	var val int
	if way == "f1" {
		val = f1(n)
	} else if way == "f2" {
		val = f2(n)
	} else {
		fmt.Println("未输入正确计算方式")
	}
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("斐波那契数列,使用方式：%s ,第 %d 是 %d,耗时：%s", way, n, val, elapsed)
}

// 计算斐波那契数列:1,1,2,3,5,8,13...
// 方式一：递归
func f1(n int) int {
	// TODO
	if n < 1 {
		return 0
	} else if n <= 2 {
		return 1
	} else {
		return f1(n-1) + f1(n-2)
	}
}

// 方式2：顺序累加
func f2(n int) int {
	// TODO
	if n <= 0 {
		return 0
	} else if n <= 2 {
		return 1
	}

	var p1, p2 int = 1, 1 //定义初始的两个值
	var i int = 3
	var val int
	for i <= n {
		val = p1 + p2
		p1 = p2
		p2 = val
		i++
	}
	return val
}
