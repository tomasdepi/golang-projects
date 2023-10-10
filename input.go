package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	count, err := fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Hello %s, your age is %d\n", name, age)
	fmt.Println("Number of inputs: ", count)
	fmt.Println("Error: ", err)

	fmt.Printf("The variable you entered is of type %T\n", name)
	fmt.Printf("The age is type %v\n", reflect.TypeOf(age))
}
