package main

/*Реализовать быструю сортировку массива
(quicksort) встроенными методами языка.*/
import (
	"fmt"
)

func main() {
	//Создаем слайс со значениями
	array := []int{8, 2, 10, 1}
	fmt.Printf("Original slice = %v\n", array)
	//Вызвать функцию
	fmt.Printf("Sorted slice = %v\n", quicksort(array))
}

// Функция, принимает slice int и возвращает slice int
func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		var less []int
		for _, value := range array[1:] {
			if value <= pivot {
				less = append(less, value)
			}
		}
		var greater []int
		for _, value := range array[1:] {
			if value > pivot {
				greater = append(greater, value)
			}
		}
		var middle []int
		middle = append(middle, pivot)
		//Немножко клоунада, но я не знаю как по другому объединить три слайса
		array = append(quicksort(less), append(middle, quicksort(greater)...)...)
		return array
	}
}
