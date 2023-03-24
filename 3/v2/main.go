package main

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(2^2+3^2+4^2….)
с использованием
конкурентных вычислений.
*/
//Улучшен - без data race
import (
	"fmt"
	"sync"
	"time"
)

/*
Функция таймер для подсчета времени выполнения программы
*/
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

var sum int = 0

func main() {
	defer timer("main")()
	var mutex sync.Mutex
	numbers := [...]int{2, 4, 6, 8, 10}
	regular := make(chan int)
	done := make(chan struct{})
	//result := make(chan int, 1)
	go func() {
		for _, value := range numbers {
			regular <- value
		}
		close(regular)
	}()
	//Запускаем горутины
	for i := 0; i < len(numbers); i++ {
		go worker(regular, &mutex, done)
	}
	for i := 0; i < len(numbers); i++ {
		<-done
	}
	fmt.Println(sum)

}
func worker(regular <-chan int, mutex *sync.Mutex, done chan<- struct{}) {
	defer func() { done <- struct{}{} }()
	mutex.Lock()
	value := <-regular
	sum += value * value
	mutex.Unlock()
}
