package main

import "fmt"

func Cloursre(name string) func() string {

	return func() string {
		return "hello," + name
	}

}
func Cloursre2() func() string {
	name := "cgy"
	age := 18
	return func() string {
		return fmt.Sprintf("Hello,%s,%d", name, age)
	}

}
func Cloursre3() func() int {
	var age = 0
	return func() int {
		age++
		return age
	}

}
