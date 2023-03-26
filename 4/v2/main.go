package main

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/
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

func main() {
	var wg sync.WaitGroup
	//Устанавливаем флаг
	numWorkers := flag.Int("N", 1, "Number of workers")
	flag.Parse()
	fmt.Printf("Number of workers %v\n", *numWorkers)
	//Создаем канал для прерываний
	interrupt := make(chan os.Signal, 1)
	//Создаем канал, в который мы отправим горутинам сигнал о завершении работы
	exit := make(chan struct{}, 1)
	//Создаем канал для значений
	values := make(chan int, 1000)
	// signal.Notify регистрирует данный канал для
	// получений уведомлений об определенных сигналах от ОС
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	//Запускаем горутины
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
				wg.Wait()
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
		//С помощью неблокирующего оператора select читаем два канал
		//Как только придет значение в exit, горутина завершится
		select {
		case <-values:
			{
				logrus.Infof("Value %v was received by worker %v\n", value, id)
			}
		case <-exit:
			{
				//Делаем любые операции необходимые для завершения горутины
				logrus.Warnf("Worker %v is done", id)
				wg.Done()
				return
			}
		}
	}
}
