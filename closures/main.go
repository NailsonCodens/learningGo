package main

import "fmt"

func main() {
	var number1, number2 int

	var operation string

	fmt.Print("Digite o primeiro número")
	fmt.Scanf("%d", &number1)
	fmt.Print("Digite a operação (+, -, *, /, $: ")
	fmt.Scanf("%s", &operation)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanf("%d", &number2)
	var operationMethod = func(n1 int, n2 int) (int, int)

	if operation == "+" {
		operationMethod = func(n1 int, n2 int) (int, int) {
			return n1 + n2, 0
		}
	} else if operation == "-" {
		operationMethod = func(n1 int, n2 int) (int, int) {
			return n1 - n2, 0
		}
	} else if operation == "*" {
		operationMethod = func(n1 int, n2 int) (int, int) {
			return n1 * n2, 0
		}
	} else if operation == "/" {
		operationMethod = func(n1 int, n2 int) (int, int) {
			return n1 / n2, n1 % n2
		}
	} else if operation == "$" {
		operationMethod = func(n1 int, n2 int) (int, int) {
			return n1 + n2, n1 - n2
		}
	}

	result1, result2 := operationMethod(number1, number2)
	fmt.Printf("%d %s %d = %d, %d \n", number1, operation, number2, result1, result2)

}
