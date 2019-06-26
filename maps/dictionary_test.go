package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		d := Dictionary{"test": "this is just test"}

		got, err := d.Search("test")
		want := "this is just test"

		assertString(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		d := Dictionary{"test": "this is just test"}

		_, err := d.Search("huh?")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this just a test"
		d.Add(word, definition)
		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}

		otherDef := "this is a nother definition"
		err := d.Add(word, otherDef)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: definition}
		newDef := "new definition"

		err := dict.Update(word, newDef)
		assertNoError(t, err)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}
	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted: %s", word, err)
	}
}

func assertDefinition(t *testing.T, d Dictionary, word, defintion string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if defintion != got {
		t.Errorf("got '%v', want '%v, given '%s'", got, defintion, word)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s', want '%s', given '%s'", got, want, "test")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("got %v, want nil", err)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatalf("didn't got error, but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
