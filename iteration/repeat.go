package iteration

// Repeat string passed
func Repeat(repeatingString string, repeatBy int) string {
	var repeatedChar string
	for i := 0; i < repeatBy; i++ {
		repeatedChar += repeatingString
	}
	return repeatedChar
}
