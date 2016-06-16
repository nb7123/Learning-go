package main

//#include<stdlib.h>
import "C"

import (
	"fmt"
)

//import "os"
//import "strings"
//import "strconv"
//import "time"

type Rope string

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(arr1[:]))

	for i, v := range arr1 {
		fmt.Println(i, v)
	}
}

func sum (a []int) (int) {
	sum := 0
	for _, v := range a {
		sum += v
		fmt.Println(v)
	}

	a[2] = 18
	return sum
}

func callback(a, b, c int, f func(int, int)(int)) int {
	return a * f(b, c)
}

func add(a, b int) int {
	return a + b
}

func returnMulti () (a int, b int) {
	return 3, 4
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
