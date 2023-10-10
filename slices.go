package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	sub_slice := slice_1[0:3]
	fmt.Println(sub_slice)

	// make([]<data_type>, lenth, capacity)
	slice_2 := make([]int, 5, 10)
	fmt.Println(slice_2)
	fmt.Println(len(slice_2))
	fmt.Println(cap(slice_2))

	slice_1[0] = 1
	fmt.Println(arr)
	// second item of arr is 1 now since a slice is a pointer to the underlying array
}
