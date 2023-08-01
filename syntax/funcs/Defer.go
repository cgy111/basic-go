package main

import "fmt"

func Defer() {
	defer func() {
		println("第一个defer")
	}()
	defer func() {
		println("第二个defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}
func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

/*func Query()  {
	db := sql.Open("","")
	defer db.Close()
	db.Query("SELECT")
}*/

func DerferRetun() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("i的地址是%p,值是%d\n", &i, i)
			//println(i)
		}()
	}
}
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("val的地址是%p,值是%d\n", &val, val)
			//println(i)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("j的地址是%p,值是%d\n", &j, j)
			//println(i)
		}()
	}
}
