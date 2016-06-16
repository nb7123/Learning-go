package main

import "fmt"

func main() {
	s := new(Stack)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)

	s.Push(2)
	s.Push(8)
	s.Push(10)
	s.Push(3)
	s.Push(15)
	fmt.Println(s)

	fmt.Println(s.Pop())
	fmt.Println(s)
}

type Stack struct {
	data [4]int
	cur int
}

//func (stack *Stack) String() string {
//	return ""
//}

func (stack *Stack) Pop() (int) {
	if stack.cur == 0 {
		return 0
	}
	result := stack.data[stack.cur-1]
	stack.data[stack.cur-1] = 0
	stack.cur--

	return result
}

func (stack *Stack) Push(a int) {
	if stack.cur == 4 {
		return
	}

	stack.data[stack.cur] = a
	stack.cur++
}


