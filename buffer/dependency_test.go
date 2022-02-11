package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Zoe")
	got := buffer.String()
	want := "Hello, Zoe"

	assertCorrectMessage(t, got, want)
}

func TestCountDown(t *testing.T){
	buffer := &bytes.Buffer{}
	spy := &spySleeper{}
	CountDown(buffer, spy)
	got := buffer.String()
	want := `3
2
1
GO!
`
	assertCorrectMessage(t, got, want)
	
	
}