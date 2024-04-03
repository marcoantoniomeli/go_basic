package main

import "fmt"

func main() {

	var (
		lastName string = "Smith"

		age    int     = 35
		salary float64 = 45857.90

		firstName string = "Mary"
	)

	boolean := false

	fmt.Printf("Last Name: %s , First Name: %s , Age: %d, Salary: %f \n", lastName, firstName, age, salary)
	fmt.Printf("Boolean: %t \n", boolean)

}
