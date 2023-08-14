package main

import "fmt"

// 二分法基础版
/*
Params:a-待查找的升序数组
       target-待查找额目标值

Returns：
       找到则返回索引
       找不到则返回1
*/
func binarySearchBasic(a []int, target int) int {

	//设置指针和初值
	var i = 0
	var j = len(a) - 1
	for i <= j { //范围内有东西
		m := (i + j) / 2
		if target < a[m] { //目标在中间值的左边
			j = m - 1
		} else if a[m] < target { //目标在右边
			i = m + 1
		} else {
			return m

		}
	}
	return -1
}

func main() {
	a := []int{7, 13, 21, 30, 38, 44, 52, 53}
	target := 52
	index := binarySearchBasic(a, target)
	fmt.Println(index)
	//println(m)
}
