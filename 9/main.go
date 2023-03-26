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
func main() {
	//Cоздаем два канала - для обычных чисел и для удвоенных
	regular := make(chan int)
	doubled := make(chan int)
	//Создаем массив из чисел
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//Создаем анонимную горутину для заполнения канала
	go func() {
		for _, value := range numbers {
			//Отправляем числа в канал
			regular <- value
		}
		//Закрываем канал
		close(regular)
	}()
	go func() {
		//Читаем канал и горутина не умирает
		for value := range regular {
			doubled <- value * 2
		}
		//Закрываем канал
		close(doubled)
	}()
	//Ждем, пока значения придут в канал и выводим их
	for value := range doubled {
		fmt.Println(value)
	}
}