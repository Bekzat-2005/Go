package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64 // 0.0
	fmt.Print("Рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Вес: ")
	fmt.Scan(&userKg)
	var IMT = userKg / math.Pow(userHeight, IMTPower)
	fmt.Printf("Ваш индек масса тела %.0f", IMT)

}
