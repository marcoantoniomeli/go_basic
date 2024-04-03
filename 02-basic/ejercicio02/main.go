package main

import "fmt"

func main() {

	var (
		year   int
		job    bool
		oldJob int
		salary float64
	)

	fmt.Println("age of the client?")
	fmt.Scanf("%d\n", &year)

	fmt.Println("do you have a job?")
	fmt.Scanf("%t\n", &job)

	fmt.Println("seniority at work?")
	fmt.Scanf("%d\n", &oldJob)

	fmt.Println("Salary greater than 100,000?")
	fmt.Scanf("%f\n", &salary)

	switch {

	case year < 22:
		fmt.Println("----------------------------------------")
		fmt.Println("The client is the one under 22 years old")
		fmt.Println("----------------------------------------")
		fallthrough

	case !job:

		fmt.Println("----------------------------------------")
		fmt.Println("The client is the one under 22 years old")
		fmt.Println("----------------------------------------")
		fallthrough

	case oldJob < 1:
		fmt.Println("----------------------------------------------------")
		fmt.Println("The client has been on the job for less than a year.")
		fmt.Println("----------------------------------------------------")
		fallthrough

	case salary < 100000.00:
		fmt.Println("-----------------------------------")
		fmt.Println("the client earns less than 100,000")
		fmt.Println("-----------------------------------")
		fallthrough

	case salary > 100000.00:
		fmt.Println("-------------------------------------------------")
		fmt.Println("congratulations, you will not be charged interest")
		fmt.Println("-------------------------------------------------")

	}

}
