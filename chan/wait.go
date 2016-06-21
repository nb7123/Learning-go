package main

import "fmt"
import "time"

func main() {
	fmt.Println("Start...")
	ch := make(chan int)
	sem := make(chan int)

	go func(ch chan int, sem chan int){
		for i := 0; i < 10; i++ {
			ch <- i * 10
			time.Sleep(time.Second * 1)
		}
		sem <- 1
	} (ch, sem)

	go func(ch chan int, sem chan int) {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

		time.Sleep(time.Second * 3)
		sem <- 2
	}(ch, sem)
	
	<- sem
	<- sem
	fmt.Println("End")
}

func wait(ch chan int) {
	time.Sleep(time.Second * 5)

	ch <- 2
}
