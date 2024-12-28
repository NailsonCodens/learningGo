package main

import "fmt"

func main() {
	/*var age int
	fmt.Print("Digite sua idade: ")
	fmt.Scanf("%d", &age)

	if age >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}*/

	var number int
	fmt.Println("Digite um número: ")
	fmt.Scanf("%d", &number)

	if number < 0 {
		fmt.Println("Este número é negativo!")
	} else if number >= 0 && number <= 9 {
		fmt.Println("Este número é positivo e contém mais de um dígito!")
	} else {
		fmt.Println("Este número é positivo e contém mais de um dígito")
	}
}
