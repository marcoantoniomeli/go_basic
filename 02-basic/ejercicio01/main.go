package main

import (
	"fmt"
)

func main() {

	var (
		palabra string
		count   int
	)

	fmt.Println("Ingresa una palabra")
	fmt.Scanf("%s", &palabra)

	for i := 0; i < len(palabra); i++ {

		temp := string(palabra[i])
		fmt.Println(temp)
		count++
	}

	fmt.Println("Total de letras", count)

}
