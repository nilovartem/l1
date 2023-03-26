package main

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/
import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	//Разбиваем строку на массив,
	//состоящий из слов (точнее из объектов)
	temp := strings.Fields(str)
	builder := strings.Builder{}
	for i := len(temp) - 1; i >= 0; i-- {
		//Создаем новую строку и ставим убранные пробелы между словами
		builder.WriteString(temp[i] + " ")
	}
	reversed := strings.TrimSpace(builder.String())
	fmt.Println(reversed)
}
