package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func assertCorrectMessage(t *testing.T, got, want interface{}){
	t.Helper()
	switch v := got.(type){
	case []int:
	if !reflect.DeepEqual(got, want){
			t.Errorf("type%v,got %v want %v", v, got, want)
		}
	default:
		check := got != want
		if check {
			t.Errorf("want '%v', but got '%v'", want, got)
		}
	}
	
}

func assertError(t *testing.T, got error, want string){
	t.Helper()
	if got == nil {
		t.Error("Want an error but got nil")
	}else if got.Error() != want{
		t.Errorf("want '%s', but got '%s'", want, got)
	}
}

func Greet(writer io.Writer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
	fmt.Println()
}

type Sleeper interface{
	Sleep()
}

type spySleeper struct{
	duration time.Duration
}

func (s *spySleeper)Sleep(){
	time.Sleep(s.duration)//1*time.Second
}

func CountDown(writer io.Writer, spy Sleeper){
	for i :=3; i>0 ; i--{
		fmt.Fprintln(writer, i)
		spy.Sleep()
	}
	
	fmt.Fprintln(writer, "GO!")
}

func main(){
	Greet(os.Stdout, "Ziv")
	spy := spySleeper{time.Second}
	CountDown(os.Stdout, &spy)
}