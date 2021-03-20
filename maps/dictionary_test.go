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

		assertStrings(t, err.Error(), want)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("Adding a word to the dictionary", func(t *testing.T) {
		dictionary.Add("newKey", "This is a new value")
		got, _ := dictionary.Search("newKey")
		want := "This is a new value"

		assertStrings(t, got, want)
	})
	t.Run("Adding an already existing word", func(t *testing.T) {
		key := "alreadyExistingKey"
		value := "Already Existing"
		dictionary := Dictionary{key: value}
		got := dictionary.Add(key, value)
		want := "already existing"

		assertStrings(t, got.Error(), want)
	})
}

func TestUpdate(t *testing.T) {
	key := "alreadyExistingKey"
	value := "sample text"
	dictionary := Dictionary{key: value}
	t.Run("Updating an existing key", func(t *testing.T) {
		dictionary.Update(key, "Updated Sample Text")
		got, _ := dictionary.Search(key)
		want := "Updated Sample Text"
		assertStrings(t, got, want)
	})

	t.Run("Updating a key that doesn't exist", func(t *testing.T) {
		err := dictionary.Update("notReallyAKey", "Doesn't really matter, not updating anyway")
		want := "key doesn't exist"
		assertStrings(t, err.Error(), want)
	})
}

func TestDelete(t *testing.T) {
	key := "alreadyExistingKey"
	value := "sample text"
	dictionary := Dictionary{key: value}
	t.Run("Deleting a key", func(t *testing.T) {
		dictionary.Delete(key)
		_, gotError := dictionary.Search(key)
		want := "could not find the word you were searching"

		assertStrings(t, gotError.Error(), want)
	})

	t.Run("Deleting a key that doesn't exist", func(t *testing.T) {
		err := dictionary.Delete("notExistentKey")
		want := "key doesn't exist"

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}

}
