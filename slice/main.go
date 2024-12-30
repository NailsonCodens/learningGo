package main

import "fmt"

func main() {
	friends := make([]string, 3)
	name := ""
	i := 0

	for name != "q" {
		fmt.Print("Digite o nome de um amigo ou 'q' para parar: ")
		fmt.Scanf("%s", &name)

		if name != "q" {
			if i < 3 {
				friends[i] = name
			} else {
				friends = append(friends, name)
			}
			i++
		}
	}

	fmt.Println(friends)
	fmt.Printf("Você tem %d amigos \n", len(friends))
	for _, friendName := range friends {
		fmt.Printf(" - %s \n", friendName)
	}

	anotherSlice := friends[1:3]
	fmt.Println(anotherSlice)

	oneMoreSlice := friends[:3]
	fmt.Println(oneMoreSlice)
}
