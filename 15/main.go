package main

import (
	"fmt"
	"math/rand"
)

var justString string

func createHugeString(size int64) (result string) {
	var i int64
	for i = 0; i < size; i++ {
		//Добавляем к строке символы Unicode,
		//преобразовав псевдослучайное число в rune
		//В качестве диапазона взяты символы до Latin Extended Additional
		result += string(rune(32 + rand.Intn(591-32)))
	}
	return
}
func someFunc() {
	//Создаем строку из 1024 байт
	v := createHugeString(1 << 10)
	justString = v[:100]
	/*Строка - это слайс из байт.
	Это значит, что нарезая строку таким образом, мы получаем 100 байт
	Однако, один символ может состоять из нескольких байт,
	которые мы могли просто не записать
	Из-за недостатка байт символы часто интерпретируются как �
	Чтобы исправить ошибку, мы должны нарезать слайс иначе*/

}
func newSomeFunc() {
	v := createHugeString(1 << 10)
	/*Нарезав строку как руны (int32 коды символов стандарта Unicode),
	мы правильно определим байтовые границы каждого символа*/
	runes := []rune(v)
	justString = string(runes[:100])
}

func main() {
	someFunc()
	fmt.Printf("Damaged string: %v\n", justString)
	newSomeFunc()
	fmt.Printf("Properly defined string: %v\n", justString)
}
