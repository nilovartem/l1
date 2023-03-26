package main

import (
	"github.com/sirupsen/logrus"
)

// Определяем интерфейс с ограничениями по типу
// | - это оператор ИЛИ
type MyConstraint interface {
	int | int8 | int16 | int32 | int64 | float64
}

// Определяем функцию, принимающую только
// слайс значений определенного типа
func binarySearch[T MyConstraint](values []T, value T) (position int, error bool) {
	low := 0
	high := len(values) - 1
	for low <= high {
		mid := (low + high)
		guess := values[mid]
		if guess == value {
			return mid, true
		}
		if guess > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
func main() {
	//Создаем слайс со значениями,
	//для бинарного поиска должен быть уже отсортирован
	values := []int{10, 20, 30, 40, 50}
	//Ищем позицию цифры 30
	position, ok := binarySearch(values, 30)
	//Если ошибка, то значение было не найдено
	if ok {
		logrus.Infof("Index of value is %v\n", position)
	} else {
		logrus.Errorln("Can't find the value!")
	}
}
