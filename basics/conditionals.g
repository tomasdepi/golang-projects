package main

import "fmt"

const LEGAL_AGE int = 18

func main() {
	var myAge int
	fmt.Print("Enter your age ")
	fmt.Scan(&myAge)

	if myAge >= LEGAL_AGE {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}

	switch myAge {
	case LEGAL_AGE:
		fmt.Println("You Are in the Limit of the Legal Age")
	case LEGAL_AGE - 1:
		fmt.Println("Your Age is 17")
	case LEGAL_AGE + 1:
		fmt.Println("Your Age is 19")
		fallthrough
	default:
		fmt.Println("Default case")

	}

}
