package main

import "fmt"

func hello() string {
	return "Hello World"
}

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func helloName(name string, language string) string {
	var prefix = englishHelloPrefix
	if name == "" {
		name = "Stranger"
	}
	if language == "Spanish" {
		prefix = spanishHelloPrefix
	}
	if language == "French" {
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(hello())
}
