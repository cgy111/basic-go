package funcs

// Func1没有任何参数
func Func1() {

}

// Func2有一个参数
func Func2(a int) {

}

// Func3有多个参数
func Func3(a int, b string) {

}

// Func4多个参数,一个类型
func Func4(a, b string) {

}

// Func5多个参数,一个类型,另一个写法
func Func5(a string, b string) {

}

func Func6(a, b string) string {
	return "hello，world"
}

// 多个返回值
func Func7(a, b string) (string, string) {
	return "hello,world", "go"
}

func Func8() (name string, age int) {
	return "cgy", 18
}
func Func9() (name string, age int) {
	name = "cgy"
	age = 18
	return
}
func Func10() (name string, age int) {
	//等价于 "",0
	return
}
