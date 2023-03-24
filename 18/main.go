package main

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

// Реализовать структуру-счетчик, которая будет
// инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	m     sync.Mutex
	value int
}

func (counter *Counter) add(value int) {
	counter.m.Lock()
	defer counter.m.Unlock()
	counter.value += value
}
func main() {
	counter := Counter{value: 0}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			for i := 0; i < 10; i++ {
				fmt.Printf("Goroutinge №%v is working\n", id)
				counter.add(1)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	logrus.Infof("Counter value = %v", counter.value)

}
