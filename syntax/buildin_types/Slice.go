package main

import (
	"fmt"
)

func Slice() {
	s1 := []int{9, 8, 7}
	fmt.Printf("s1:%v,len=%d,cap=%d \n", s1, len(s1), cap(s1))
	/*i := s1[0]
	fmt.Printf(strconv.Itoa(i))*/

	s2 := make([]int, 3, 4)
	fmt.Printf("s2:%v,len=%d,cap=%d\n", s2, len(s2), cap(s2))

	//{0,0,0,0}
	s3 := make([]int, 4)
	//make([]int,0,4)
	fmt.Printf("s3:%v,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	s4 := make([]int, 0, 4)
	s4 = append(s4, 1)
	fmt.Printf("s4:%v,len=%d,cap=%d", s4, len(s4), cap(s4))
}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3]
	fmt.Printf("s2:%v,len=%d,cap=%d\n", s2, len(s2), cap(s2))

	s3 := s1[2:]
	fmt.Printf("s3:%v,len=%d,cap=%d\n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4:%v,len=%d,cap=%d", s4, len(s4), cap(s4))
}
