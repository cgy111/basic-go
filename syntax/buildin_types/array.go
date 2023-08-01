package main

import "fmt"

func Array() {
	a1 := [3]int{9, 8, 7}
	fmt.Printf("a1:%v,len=%d,cap=%d 、\n", a1, len(a1), cap(a1))

	a2 := [3]int{9, 8}
	fmt.Printf("a2:%v,len=%d,cap=%d 、\n", a2, len(a2), cap(a2))

	a3 := [3]int{}
	fmt.Printf("a3:%v,len=%d,cap=%d 、\n", a3, len(a3), cap(a3))
}

func UseSumInt64() {

	s1 := []int{1, 2, 3}
	res := SumInt(s1)
	println(res)

	s2 := make([]int64, 0, len(s1))

	for _, v := range s1 {
		s2 = append(s2, int64((v)))
	}
}
func SumInt(vals []int) int {
	var res int
	for _, val := range vals {
		res += val
	}
	return res
}

func SumInt64(vals []int64) int64 {
	var res int64
	for _, val := range vals {
		res += val
	}
	return res
}
