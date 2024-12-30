package main

import "fmt"

func main() {
	fmt.Println("Estou abrindo a conexão com o banco de dados ")
	defer fmt.Println("Estou fechando a conexão com o banco de dados!")
	executSql()
}

func executSql() {
	fmt.Println("Estou executando a consulta ao banco de daods...")
}
