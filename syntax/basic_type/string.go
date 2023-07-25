package main

import "fmt"

func String() {
	//He said:"Hello Go!"
	println("He said:\"Hello Go!\"")
	println("Hello,\"go!")
	println(`
可以换行
再一行
`)

	println("Hello" + "go")
	//println("hello"+string(123))
	fmt.Printf("hello %d", 123)
}
