package main

import "fmt"

func main() {
	x := t{2, 3}

	fmt.Println(x)
	fmt.Println(x.Add())
	fmt.Println(x)
}

type t struct {
	a int
	b int
}

func (a t) Add() (int) {
	tmp := a.a + a.b;
	a.a = 10
	a.b = 10

	return tmp
}
