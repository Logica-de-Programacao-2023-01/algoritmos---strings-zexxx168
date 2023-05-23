package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string

	fmt.Println("digite um texto")
	fmt.Scanln(&texto)

	nova1 := strings.Fields(texto)
	nova2 := len(nova1)

	fmt.Println(nova2)

}
