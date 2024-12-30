package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	courseMap := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	course := ""

	for course != "q" {
		var workload int
		fmt.Print("Digite um curso ou digite q para sair: ")
		scanner.Scan()
		course = scanner.Text()

		if course != "q" {
			fmt.Print("Digite a carga horária do curso: ")
			fmt.Scanf("%d", &workload)
			courseMap[course] = workload
		}
	}

	fmt.Println("Cursos registrado: ")
	for courseName, workLoad := range courseMap {
		fmt.Printf("- %s: %dh \n", courseName, workLoad)
	}

	course = ""

	for course != "q" {
		fmt.Print("Digite o nome do curso a ser excluído ou digite q para cancelar: ")
		scanner.Scan()
		course = scanner.Text()
		if course != "q" {
			workLoad, courseAlreadyExists := courseMap[course]
			if courseAlreadyExists {
				delete(courseMap, course)
				fmt.Printf("O curso %s com %dh foi removido \n", course, workLoad)
			} else {
				fmt.Println("O curso digitado não existe")
			}

		}
	}
	fmt.Println("Cursos registrado: ")
	for courseName, workLoad := range courseMap {
		fmt.Printf("- %s: %dh \n", courseName, workLoad)
	}
}
