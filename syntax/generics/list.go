package main

// T类型参数，名字叫做T，约束是any，等于没有约束
type List[T any] interface {
	Add(idx int, t T)
	Append(t T)
}

func UseList() {
	var l List[int]
	l.Append(1)

	//var f List[float64]
	//f.Append(1.1)

}

type LinkedList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	val T
}

func main() {
	UseList()
}
