package main

import (
	"fmt"
	"sync"
)

func sqr(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(number * number)
}
func main() {
	fmt.Println("Starting main")
	var wg sync.WaitGroup
	defer wg.Wait() //вызовется последним
	array := [5]int{2, 4, 6, 8, 10}
	for _, value := range array {
		wg.Add(1)
		go sqr(value, &wg)
	}
	//fmt.Scanln()
	//runtime.Goexit()

}
