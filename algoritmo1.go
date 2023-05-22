package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string

	fmt.Println("escreva um nome")
	fmt.Scanln(&nome)

	maiusculo := strings.ToUpper(nome)
	fmt.Println("seu nome em maisculo: ", maiusculo)
}
