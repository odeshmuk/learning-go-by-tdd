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

func TestHelloMultiple(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testing with a regular name", func(t *testing.T) {
		got := helloName("Omkar")
		want := "Hello Omkar"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing with an empty name", func(t *testing.T) {
		got := helloName("")
		want := "Hello Stranger"
		assertCorrectMessage(t, got, want)
	})

}
