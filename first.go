package main

import (
	"fmt"
	"slices"
)

func lox() {
	var otvet string
	massive := []string{"egg", "tea", "f"}
	println("Привет что ты ел на завтрак?")
	fmt.Scan(&otvet)

	found := slices.Contains(massive, otvet)
	if found == true {
		println("Я угадал!!")
	} else {
		println("Какашки")
	}

}
