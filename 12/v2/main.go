package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
// если нужно удалить элемент, ставим false
type Set map[string]bool

func main() {
	values := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(Set)
	for _, value := range values {
		set[value] = true
	}
	fmt.Println(values)
	fmt.Println(set)

}
