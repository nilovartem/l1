package main

import (
	"fmt"
	"sync"
)

/*Функция для вычисления квадрата числа*/
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
	//Ждем, пока счетчик ожидаемых горутин не будет равен 0
	defer wg.Wait() //выполнится перед return
	array := [5]int{2, 4, 6, 8, 10}
	for _, value := range array {
		//Увеличиваем счетчик ожидаемых горутин
		wg.Add(1)
		//функциональность горутины оставляем в анонимной функции
		go func(value int) {
			sqr(value)
			//Уменьшаем счетчик ожидаемых горутин
			wg.Done()
		}(value)
	}
}
