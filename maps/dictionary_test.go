package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing a map"}
	got := Search(dictionary, "test")
	want := "testing a map"
	assertSearch(t, got, want)
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}

}
