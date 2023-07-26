package main

const External = "包外"
const internal = "包内"
const (
	a = 123
)

const (
	Init = iota
	Running
	Paused
	Stop

	Status = 100
)

const (
	DayA = iota*12 + 13
	DayB
)

const (
	NumA = iota << 1
	NumB
	NumC
)

func main() {
	const a = 123

	//	常量不需改变
	//a=456
}
