package main

/*Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

import (
	"fmt"
	"math"
)

// Выделяем группу из числа
func getGroup(value float64) int {
	//Обрезает до целой части - лучшая функция ever, долго не мог найти
	group := math.Trunc(value / 10)
	group *= 10
	return int(group)
}

// Создаем именованный тип
type Subset map[float64]struct{}

// Переделал на ключи и пустые структуры,
// Так как ключи не повторяются по определению

func main() {
	//Создаем массив температур
	data := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Объявляем и инициализируем map
	set := make(map[int]Subset)
	for _, value := range data {
		group := getGroup(value)
		//Если группа не существует, создаем подмножество
		if _, exists := set[group]; !exists {
			set[group] = make(Subset)
		}
		subset := set[group]
		subset[value] = struct{}{} //или bool = true
		set[group] = subset
	}
	fmt.Println(set)
}
