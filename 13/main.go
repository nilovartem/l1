package main

/*Поменять местами два числа без
создания временной переменной.*/

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Printf("Value of a = %v, value of b = %v\n", a, b)
	//Либо так
	a, b = b, a
	//Либо так
	//a = a + b // 5 + 10 = 15
	//b = a - b // 15 - 10 = 5
	//a = a - b // 15 - 5 = 10
	fmt.Printf("Value of a = %v, value of b = %v\n", a, b)
}
