package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// 1. Используем отдельный сигнальный канал,
// и используя неблокирующий оператор select ждем,
// пока в канал поступит любое значение, пусть даже struct{}{}
// Такое значение будет означать остановку горутины
func v1(exit <-chan struct{}) {
	for {
		select {
		case <-exit:
			fmt.Println("v1 finished")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("Goroutine signal is working now")
		}
	}

}
func v2(data chan int) {
	for value := range data {
		fmt.Printf("Goroutine close is working on value %v\n", value)
	}
	fmt.Println("v2 finished")
}
func v3(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("v3 finished")
}
func v4(data chan int) {
	<-data
	fmt.Println("v4 finished")
}
func main() {
	//1. Остановка горутины сигнальным каналом
	exit := make(chan struct{})
	go v1(exit)
	fmt.Println("Main is working now")
	time.Sleep(time.Second * 3)
	exit <- struct{}{}
	//2. Остановка "опустевшим каналом" через range
	data := make(chan int)
	go v2(data)
	for i := 0; i < 5; i++ {
		data <- i
	}
	close(data)
	//3. Закрытие с помощью контекста.
	ctx, cancel := context.WithCancel(context.Background())
	go v3(ctx)
	cancel()
	//4. Читаем из канала и завершаем горутину
	c1 := make(chan int)
	go v4(c1)
	c1 <- 0
}
