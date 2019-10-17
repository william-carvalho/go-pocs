package main

import (
	"fmt"
	"math"
)

func main() {
	a := 8
	b := 5

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Módulo = >", a%b)

	//bitwise
	fmt.Println("E =>", a&b)  //11 & 10 = 10
	fmt.Println("Ou =>", a|b) //11 | 10 = 11
	fmt.Println("Ou =>", a^b) //11 ^ 10 = 01

	// Outras operacoes usando math

	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
