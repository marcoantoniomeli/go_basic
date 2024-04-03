package main

import "fmt"

func main() {

	var (
		firstName      string = " Marco "
		lastName       string = " Mena"
		age            int    = 17
		driver_license bool   = true
		person_height  int    = 33
	)

	childsNumber := 2

	fmt.Printf("First Name %s Last Name %s \n", firstName, lastName)
	fmt.Printf("Years %d \n", age)
	fmt.Printf("Has driver license ? %t \n", driver_license)
	fmt.Printf("Height person? %d \n", person_height)
	fmt.Printf("Numbre of childs %d \n", childsNumber)

}
