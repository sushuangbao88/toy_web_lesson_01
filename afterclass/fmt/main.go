package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	var f1 float64 = 3.1415926
	str1 := printNumWith2(f1)
	fmt.Printf("%f 保留两位小数的结果是：%s \n", f1, str1)

	fmt.Println(printBytes([]byte("你好，世界")))
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	//var str = fmt.Sprintf("%0.2f", float642)
	var str = strconv.FormatFloat(float642, 'f', 2, 64)

	return str
}

// 将[]byte 输出为16进制
func printBytes(data []byte) string {
	sign := md5.Sum(data)
	str := fmt.Sprintf("%x", sign)
	return str
}
