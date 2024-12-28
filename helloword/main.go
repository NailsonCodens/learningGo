package main

import "fmt"

func main() {
	/*fmt.Println("Olá mundo!")
	fmt.Println("Treina Web")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("1.1 + 2.2 =", 1.1+2.2)
	fmt.Println(true)

	fmt.Println("Inteiros sem sinal")
	var u1 uint8 = 255 //uint8 could be byte
	fmt.Println(u1)
	var u2 uint16 = 33
	fmt.Println(u2)
	var u3 uint32 = 44
	fmt.Println(u3)
	var u4 uint64 = 55
	fmt.Println(u4)

	fmt.Println("Inteiros com sinal")
	var i1 int8 = 127
	fmt.Println(i1)
	var i2 int16 = 1000
	fmt.Println(i2)
	var i3 int32 = 1001 //int32 could be rune
	fmt.Println(i3)
	var i4 int64 = 1002
	fmt.Println(i4)

	var t1 uint = 1010
	fmt.Println(t1)
	var t2 int = 2000
	fmt.Println(t2)

	fmt.Println("Pontos flutuantes")
	var f1 float32 = 1
	fmt.Println(f1)
	var f2 float64 = 2
	fmt.Println(f2)
	var f3 complex64 = 3
	fmt.Println(f3)
	var f4 complex128 = 4
	fmt.Println(f4)

	fmt.Println("Strings")
	var s1 string = "Treina Web \n Cursos"
	fmt.Println(s1)
	var s2 string = `Treina

	Cursos
	`
	fmt.Println(s2)
	fmt.Println(s1[0])

	fmt.Println("BOOLEANOS")
	var b1 bool = true
	fmt.Println(b1)

	fmt.Println("Multiplas declaracoes")

	//var nome, sobrenome string = "Treina web", "Cursos"
	var (
		nome      string = "Treina Web"
		sobrenome string = "Crusos"
		idade     int    = 10
	)
	fmt.Println(nome)
	fmt.Println(sobrenome)
	fmt.Println(idade)*/

	//inferencia de tipo
	var nome = "treinaweb"
	fmt.Println(nome)

	nomeCompleto := "Treina web Cursos"

	fmt.Println(nomeCompleto)

	const constante string = "Treina Web"
	fmt.Println(constante)

}
