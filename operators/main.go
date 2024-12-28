package main

import "fmt"

func main() {
	var numberOne int
	var numbertwo int

	fmt.Print("Digite o primeiro número:")
	fmt.Scanln(&numberOne)
	fmt.Print("Digite o segundo número:")
	fmt.Scanln(&numbertwo)

	fmt.Printf("%d + %d = %d \n", numberOne, numbertwo, numberOne+numbertwo)
	fmt.Printf("%d - %d = %d \n", numberOne, numbertwo, numberOne-numbertwo)
	fmt.Printf("%d x %d = %d \n", numberOne, numbertwo, numberOne*numbertwo)
	fmt.Printf("%d / %d = %d \n", numberOne, numbertwo, numberOne/numbertwo)
	fmt.Printf("%d %% %d = %d \n", numberOne, numbertwo, numberOne%numbertwo)

	increment := numberOne
	increment += numbertwo

	fmt.Printf("O incremento de %d com %d é %d \n", numberOne, numbertwo, increment)

	decrement := numberOne
	decrement += numbertwo

	fmt.Printf("O decremento de %d com %d é %d \n", numberOne, numbertwo, decrement)

}
