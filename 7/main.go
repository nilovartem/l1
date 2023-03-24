package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// https://habr.com/ru/post/338718/
// type mutex sync.RWMutex

type example struct {
	key   int
	value int
}

type concurrentMap struct {
	mutex sync.Mutex
	data  map[int]int
}

func (cMap *concurrentMap) write(key, value int) {
	cMap.mutex.Lock() //=недоступен ни для чего
	defer cMap.mutex.Unlock()
	cMap.data[key] = value
}
func (cMap *concurrentMap) read(key int) (int, bool) {
	cMap.mutex.Lock()         //=доступен только для чтения
	defer cMap.mutex.Unlock() //=доступен только для чтения и записи
	value, ok := cMap.data[key]
	return value, ok
}

// Реализовать конкурентную запись данных в map.
func main() {
	var wg sync.WaitGroup
	storage := concurrentMap{data: make(map[int]int)}
	data := make(chan example)
	//создаем горутины
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(data <-chan example, storage *concurrentMap, id int) {
			defer wg.Done()
			for object := range data {
				fmt.Printf("Writer %v working now on key %v\n", id, object.key)
				storage.write(object.key, object.value)
			}
		}(data, &storage, i)
	}
	for i := 0; i < 100; i++ {
		data <- example{key: i, value: rand.Intn(100)}
	}
	close(data)
	wg.Wait()
	fmt.Println(storage.data)
}
