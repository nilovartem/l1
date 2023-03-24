package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет,
// что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// основана на том, что в мапе должны быть уникальные ключи
func distinctionCheck(str string) bool {
	var mapping = make(map[rune]struct{})
	for _, value := range str {
		if _, exists := mapping[value]; exists {
			return false
		}
		mapping[value] = struct{}{}
	}
	return true
}
func main() {
	var str string
	fmt.Print("Enter your string:")
	fmt.Scanln(&str)
	fmt.Printf("Result of check = %v\n", distinctionCheck(strings.ToLower(str)))
}
