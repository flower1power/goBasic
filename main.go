package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight float64
	var userKg float64

	fmt.Println("__Калькулятор индекса массы тела__")

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	printResult(calculateIMT(userKg, userHeight))
}

func calculateIMT(kg float64, height float64) float64 {
	const IMTPower = 2

	return kg / math.Pow(height/100, IMTPower)
}

func printResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}
