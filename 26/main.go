package main

/*Разработать программу, которая проверяет,
что все символы в строке уникальные
(true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.*/
import (
	"fmt"
	"strings"
)

// Основана на том, что в map должны быть уникальные ключи
// По сути, то же воплощение, что и в задании со множествами
func distinctionCheck(str string) bool {
	//создаем map с ключами из "символов" и с "пустым значением"
	var mapping = make(map[rune]struct{})
	for _, value := range str {
		//Если символ уже существует, то символы не уникальны
		if _, exists := mapping[value]; exists {
			return false
		}
		//Добавляем символ в качестве ключа map
		mapping[value] = struct{}{}
	}
	return true
}
func main() {
	var str string
	fmt.Print("Enter your string:")
	//Передаем указатель на строку, экономя память
	fmt.Scanln(&str)
	//Вызовем функцию, предварительно переведя строку в нижний регистр
	fmt.Printf("Result of check = %v\n", distinctionCheck(strings.ToLower(str)))
}
