package main

import "fmt"

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.
// пустой интерфейс может держать значения любого типа
func printType(value interface{}) {
	fmt.Printf("Type = %T\n", value)
}
func main() {
	printType(int(10))
	printType(string("Hello"))
	printType(bool(true))
	printType(make(chan int))
}
