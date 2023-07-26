package main

func Functional4() {
	println("hello,functional4")
}
func Functional5(age int) {

}

var Abc = func() string {
	return "hello"
}

func UseFunctional4() {
	myFun := Functional4
	myFun()

	fun5 := Functional5
	fun5(18)
}

func Functional6() {
	//新定义了一个方法，赋值给了fun
	fn := func() string {
		return "hello"
	}
	fn()
}

// 匿名方法立即发起调用
func Functional8() {
	//新定义了一个方法，赋值给了fun
	fn := func() string {
		return "hello"
	}()
	println(fn)
}

// 他的意思是我返回一个，返回string的无参方法
func Functional7() func() string {
	return func() string {
		return "hello"
	}
}
