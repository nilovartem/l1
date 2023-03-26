package main

/*
Разработать программу, которая перемножает, делит, складывает,
вычитает две числовых переменных a,b, значение которых > 2^20.
*/

import (
	"fmt"
	"math/big"
)

func main() {
	var a = big.NewInt(2 << 20)
	var b = big.NewInt(2 << 25)
	var bigInt big.Int
	fmt.Printf("A = %v, B = %v\n", a, b)
	fmt.Printf("A * B = %v\n", bigInt.Mul(a, b))
	fmt.Printf("A / B = %v\n", bigInt.Div(a, b))
	fmt.Printf("A + B = %v\n", bigInt.Add(a, b))
	fmt.Printf("A - B = %v\n", bigInt.Sub(a, b))
}
