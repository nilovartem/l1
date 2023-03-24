package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива
// (quicksort) встроенными методами языка.
// разобрал и написал код из грокаем алгоритмы
func main() {
	array := []int{8, 2, 10, 1}
	//array = [23]int(quicksort(array[:]))
	fmt.Printf("Original array = %v\n", array)
	fmt.Printf("Sorted array = %v\n", quicksort(array))
	//array = sort(array)
}
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

		//fmt.Printf("Pivot = %v\n", pivot)
		//fmt.Printf("Less = %v\n", less)
		//fmt.Printf("Greater = %v\n", greater)
		var middle []int
		middle = append(middle, pivot)
		//less = append(quicksort(less), middle...)
		array = append(quicksort(less), append(middle, quicksort(greater)...)...)
		//fmt.Printf("Array = %v\n", array)
		return array
	}
}
