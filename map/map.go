package main

import "fmt"

func main() {
	var m map[int]int
	mb := map[int]int{}
	m = make(map[int]int)

	mb[3] = 4
	fmt.Println("mb[3]:", mb[3])

	fmt.Println("Len:", len(m))

	fmt.Println(m)
	a := m[2]
	fmt.Println(a)

	m[2] = 3
	fmt.Println(a)
	a = m[2]
	fmt.Println(a)
	fmt.Println("Len:", len(m))

//any map
	mc := map[int]func (_ int) int {
		1: func (_ int) (int) { return 1; },
		2: func (_ int) (int) { return 2; },
		3: func (_ int) (int) { return 3; },
	}
	fmt.Println("Any map:", mc)
	fmt.Println("mc[1](2):", mc[1](2))

	if _, ok := mc[1]; ok {
		fmt.Println(ok)
	} else {
		fmt.Println("false")
	}

	delete(mc, 1)
	_, ok := mc[1]
	fmt.Println(ok)
}
