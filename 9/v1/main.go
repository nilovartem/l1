package main

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
import (
	"fmt"
)

// Нам не нужно использовать здесь WaitGroup,
// т.к чтение из канала блокирует
// возможное завершение main, пока канал не опустеет
// Горутины выполняются не по порядку, а асинхронно (наверное)
func main() {
	regular := make(chan int)
	doubled := make(chan int)
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	go func() {
		for _, value := range numbers {
			//fmt.Printf("Send value %v\n", value)
			regular <- value
		}
		close(regular)
	}()
	go func() {
		for value := range regular {
			//fmt.Printf("Double the value %v\n", value)
			doubled <- value * 2
		}
		close(doubled)
	}()
	for value := range doubled {
		//fmt.Printf("Receiving the value %v\n", value)
		fmt.Println(value)
	}
}
