package main

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T){

	t.Run("Say hi to people", func(t *testing.T){
		
		got := SayHi("Peter")
		want := "Hello, Peter"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("default hello", func(t *testing.T){
		got := SayHi("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleAdd(){
	sum := Adder(1, 5)
	fmt.Println(sum)
	// output:6
}

func TestAdder(t *testing.T){
	sum := Adder(2, 2)
	expected := 4
	assertCorrectMessage(t, sum, expected)
	
}

func TestRepeater(t *testing.T){
	got := Repeater("u", 3)
	want := "uuu"
	assertCorrectMessage(t, got, want)
}

/*
基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
*/
func BenchmarkRepeater(b *testing.B){
	for i := 0; i < b.N; i++{
		Repeater("u", 1)
	}
}