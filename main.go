package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("__Калькулятор индекса массы тела__")

	userHeight, userKg := getUserInput()
	printResult(calculateIMT(userHeight, userKg))
}

func calculateIMT(height float64, kg float64) float64 {
	const IMTPower = 2

	return kg / math.Pow(height/100, IMTPower)
}

func printResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}

func getUserInput() (height float64, kg float64) {
	var userHeight float64
	var userKg float64

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}
