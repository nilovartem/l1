package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// Разработать программу, которая будет
// последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
func main() {
	duration := flag.Int("N", 3, "-N=time")
	flag.Parse()
	//таймер пишет значение в свой канал, когда завершает работу
	timer := time.NewTimer(time.Second * time.Duration(*duration))
	read := make(chan int)
	go func() {
		for {
			select {
			case <-timer.C:
				{
					fmt.Println("Timer expired")
					close(read)
					return
				}
			default:
				{
					//отправляем значения в канал
					read <- rand.Intn(100)
				}
			}
		}
	}()
	//читаем значения из канала
	for value := range read {
		fmt.Println(value)
	}
}
