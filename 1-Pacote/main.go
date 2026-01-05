package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(auxiliar.EscrevaMensagem())
	erro := checkmail.ValidateFormat("email@example.com")

	fmt.Println("Validar email:", erro)

}
