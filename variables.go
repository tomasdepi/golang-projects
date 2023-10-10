package main

import "fmt"

var global_var string = "Depi"

func main() {
	// variables

	// var <variable name> <data type> = <type>
	var s string = "Hello World!!"
	//var i int = 100
	//var b bool = false
	//var f float64 = 465.6246

	// shorthand way (just for same data types)
	var st, t string = "foo", "bar"

	// var (
	// s string = "foo"
	// i int = 5)

	// short variable declaration
	// s := "Hello!"

	fmt.Printf("%s asd\n", s)
	fmt.Println("Hello!!!")
	fmt.Println(st)
	fmt.Println(t)
	fmt.Println(global_var)
}
