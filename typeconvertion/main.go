package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string
	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &text)

	/*var number int
	number, _ = strconv.Atoi(text)
	fmt.Println(number)*/

	number, _ := strconv.ParseInt(text, 10, 32)
	var conv = float64(number)
	fmt.Println(number)
	fmt.Println(conv)
}
