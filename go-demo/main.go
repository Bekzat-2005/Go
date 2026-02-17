package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	userHeight, userKg := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	// isLean := IMT < 16
	outputResult(IMT)

	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас  дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный массы тела")
	case IMT < 30:
		fmt.Println("У вас избыточный массы тела")
	default:
		fmt.Println("У вас степень ожирения массы тела")
	}

	// if IMT < 16 {
	// 	fmt.Println("У вас сильный дефицит массы тела")
	// } else if IMT < 18.5 {
	// 	fmt.Println("У вас  дефицит массы тела")
	// } else if IMT < 25 {
	// 	fmt.Println("У вас нормальный массы тела")
	// } else if IMT < 30 {
	// 	fmt.Println("У вас избыточный массы тела")
	// } else {
	// 	fmt.Println("У вас степень ожирения массы тела")
	// }

}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индек масса тела %.1f", imt)
	fmt.Println(result)
}
func calculateIMT(userKg float64, userHeight float64) float64 {

	var IMT = userKg / math.Pow(userHeight/100, IMTPower)
	return IMT
}
func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64 // 0.0
	fmt.Print("Рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Вес: ")
	fmt.Scan(&userKg)
	return userHeight, userKg
}
