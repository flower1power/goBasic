package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")

	printResult(calculateIMT(getUserInput()))
}

func calculateIMT(height float64, kg float64) float64 {
	return kg / math.Pow(height/100, IMTPower)
}

func printResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f", IMT)
}

func getUserInput() (height float64, kg float64) {
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&kg)

	return height, kg
}
