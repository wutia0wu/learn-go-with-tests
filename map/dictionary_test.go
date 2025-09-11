package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("know word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("abcd")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want.Error())
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefintion(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		// err := dictionary.Add(word, "existing word")
		err := dictionary.Add(word, "new word")

		assertError(t, err, ErrWordExists)
		assertDefintion(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefintion(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted ", word)
	}
}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertDefintion(t *testing.T, dic Dictionary, word, definition string) {
	t.Helper()

	got, err := dic.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertStrings(t, got, definition)
}

func assertError(t *testing.T, addErr, typeErr error) {
	if addErr != typeErr {
		t.Errorf("(%s) should be (%s)", addErr, typeErr)
	}
}
