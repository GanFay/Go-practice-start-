package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "горутина 1"
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

	}()
	go func() {
		ch <- "горутина 2"
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
	}()
	fmt.Println(<-ch)
}
