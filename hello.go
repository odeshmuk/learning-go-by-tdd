package main

import "fmt"

func hello() string {
	return "Hello World"
}

const englishHelloPrefix = "Hello "

func helloName(name string) string {
	if name == "" {
		name = "Stranger"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(hello())
}