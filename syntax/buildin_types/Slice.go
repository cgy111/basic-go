package main

import "fmt"

func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("a1:%v,len=%d,cap=%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4)
	fmt.Printf("a2:%v,len=%d,cap=%d", s2, len(s2), cap(s2))
}
