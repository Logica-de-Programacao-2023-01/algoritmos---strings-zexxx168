package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, mudar, nova string

	fmt.Println("digite nomes: ")
	fmt.Scanln(&str)

	fmt.Println("digite o nome a ser removido")
	fmt.Scanln(&mudar)

	fmt.Println("digite o nome a ser adicionado: ")
	fmt.Scan(&nova)

	replace := strings.ReplaceAll(str, mudar, nova)

	fmt.Println(replace)
}
