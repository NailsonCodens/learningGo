package main

import "fmt"

func main() {
	var number int
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &number)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", number, i, number*i)
	}
}
