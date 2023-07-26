package main

func Cloursre(name string) func() string {

	return func() string {
		return "hello," + name
	}
}
