package main

func ForLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	for i := 0; i < 10; {
		println(i)
		i++
	}
}
func Loop2() {
	i := 0
	for i < 10 {
		i++
		println(i)
	}
	for true {
		i++
		println(i)
	}
	for {
		i++
		println(i)
	}
}

func ForArr() {
	arr := [3]int{1, 2, 3}
	for index, val := range arr {
		println("下标 ", index, "值", val)
	}
	println()
	for index := range arr {
		println("下标 ", index, "值", arr[index])
	}
}
