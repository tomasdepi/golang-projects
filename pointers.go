package main

import "fmt"

func main() {
	x := 80
	address_x := &x
	x_2 := *address_x
	fmt.Printf("The variable x it's located at %v and the data type is %T \n", address_x, address_x)
	fmt.Println(x_2)

	// zero value of a pointer is nil

	s := "hello"
	fmt.Printf("%T %v \n", s, s)
	ps := &s
	*ps = "world"
	fmt.Printf("%T %v \n", s, s)

}
