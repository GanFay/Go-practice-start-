package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := randomizer()
	fmt.Println(num)

	for i := 0; i < 10; {
		input(&num)
	}
}

func input(num *[]int) {
	input1 := ""
	fmt.Scanln(&input1)
	if input1 == "max" {
		fmt.Println(maxMas(*num))
	} else if input1 == "min" {
		fmt.Println(minMis(*num))
	} else if input1 == "generate" {
		*num = randomizer()
		fmt.Println(*num)
	} else {
		fmt.Println("Пиши max or min or generate")
	}
}

func maxMas(massive []int) int {
	max_massive := massive[0]
	for i := 0; i < len(massive)-1; i++ {
		if max_massive < massive[i+1] {
			max_massive = massive[i+1]
		}
	}
	return max_massive
}

func minMis(massive []int) int {
	min_Massive := massive[0]
	for i := 0; i < len(massive)-1; i++ {
		if min_Massive > massive[i+1] {
			min_Massive = massive[i+1]
		}
	}
	return min_Massive
}

func randomizer() []int {
	numbers := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		num := rand.Intn(1001)
		numbers = append(numbers, num)
	}
	return numbers
}
