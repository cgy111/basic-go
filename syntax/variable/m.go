package main

var Global = "全局变量"
var interna = "私有变量"

var (
	First  string = "1"
	second int    = 2
)

func main() {
	var a int = 123
	println(a)

	var b = 234
	println(b)

	var c uint = 456
	println(c)

	var (
		d string = "aaa"
		e int    = 123
	)
	println(d, e)

	f := 123
	println(f)
}
