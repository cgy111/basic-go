package main

import "basic-go/syntax/variable"

func main() {
	global := variable.Global
	println(global)
	//variable.internal     私有变量不能调用

}
