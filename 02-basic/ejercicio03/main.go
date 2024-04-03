package main

import (
	"fmt"
	"strings"
)

/*
Realizar una aplicación que reciba  una variable con el número del mes.

Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
Ej: 7, Julio.
*/

func main() {

	var (
		fecha   string
		flag    bool = false
		flagTwo bool = false
		temp    string
	)

	month := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}

	fmt.Println("Enter the date")
	fmt.Scanf("%s", &fecha)

	dateSplit := strings.Split(fecha, "/")
	monthActua := dateSplit[1]

	for key, value := range month {

		switch {

		case key == dateSplit[1]:
			flag = true
			temp = month[monthActua]

		case value == monthActua:

			temp = monthActua
			flag = true
		}

	}

	if !flag && !flagTwo {
		fmt.Println("Value of Month no valid")
	} else {
		fmt.Println("Month is", temp)
	}

}
