package main

import (
	"fmt"
	"slices"
)

func main() {
	otvet := "losharar"
	massive := []string{"egg", "tea", "f"}
	println("Привет что ты ел на завтрак?")
	fmt.Scan(&otvet)

	found := slices.Contains(massive, otvet)
	if found == true {
		println("Я угадал!!")
	} else {
		println("Я не угадал!!esketit")
	}

}
