package main

const External = "包外"
const internal = "包内"
const (
	a = 123
)

const (
	StatusA = iota
	StatusB
	StatusC
	StatusD

	Status = 100
)

const (
	DayA = iota*12 + 13
	DayB
)

func main() {
	const a = 123

	//	常量不需改变
	//a=456
}
