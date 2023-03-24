package main

import (
	"fmt"
	"sync"
)

func sqr(number int) {
	fmt.Println(number * number)
}

/*
Лучшая реализация, т.к код с WaitGroup
вынесен из функции sqr() и остается внутри main
*/
func main() {
	fmt.Println("Starting main")
	var wg sync.WaitGroup
	defer wg.Wait()
	array := [5]int{2, 4, 6, 8, 10}
	for _, value := range array {
		wg.Add(1)
		go func(value int) {
			sqr(value)
			wg.Done()
		}(value)
	}
}
