package main

import (
	"fmt"
)

func reverse(str *string) (result string) {
	//buffer := []rune(*str)
	// В Go присутствует синтаксический сахар при обходе строки.
	// Если использовать конструкцию for range,
	// строка автоматически будет преобразована в []rune,
	//  то есть обход будет по Юникод символам:
	for _, symbol := range *str {
		result = string(symbol) + result
	}
	return
}
func main() {
	//default option
	var str string = "🥶🥵🧐"
	fmt.Print("Enter string:")
	fmt.Scanln(&str)
	fmt.Println("Your string:", str)
	str = reverse(&str)
	fmt.Println("Reversed string:", str)
}
