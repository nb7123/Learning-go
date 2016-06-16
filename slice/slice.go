package main

import "fmt"
import "sort"

func main() {
	s := make([]int, 2, 10)

	fmt.Println("Len:", len(s), "Cap:", cap(s))

//	s[3] = 10

	sb := make([]byte, 2, 10)
	fmt.Println("SB Len:", len(sb), "SB Cap:", cap(sb))

	fmt.Println(s)
	s = s[2:3]
	fmt.Println(s)
//Sort
	arr1 := []int{3, 9, 2, 7, 9, 7, 10}
	s1 := arr1[:]

	fmt.Println(sort.IntsAreSorted(arr1))
	
	fmt.Println(arr1)
	fmt.Println(s1)

	sort.Ints(s1)
	fmt.Println(s1)
	fmt.Println(arr1)
}
