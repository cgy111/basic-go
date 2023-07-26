package main

func Recursive(n int) {
	if n > 10 {
		return
	}

	Recursive(n + 1)
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	A()
}
