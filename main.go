package main

import (
	"fmt"
	"math"
)

func main() {
	marks := []float64{2, 4, 5}
	var sum float64 = 0

	for i := 0; i < len(marks); i++ {
		sum += marks[i]
	}
	var result float64 = sum / float64(len(marks))
	result_final := fmt.Sprintf("%.2f", math.Round(result*100)/100)
	println(result_final)

}
