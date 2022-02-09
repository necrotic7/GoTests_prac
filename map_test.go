package main

import (
	"testing"
)

func TestSearch(t *testing.T){
	t.Run("known word", func(t *testing.T){
		dict := Dictionary{"test": "Just a test"}
		got, _ := dict.Search("test")
		want := "Just a test"

		assertCorrectMessage(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T){
		dict := Dictionary{"test": "Just a test"}
		_, got := dict.Search("haha")
		want := "key not found"

		assertError(t, got, want)
	})
	
}

func TestAdd(t *testing.T){
	t.Run("Normal Add", func(t *testing.T){
		dict := Dictionary{}
		dict.Add("test2", "Adding test")
		got, _ := dict.Search("test2")
		want := "Adding test"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Adding exist word", func(t *testing.T){
		dict := Dictionary{"test3": "exist"}
		err := dict.Add("test3", "double")

		assertError(t, err, "key already exist")
	})
}

func TestUpdate(t *testing.T){
	t.Run("normal update", func(t *testing.T){
		dict := Dictionary{"test4":"original value"}
		dict.Update("test4", "new value")
		got, _:=dict.Search("test4")
		want := "new value"

		assertCorrectMessage(t, got, want)
	})
}

func TestDelete(t *testing.T){
	t.Run("Normal Delete", func(t *testing.T){
		dict := Dictionary{"test5":"Delete test"}
		dict.Delete("test5")
		_, err := dict.Search("test5")
		want := "key not found"
		assertError(t, err, want)
		
	})

	t.Run("Delete not found", func(t *testing.T){
		dict := Dictionary{}
		got := dict.Delete("test6")
		want := "key not found"
		assertError(t, got, want)
	})
}