package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/
import "fmt"

// Создаем map с ключами из строк и значениями из структур
type Set map[string]struct{}

func main() {
	//Создаем slice из строк
	values := []string{"cat", "cat", "dog", "cat", "tree"}
	//Инициализируем тип
	set := make(Set)
	//Проходим по slice из строк и задаем значения в качестве ключей,
	//а пустые структуры в качестве значений
	for _, value := range values {
		set[value] = struct{}{}
	}
	fmt.Println(values)
	fmt.Println(set)
}
