package main

import "fmt"

func main() {
	var str1, str2 string

	fmt.Println("digite um nome: ")
	fmt.Scan(&str1)

	fmt.Println("digite outro nome: ")
	fmt.Scan(&str2)

	if str1 == str2 {
		fmt.Println("os nomes sao iguais")
	} else {
		fmt.Println("os nome sao diferente")
	}

}
