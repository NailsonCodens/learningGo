package main

import "fmt"

func main() {
	var month int
	fmt.Print("Digite o número do mês: ")
	fmt.Scanf("%d", &month)
	switch month {
	case 1:
		fmt.Println("Este m6es é janeiro.")
	case 2:
		fmt.Println("Este m6es é fevereiro.")
	case 3:
		fmt.Println("Este m6es é março.")
	case 4:
		fmt.Println("Este m6es é abril.")
	case 5:
		fmt.Println("Este m6es é maio.")
	case 6:
		fmt.Println("Este m6es é junho.")
	case 7:
		fmt.Println("Este m6es é julho.")
	case 8:
		fmt.Println("Este m6es é agosto.")
	case 9:
		fmt.Println("Este m6es é setembro.")
	case 10:
		fmt.Println("Este m6es é outubro.")
	case 11:
		fmt.Println("Este m6es é novembro.")
	case 12:
		fmt.Println("Este m6es é dezembro.")
	default:
		fmt.Println("Mês inválido.")
	}
	switch month {
	case 1, 2, 3:
		fmt.Println("Primeiro trimestre")
	case 4, 5, 6:
		fmt.Println("Segundo trimestre")
	case 7, 8, 9:
		fmt.Println("Terceiro trimestre")
	case 10, 11, 12:
		fmt.Println("Quarto trimestre")
	}
}
