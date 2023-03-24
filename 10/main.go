package main

import (
	"fmt"
	"math"
)

// https://stackoverflow.com/questions/47544156/what-uses-a-type-with-empty-struct-has-in-go
// https://www.sohamkamani.com/golang/sets/

// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func getGroup(value float64) int {
	group := math.Trunc(value / 10) //обрезает до целой части - лучшая функция ever
	group *= 10
	return int(group)
}

type Subset map[float64]struct{}

// переделать на ключи и пустые структуры,
// так как ключи не повторяются по определению
func main() {
	data := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	set := make(map[int]Subset)
	for _, value := range data {
		group := getGroup(value)
		if _, exists := set[group]; !exists {
			set[group] = make(Subset)
		}
		subset := set[group]
		subset[value] = struct{}{} //или bool = true
		set[group] = subset
	}
	fmt.Println(set)
}
