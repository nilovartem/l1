package main

/*
Реализовать структуру-счетчик, которая будет
инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

// Создаем именованный тип Counter
type Counter struct {
	//Добавляем Mutex для обеспечения единиличного доступа
	m sync.Mutex
	//Значение счетчика
	value int
}

// Создаем для удобства метод добавления в счетчик
func (counter *Counter) add(value int) {
	//Блокируем доступ для всех горутин, кроме текущей
	counter.m.Lock()
	//Разрешаем доступ для всех горутин, кроме текущей
	defer counter.m.Unlock()
	counter.value += value
}
func main() {
	//Объявляем и инициализируем счетчик
	counter := Counter{value: 0}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1) //Повышаем кол-во ожидаемых горутин
		//Создаем анонимную горутину
		//Параметр i передаем явно, так как возможна ошибка
		//из-за задержки выполения
		go func(id int) {
			for i := 0; i < 10; i++ {
				fmt.Printf("Goroutinge №%v is working\n", id)
				counter.add(1)
			}
			wg.Done() //Понижаем кол-во ожидаемых горутин
		}(i)
	}
	wg.Wait() //Ждем, пока счетчик ожидаемых горутин не будет равен 0
	logrus.Infof("Counter value = %v", counter.value)

}
