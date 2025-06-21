package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2

	var userHeight float64
	var userKg float64

	fmt.Println("__Калькулятор индекса массы тела__")

	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userKg)

	fmt.Print("Ваш индекс массы тела: ")
	IMT := userKg / math.Pow(userHeight, IMTPower)

	fmt.Println(IMT)
}
