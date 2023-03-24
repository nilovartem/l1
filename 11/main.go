package main

import "fmt"

func main() {
	first := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	second := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := []int{}
	check := map[int]struct{}{}
	for _, value := range first {
		check[value] = struct{}{}
	}
	for _, value := range second {
		if _, ok := check[value]; ok {
			result = append(result, value)
		}
	}
	fmt.Println(result)
}
