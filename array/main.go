package main

import "fmt"

func main() {
	var friends [5]string

	fmt.Println(friends)
	for i := 0; i < len(friends); i++ {
		fmt.Print("Digite o nome dos seus amigos: \n")
		fmt.Scanf("%s", &friends[i])
	}

	fmt.Println(friends)
	fmt.Println("Seus amigos são: ")
	for _, friend := range friends {
		fmt.Printf("- %s \n", friend)
	}
	initArray := [3]int{2, 3, 6}

	fmt.Println(initArray)
	var headquarters [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um número: ")
			fmt.Scanf("%d", &headquarters[i][j])
		}
	}
	fmt.Println(headquarters)
}
