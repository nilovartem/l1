package main

import (
	"fmt"
	"unsafe"
)

// Объявляем именованный тип Human
type Human struct {
	name     string
	language string
}

/*Объявляем методы*/
func (h *Human) say() {
	fmt.Printf("%v is speaking %v\n", h.name, h.language)
}
func (h *Human) run() {
	fmt.Printf("%v is running\n", h.name)
}

/*
Первый вариант - встраивание (embedding) поля
с именованным типом
*/
//Объявляем именованный тип ActionF
type ActionF struct {
	//Объявим анонимное поле
	Human
}

/*
Второй вариант - встраивание (embedding) поля
с указателем на именованный тип
Позволит сэкономить память, так как
поля типа не копируются
Это будет проверено в main
*/
//Объявляем именованный тип ActionS
type ActionS struct {
	//Объявим анонимное поле - указатель на именованный тип
	*Human
}

func main() {
	//Проверим стандартное поведение Human
	fmt.Println("Стандартное поведение Human")
	artem := Human{name: "Artem", language: "Russian"}
	artem.say()
	artem.run()
	fmt.Println("ActionF - обращаемся к экземпляру")
	//Передадим в ActionF экземляр структуры Human
	actionF := ActionF{artem}
	actionF.say()
	//Передадим в ActionS адрес экземляра структуры Human
	fmt.Println("ActionS - обращаемся к указателю на экземпляр")
	actionS := ActionS{&artem}
	actionS.run()
	/*Узнаем размер ActionF и ActionS
	Думаю, что при работе с большими структурами размеры
	могут повлиять на производительность*/
	fmt.Printf("Размер ActionF в байтах = %d\n", unsafe.Sizeof(actionF))
	fmt.Printf("Размер ActionS в байтах = %d\n", unsafe.Sizeof(actionS))
}
