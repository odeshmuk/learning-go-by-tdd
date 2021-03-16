package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing a map"}
	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "testing a map"
		assertStrings(t, got, want)
		if err != nil {
			t.Fatal("Expected no error")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were searching"

		if err == nil {
			t.Fatal("expected an error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}

}
