package main

import "fmt"
import "bytes"

func main() {
	bs := new(bytes.Buffer)

	ss := []string{"A", "b", "C", "D"}

	for _, v := range ss {
		bs.WriteString(v)
	}

	fmt.Println(bs.String())
}
