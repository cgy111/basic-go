package main

func main() {
	var a int = 456
	var b int = 123

	println(a + b)
	println(a - b)
	println(a * b)

	if b != 0 {
		println(a / b)
		//取余
		println(a % b)
	}

	//println(a / 0)

	var c float64 = 12.3
	println(float64(a) + c)

	/**
	 *不兼容
	var d int32 = 12
	println(a+d)
	*/
	String()

}
