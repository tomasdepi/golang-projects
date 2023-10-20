package main

import (
	"fmt"

	"github.com/tomasdepi/golang-projects/modules/cryptic/decrypt"
	"github.com/tomasdepi/golang-projects/modules/cryptic/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Hola Mundo")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
