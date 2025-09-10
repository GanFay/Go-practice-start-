package main

import (
	"fmt"
)

func main() {
	start()
}

func start() {
	fmt.Println("Введите число факториал которого хотите найти:")
	var number uint
	scanln, err := fmt.Scanln(&number)
	if err != nil {
		return
	}
	fmt.Println(factorial(uint(scanln)))
	start()
}

func factorial(n uint) uint {
	var result uint = 1
	var i uint = 2
	for ; i <= n; i++ {
		result = result * i
	}
	return result
}
