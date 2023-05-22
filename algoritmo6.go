package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	fmt.Println("digite um texto")
	fmt.Scanln(&texto)

	nova := len(strings.Fields(texto))

	fmt.Println(nova)

}
