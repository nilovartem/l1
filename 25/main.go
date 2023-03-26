package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	//Таймер отправит значение в канал, когда истечет указанное время
	//Мы ожидаем значения из канала, блокируя функцию
	//Когда таймер завершится, функция прочитает канал и тоже завершится
	<-timer.C
}
func main() {
	//Ждем 5 секунд
	sleep(5 * time.Second)
	fmt.Println("Timer expired")
}
