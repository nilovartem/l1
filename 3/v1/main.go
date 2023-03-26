package main

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….)
с использованием
конкурентных вычислений.
*/
import (
	"fmt"
)

func main() {
	//Создаем массив, так как числа уже известны
	numbers := [...]int{2, 4, 6, 8, 10}
	//Создаем канал, куда будут отправляться квадраты чисел
	squared := make(chan int)
	go func() {
		//Проходим по массиву
		for _, value := range numbers {
			fmt.Printf("Sending value %v\n", value)
			//Отправляем квадраты чисел в канал
			squared <- value * value
		}
		//Закрываем канал
		close(squared)
	}()
	sum := 0
	//Читаем канал
	for value := range squared {
		fmt.Printf("Squared value equals %v\n", value)
		//Прибавляем значение из канала в сумму
		sum += value
	}
	fmt.Printf("Sum = %v\n", sum)
}
