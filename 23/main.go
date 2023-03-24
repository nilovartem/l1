package main

import "fmt"

// Удалить i-ый элемент из слайса.
// generics in GO
func removeAt[T any](slice []T, position int) []T {
	return append(slice[:position], slice[position+1:]...)
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Initial slice", slice)
	slice = removeAt(slice, 1)
	fmt.Println("Updated slice", slice)

}
