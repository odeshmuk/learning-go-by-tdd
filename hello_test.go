package main

import "testing"

//The test function must start with the word Test
func TestHello(t *testing.T) {
	got := hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := helloName("Omkar")
	want := "Hello Omkar"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
