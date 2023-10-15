package main

import "fmt"

func main() {
	// unordered collection of key/value pairs
	// implemented by hash tables
	// provide efficient add, get and delete operations

	codes := map[string]string{"en": "English", "hi": "Hindi", "es": "Espaniol"}
	codes["fr"] = "French"
	fmt.Println(codes)
	fmt.Println(len(codes))

	codes_1 := make(map[string]string) // creates empty map
	fmt.Println(codes_1)

	value, found := codes["en"]
	fmt.Println(value, found)
	value, found = codes["hh"]
	fmt.Println(value, found)

	delete(codes, "fr")
	fmt.Println(codes)

	for key, value := range codes {
		fmt.Println(key, "=>", value)
	}

}
