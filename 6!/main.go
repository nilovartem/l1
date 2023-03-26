package main

import (
	"fmt"
)

//Реализовать все возможные способы остановки выполнения горутины.

// 1. Используем отдельный сигнальный канал,
// и используя неблокирующий оператор select ждем,
// пока в канал поступит любое значение, пусть даже struct{}{}
// Такое значение будет означать остановку горутины
// func signal(exit <-chan struct{}) {
// 	for {
// 		select {
// 		case <-exit:
// 			fmt.Println("Exit")
// 			return
// 		default:
// 			//Делаем какие-либо задачи
// 			time.Sleep(time.Second * 2)
// 			fmt.Println("Goroutine signal is working now")
// 		}
// 	}

// }
func close(data chan int) {
	for value := range data {
		fmt.Printf("Goroutine close is working on value %v\n", value)
	}
}
func main() {
	//1. Остановка горутины сигнальным каналом
	// exit := make(chan struct{})
	// go signal(exit)
	// fmt.Println("Main is working now")
	// time.Sleep(time.Second * 3)
	// exit <- struct{}{}
	//2. Остановка "опустевшим каналом"
	data := make(chan int)
	go close(data)
	data <- 0
	data <- 1
	data <- 2
	close(data)
	//time.Sleep(time.Second * 2)
}
