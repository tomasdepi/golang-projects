package decrypt

// functions that start with uppercase are exported, can be accessed from outside
// functions that start with lowercase are internal to the package
func Nimbus(s string) string {
	decryptedStr := ""

	for _, c := range s {
		asciiCode := int(c)
		character := string(asciiCode - 3)
		decryptedStr += character
	}

	return decryptedStr
}
