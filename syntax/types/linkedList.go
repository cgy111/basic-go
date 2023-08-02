package main

type LinkedList struct {
	head *node
	tail *node

	//包外可以访问
	Len int
}

func (l *LinkedList) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(idx int) (any error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) toSlise() ([]any, error) {
	//TODO implement me
	panic("implement me")
}

/*
func (l LinkedList) Add(idx int, val any) {

}

// 方法接收器，receiver
func (l *LinkedList) AddV1(idx int, val any) {

}
*/
type node struct {
	prev *node
	next *node
}
