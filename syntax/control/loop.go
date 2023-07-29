package main

import "fmt"

func ForLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	for i := 0; i < 10; {
		println(i)
		i++
	}
}
func Loop2() {
	i := 0
	for i < 10 {
		i++
		println(i)
	}
	for true {
		i++
		println(i)
	}
	for {
		i++
		println(i)
	}
}

func ForArr() {
	arr := [3]int{1, 2, 3}
	for index, val := range arr {
		println("下标 ", index, "值", val)
	}
	println()
	for index := range arr {
		println("下标 ", index, "值", arr[index])
	}
}
func ForSlice() {
	arr := []int{1, 2, 3}
	for index, val := range arr {
		println("下标 ", index, "值", val)
	}
	println()
	for index := range arr {
		println("下标 ", index, "值", arr[index])
	}
}

func ForMap() {
	m := map[string]int{
		"key1": 100,
		"key2": 102,
	}

	for k, v := range m {
		println(k, v)
	}
	println()
	for k := range m {
		println(k, m[k])
	}
}

type User struct {
	name string
}

func LoopBug() {
	users := []User{
		{
			name: "cgy",
		},
		{
			name: "Cgy",
		},
	}
	m := make(map[string]*User)
	for _, u := range users {
		fmt.Printf("%p\n", &u)
		m[u.name] = &u
	}
	fmt.Printf("%v\n", m)
}
func LoopBreak() {
	i := 0
	for true {
		if i > 10 {
			break
		}
		i++
	}
}
func LoopContinue() {
	i := 0
	for i < 10 {
		if i%2 == 1 {
			continue
		}
		println(i)
		i++
	}
}
