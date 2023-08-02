package main

import "fmt"

func NewUser() {

	u := User{}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)

	//up是一个指针
	up := &User{}
	fmt.Printf("up: %+v\n", up)
	up2 := new(User)
	println(up2.FirstName)
	fmt.Printf("up2: %+v\n", up2)

	u4 := User{Name: "Tom", Age: 0}
	//不推荐使用
	u5 := User{"Tom", "Tony", 0}

	u4.Name = "Mike"
	u5.Age = 18

	var up3 *User = nil
	//nil 上访问字段，或者方法
	println(up3.FirstName)
	println(up3)

}

type User struct {
	Name      string
	FirstName string
	Age       int
}

func (u User) ChangeName(name string) {
	fmt.Printf("Change name中u的地址%p\n", &u)
	u.Name = name
}
func ChangeName(u User, name string) {

}
func (u *User) ChangeAge(age int) {
	fmt.Printf("Change age中u的地址%p\n", u)
	u.Age = age
}
func ChangeAge(u *User, age int) {

}

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	fmt.Printf("u1 的地址%p\n", &u1)
	//(&u1).ChangeAge(35)
	u1.ChangeAge(35)
	//这一步执行的时候，其实相当于复制了一个U1,改的是复制体
	//所以u1原封不动
	u1.ChangeName("Jerry")
	fmt.Printf("%+v", u1)

	up1 := &User{}
	//(*up1).ChangeName("Jerry")
	up1.ChangeName("Jerry")

	up1.ChangeAge(35)
	fmt.Printf("%+v", up1)
}
