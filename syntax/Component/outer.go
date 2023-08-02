package main

type Inner struct {
}

func (i Inner) DoSomething() {
	println("这是Inner")
}

func (i Inner) Name() string {
	return "inner"
}

func (i Inner) sayHello() {
	println("hello", i.Name())
}

type Outer struct {
	Inner
}

func (Outer) Name() string {
	return "outer"
}

type OuterV1 struct {
	Inner
}

func (o OuterV1) DoSomething() {
	println("这是outer")
}

type OuterPtr struct {
	*Inner
}

type OOOuter struct {
	Outer
}

func UseInner() {
	var o Outer
	o.DoSomething()

	var op *OuterPtr
	op.DoSomething()

	o1 := Outer{
		Inner: Inner{},
	}
	op1 := OuterPtr{
		Inner: &Inner{},
	}
	o1.DoSomething()
	op1.DoSomething()
}

func main() {
	var o1 OuterV1
	o1.DoSomething()
	o1.Inner.DoSomething()

	var o Outer
	o.sayHello()
}
