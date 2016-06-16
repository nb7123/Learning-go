package main

import "fmt"

func main() {
	m := Method{3, 4, A{5, 6}}
	fmt.Println(m)

	m.PrintA()
	m.PrintB()

	fmt.Println(m.a, m.b, m.c)

	i := new(Integer)

	fmt.Println(i.get())
}

type Integer int

func (p Integer) get() int {
	return int(p)
}

type Method struct {
	a int
	b int

	A
}

type A struct {
	a int
	c float32
}

func (a *A) PrintA() {
	fmt.Println(a.a)
}

func (a *A) PrintC() {
	fmt.Println(a.c)
}

func (m*Method) PrintB() {
	fmt.Println("B")
}
