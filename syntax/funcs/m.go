package main

func main() {

	name, age := Func10()
	println(name, age)

	name1, _ := Func10()
	println(name1)

	//使用 := 的前提,就是左边必须有至少一个新变量
	Func6("cgy", "hello")
	name1, name2 := Func10()
	println(name1)
	println(name2)
}