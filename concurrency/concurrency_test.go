package main

import (
	"fmt"
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
	case map[string]bool:
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v",  got, want)
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

func mockWebsiteChecker(url string) bool{
	if url == "waat://furhurterwe.geds"{
		return false
	}
	return true
}

func SlowWebsiteChecker(_ string)bool{
	time.Sleep(20*time.Second)
	return true
}

func TestCheckWebsites(t *testing.T){
	websites := []string{
        "http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
    }

	ActualResults := CheckWebsites(mockWebsiteChecker, websites)
	got := len(ActualResults)
	want := len(websites)

	assertCorrectMessage(t, got, want)

	ExpectedResults := map[string]bool{
        "http://google.com":          true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds":    false,
    }

	assertCorrectMessage(t, ActualResults, ExpectedResults)
	
}

func BechmarkCheckWebsites(b *testing.B){
	urls := make([]string, 100)
	for i:=0 ; i<len(urls) ; i++{
		urls[i] = "a.com"
	}
	fmt.Println(len(urls))
	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowWebsiteChecker, urls)
	}
}