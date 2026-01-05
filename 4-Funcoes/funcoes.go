package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func calculosMatematicos(n1, n2 int) (int, int) {
	// função anônima
	soma := n1 + n2
	subitracao := n1 - n2
	return soma, subitracao
}

func main() {
	fmt.Println(somar(10, 20))

	var funcaoSomar = func(a int, b int) int {
		return a + b
	}
	fmt.Println(funcaoSomar(30, 40))

	fmt.Println(funcaoSomar(50, 60))

	resultadoSoma, resultadoSubtracao := calculosMatematicos(100, 50)
	fmt.Println(resultadoSoma, resultadoSubtracao)

}
