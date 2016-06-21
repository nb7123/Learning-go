package main

import "fmt"
import "encoding/json"

type A struct {
	a string
	b string
	s string
}

func main() {
	x := A{"a", "b", "abc"}

	fmt.Println(x)

	js, er := json.Marshal(x)
	if nil != er {
		fmt.Println(er)
		return
	}
	fmt.Println("Endoded:", js)
}
