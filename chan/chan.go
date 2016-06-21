package main

import "fmt"
import "time"

func main() {
	var ch chan int
	ch = make(chan int)

	go send(ch)
	go get(ch)

	time.Sleep(50000000)
}


func send(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("send to chan", i)
		ch <- i
		time.Sleep(5000)
	}
}

func get(ch chan int) {
	for {
		fmt.Println("Get data from chan")
		fmt.Println(<- ch)
		time.Sleep(50)
	}
}
