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
	fmt.Scanln(&number)
	fmt.Println(factorial(number))
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
