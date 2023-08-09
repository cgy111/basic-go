package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	fmt.Println(a, len(a), cap(a))

	a = append([]int{0}, a...)
	fmt.Println(a, len(a), cap(a))

	/*a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a, len(a), cap(a))*/

	a = append([]int{-3, -2}, a...)
	fmt.Println(a, len(a), cap(a))

	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a, len(a), cap(a))

	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a, len(a), cap(a))

	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a, len(a), cap(a))

	a = append([]int{1, 2, 3}, a...)
	fmt.Println(a, len(a), cap(a))
}
