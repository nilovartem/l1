package main

// Написать программу, которая конкурентно рассчитает
// значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

import (
	"fmt"
	"sync"
)

/*Функция для вычисления квадрата числа*/
func sqr(number int, wg *sync.WaitGroup) {
	//выполнится перед return : уменьшаем счетчик ожидаемых горутин
	defer wg.Done()
	fmt.Println(number * number)
}
func main() {
	fmt.Println("Starting main")
	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10} //создаем массив
	//Проходим по массиву
	for _, value := range array {
		//Увеличиваем счетчик ожидаемых горутин
		wg.Add(1)
		go sqr(value, &wg)
	}
	//Ждем, пока счетчик ожидаемых горутин не будет равен 0
	defer wg.Wait()
}
