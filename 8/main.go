package main

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"

func toggleBit(num *int64, position uint) {
	*num ^= 1 << position
}
func main() {
	var value int64 = 8
	fmt.Printf("%.8b\n", value)
	toggleBit(&value, 3)
	fmt.Printf("%.8b\n", value)
}
