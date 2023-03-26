package main

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

import "fmt"

// Инвертировать бит
func toggleBit(num *int64, position uint) {
	//Используем оператор деления по модулю 2 (XOR)
	*num ^= 1 << position
}
func main() {
	var value int64 = 8
	//Выводим 8 бит
	fmt.Printf("%.8b\n", value)
	toggleBit(&value, 3)
	//Выводим 8 бит
	fmt.Printf("%.8b\n", value)
}
