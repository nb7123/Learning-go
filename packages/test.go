package main

import "fmt"
import "unsafe"

import "../pack/pack"

func main() {
	fmt.Println(unsafe.Sizeof(string{}))
	fmt.Println(pack.A)
}
