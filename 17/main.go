package main

import (
	"github.com/sirupsen/logrus"
)

// определяем интерфейс с ограничениями
// | - это оператор ИЛИ
type MyConstraint interface {
	int | int8 | int16 | int32 | int64 | float64
}

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
	//values := []string{"John", "James", "Jacob", "Jeremy", "Alex"}
	values := []int{10, 20, 30, 40, 50}
	position, ok := binarySearch(values, 10)
	if ok {
		logrus.Infof("Index of value is %v\n", position)
	} else {
		logrus.Errorln("Can't find the value!")
	}
	//fmt.Printf("Index of Jacob is %v", binarySearch(values))
}
