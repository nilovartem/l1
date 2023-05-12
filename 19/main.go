package main

import (
	"fmt"
)

func reverse(str *string) (result string) {
	//buffer := []rune(*str)
	// В Go присутствует синтаксический сахар при обходе строки.
	// Если использовать конструкцию for range,
	// строка автоматически будет преобразована в []rune,
	//  то есть обход будет по Юникод символам,
	// поэтому нужно преобразовывать результат в string
	for _, symbol := range *str {
		result = string(symbol) + result
	}
	return
}
func main() {
	//стандартный вариант
	var str string = "🥶🥵🧐"
	fmt.Print("Enter string:")
	fmt.Scanf("%q", &str)
	fmt.Println("Your string:", str)
	str = reverse(&str)
	fmt.Println("Reversed string:", str)
}
