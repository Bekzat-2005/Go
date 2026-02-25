package main

import "fmt"

func main() {

	tr := make([]string, 0, 2)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1", "2")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	fmt.Println(len(tr), cap(tr))
}

// 5

// tr1 := []int{1, 2, 3}
// tr2 := []int{100, 200, 300}
// tr1 = append(tr1, tr2...)
// fmt.Println(tr1)

// for i, v := range tr1 {
// 	fmt.Println(i, v)
// }

// 4

// 	fmt.Print("San jaz: ")
// 	tran := []float64{}
// 	// var sum float64
// 	for {
// 		tranScan := sanScan()
// 		if tranScan == 0 {
// 			break
// 		}
// 		tran = append(tran, tranScan)
// 		// sum += tranScan
// 	}
// 	fmt.Println(tran)
// 	fmt.Println(calculate(tran))
// 	// calculate(tran)

// 	// fmt.Println(sum)
// }
// func sanScan() float64 {
// 	var num float64

// 	fmt.Scan(&num)
// 	return num
// }
// func calculate(tran []float64) float64 {
// 	var balance float64
// 	for _, v := range tran {
// 		balance += v
// 	}
// 	return balance
// 	// fmt.Println(balance)
// }

// 3

// tran := []int{10, 2}
// tran = append(tran, 100)
// fmt.Println(tran)

// 2

// tran := []int{10, 2, 8, 19, 5, 7}
// tranPartial := tran[1:5]
// tranNewPartial := tranPartial[:1]
// tranNewPartial[0] = 30

// tranNewPartial = tranNewPartial[0:4]
// fmt.Println(tran)
// fmt.Println(tranNewPartial)
// fmt.Println(len(tranPartial), cap(tranPartial))
// fmt.Println(len(tranNewPartial), cap(tranNewPartial))

// 1
// tran := [5]int{10, 2, 19, 5, 7}
// banks := [2]string{"kaspi", "halyk"}

// fmt.Println(tran)
// fmt.Println(banks)
// partial := tran[2:3]
// fmt.Println(partial)
