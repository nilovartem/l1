package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
type Set map[string]struct{}

func main() {
	values := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(Set)
	for _, value := range values {
		set[value] = struct{}{}
	}
	fmt.Println(values)
	fmt.Println(set)
}
