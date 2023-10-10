package main

import "fmt"

// multiple returning function
func operation(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// named return function
func operation_2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// variadic functions, function that accepts variable number of arguments
func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("hey ", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Println(sub)
	}
}

// recursive functions
func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	sum, difference := operation(20, 5)
	fmt.Println(sum, difference)

	fmt.Println(sumNumbers(10, 3, 7))

	printDetails("Joe", "Physics", "Biology")

	fmt.Println(factorial(5))

	// anonymous functions
	x := func(a int, b int) int {
		return a * b
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x(5, 8))
}
