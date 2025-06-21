package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")

	for {
		IMT, err := calculateIMT(getUserInput())

		if err != nil {
			fmt.Println(err)
			if !checkRerun() {
				break
			}
			continue
		}

		printResult(IMT)

		if !checkRerun() {
			break
		}
	}
}

func calculateIMT(height float64, kg float64) (float64, error) {
	if height <= 0 || kg <= 0 {
		return 0, errors.New("не указан вес или рост")
	}

	IMT := kg / math.Pow(height/100, IMTPower)

	return IMT, nil
}

func printResult(IMT float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f \n", IMT)

	switch true {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("Норма")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func getUserInput() (height float64, kg float64) {
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&kg)

	return height, kg
}

func checkRerun() bool {
	var rerun string

	fmt.Println("Посчитать нового человека? (Y|n)")
	fmt.Scan(&rerun)

	if rerun == "Y" || rerun == "y" {
		return true
	}

	return false
}
