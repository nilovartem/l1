package main

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
если нужно удалить элемент, ставим false*/
import "fmt"

// Создаем map с ключами из строк и значениями из структур
type Set map[string]bool

func main() {
	//Создаем slice из строк
	values := []string{"cat", "cat", "dog", "cat", "tree"}
	//Инициализируем тип
	set := make(Set)
	//Проходим по slice из строк и задаем значения в качестве ключей,
	//а в качестве значений - true
	for _, value := range values {
		set[value] = true
	}
	fmt.Println(values)
	fmt.Println(set)

}
