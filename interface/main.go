package main

import "fmt"
import "reflect"

func main() {
	var obj Person = Person{}
	obj.s = "Someone"
	var i I = obj

	fmt.Println(i.(Person))

//empty interface
	var empty Any
	fmt.Println(empty)

	empty = obj.s
	fmt.Println(empty)

	i.SayHello()
	obj.SayHello()

	TestReflect()
}

func TestReflect() {
	var f float64 = 3.14159

	v := reflect.ValueOf(f)
	fmt.Println("Type:", reflect.TypeOf(f))
	fmt.Println("Value:", reflect.ValueOf(f))
	fmt.Println("Kind:", v.Kind())
}

type Any interface{}

type I interface {
	SayHello() (string)
}

type Person struct {
	s string
}

func (obj Person) SayHello() (string) {
	fmt.Println("Hello:", obj.s)

	return obj.s
}
