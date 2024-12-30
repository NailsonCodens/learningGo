package main

import "fmt"

func main() {
	var number1, number2 int
	var operation string
	fmt.Print("Digite o primero número: ")
	fmt.Scanf("%d", &number1)
	fmt.Print("Digite a operação (+, -, *, /, $)")
	fmt.Scanf("%s", &operation)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d", &number2)

	if operation == "+" {
		plus(number1, number2)
	} else if operation == "-" {
		result := minus(number1, number2)
		fmt.Printf("%d - %d = %d \n", number1, number2, result)
	} else if operation == "*" {
		result := multiplied(number1, number2)
		fmt.Printf("%d x %d = %d", number1, number2, result)
	} else if operation == "/" {
		result, rest := division(number1, number2)
		fmt.Printf("Quociente = %d; Resto = %d \n", result, rest)
	} else if operation == "$" {
		increment, decrement := incrementAndDecrement(number1, number2)
		fmt.Printf("Incremento = %d, Decrement = %d \n", increment, decrement)
	} else {
		fmt.Println("Operação inválida")
	}
}

func plus(n1 int, n2 int) {
	fmt.Printf("%d + %d = %d \n", n1, n2, n1+n2)
}

func minus(n1 int, n2 int) int {
	return n1 - n2
}

func multiplied(n1 int, n2 int) (result int) {
	result = n1 * n2
	return
}

func division(n1 int, n2 int) (int, int) {
	quocient := n1 / n2
	rest := n1 % n2
	return quocient, rest
}

func incrementAndDecrement(n1 int, n2 int) (inc int, dec int) {
	inc = n1 + n2
	dec = n1 - n2
	return
}
