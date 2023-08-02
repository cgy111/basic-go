package main

type Interage int

func UseInt() {
	i1 := 10
	i2 := Interage(i1)
	var i3 Interage = 11
	println(i2, i3)
}

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	println("fist 在游")
}

type FakeFish Fish

func UseFish() {
	f1 := Fish{}
	f2 := FakeFish(f1)
	//f2.Swim()
	println(f2)
}

/*
type MyTime time.Time

func (m MyTime) MyFunc() {

}*/
//向后兼容
type Yu = Fish
