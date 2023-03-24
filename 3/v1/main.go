package main

import (
	"fmt"
	"time"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….)
с использованием
конкурентных вычислений.
*/
/*
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}*/
func main() {
	//defer timer("main")()
	start := time.Now()
	numbers := [...]int{2, 4, 6, 8, 10}
	squared := make(chan int)
	go func() {
		for _, value := range numbers {
			fmt.Printf("Sending value %v\n", value)
			squared <- value * value
		}
		close(squared)
	}()
	sum := 0
	for value := range squared {
		fmt.Printf("Squared value equals %v\n", value)
		sum += value
	}
	fmt.Printf("Sum = %v\n", sum)
	fmt.Printf("took %v\n", time.Since(start))

}
