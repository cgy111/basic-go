package main

import "fmt"

func main() {
	/*var x = 10
	fmt.Printf("x的地址%p\n", &x)

	y := &x;
	fmt.Printf("y的地址%p\n", &y)

	x = 100
	fmt.Println(*y)
	*/
	/*	var a = []int{1, 2, 3}
		fmt.Printf("a的地址%p\n", &a)
		var b = &a
		fmt.Printf("b的地址%p\n", &b)
		a[0] = 100
		fmt.Println(*b)*/

	var c = []int{1, 2, 3}
	fmt.Printf("c的地址%p\n", &c)
	d := c
	fmt.Printf("b的地址%p\n", &d)
	c[0] = 100
	fmt.Println(d)

}
