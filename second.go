package main

import "fmt"

func main() {
	var num int16
	fmt.Scan(&num)
	switch {
	case num > 52:
		println("неписядва")
	case num < 52:
		println("неписядваточно")
	case num == 52:
		println("писядва")
	case num != 52:
		println("не писятдва точно точно")
	default:
		println("NO.")
	}
}
