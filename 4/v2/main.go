package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают
// произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	//Создаем
	var wg sync.WaitGroup
	//разобраться во всех параметрах
	numWorkers := flag.Int("N", 1, "Number of workers")
	flag.Parse()
	fmt.Printf("Number of workers %v\n", *numWorkers)
	interrupt := make(chan os.Signal, 1)
	exit := make(chan struct{}, 1)
	values := make(chan int, 1000)
	// signal.Notify registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(exit, values, &wg, i)
	}
	for {
		select {
		case <-interrupt:
			{
				logrus.Errorln("Received SIGINT, quiting")
				exit <- struct{}{}
				close(exit)
				close(values)
				defer wg.Wait()
				return
			}
		default:
			{
				values <- rand.Intn(100)
			}
		}
	}
}

func worker(exit <-chan struct{}, values <-chan int, wg *sync.WaitGroup, id int) {
	for value := range values {
		select {
		case <-values:
			{
				logrus.Infof("Value %v was received by worker %v\n", value, id)
			}
		case <-exit:
			{
				//clean up
				logrus.Warnf("Worker %v is done", id)
				wg.Done()
				return
			}
		}
	}
}
