package main

import "fmt"

func main() {
	fmt.Println("Bem-vindo!")
	defer fmt.Println("Obrigado por utilizar nosso software")
	fmt.Println("Digite um número maior do que 10")
	var number int
	fmt.Scanf("%s", number)

	if number <= 10 {
		panic("Seu número é menor que 10.")
	} else {
		fmt.Println("Número ok")
	}
}
