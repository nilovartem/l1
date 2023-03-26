package main

//Реализовать конкурентную запись данных в map.
import (
	"fmt"
	"math/rand"
	"sync"
)

// Пример данных, которые мы будем получать
// по каналу для записи в map
type example struct {
	key   int
	value int
}

// Создаем структура с безопасными для записи map
type concurrentMap struct {
	mutex sync.Mutex
	data  map[int]int
}

// Создаем безопасный метод для записи в map
func (cMap *concurrentMap) write(key, value int) {
	cMap.mutex.Lock() //недоступен ни для чего
	defer cMap.mutex.Unlock()
	cMap.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	//Объявляем и инициализируем хранилище
	storage := concurrentMap{data: make(map[int]int)}
	//Создаем канал для приема данных
	data := make(chan example)
	//Создаем горутины
	for i := 1; i <= 10; i++ {
		//Увеличиваем количество ожидаемых горутин
		wg.Add(1)
		go func(data <-chan example, storage *concurrentMap, id int) {
			//Уменьшаем количество ожидаемых горутин на 1
			defer wg.Done()
			//Ожидаем данные из канала, и записываем безопасным методом в map
			for object := range data {
				fmt.Printf("Writer %v working now on key %v\n", id, object.key)
				storage.write(object.key, object.value)
			}
		}(data, &storage, i)
	}
	//Рандомим данные и отправляем в канал
	for i := 0; i < 100; i++ {
		data <- example{key: i, value: rand.Intn(100)}
	}
	//Закрываем канал
	close(data)
	//Ждем, пока количество ожидаемых горутин не станет равным 0
	wg.Wait()
	fmt.Println(storage.data)
}
