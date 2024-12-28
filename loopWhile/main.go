package main

import "fmt"

func main() {
	var number int
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &number)

	i := 0
	for i <= 10 {
		fmt.Printf("%d x %d = %d \n", number, i, number*i)
		i++
	}
}
