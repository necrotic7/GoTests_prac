package main

import (
	
	"reflect"
	"testing"
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

func assertError (t *testing.T, got error, want string){
	t.Helper()
	if got == nil {
		t.Error("Want an error but got nil")
	}else if got.Error() != want{
		t.Errorf("want '%s', but got '%s'", want, got)
	}
}

func SayHi(name string) string{
	if name == "" {
		name = "World"
	}
	return "Hello, "+ name
}

func Adder(x, y int)int{
	return x+y
}

func Repeater(character string, times int)string{
	var result string
	for i := 0 ; i < times ; i++{
		result += character
	}
	return result
}

func Summer(numbers []int)int{
	var sum int
	for _, numbers := range numbers{
		sum += numbers
	}
	return sum
}

func SumAll(numbers ...[]int)[]int{
	lenOfNumbers := len(numbers)

	sums := make([]int, lenOfNumbers)
	for i, number := range(numbers){
		sums[i] = Summer(number)
		// sums = append(sums, Sum(number))
	}
	return sums
}

func CollectAllTail(numbers ...[]int)[]int{
	var sums []int
	for _, number := range(numbers){
		if len(number) == 0{
			continue
		}else{
			tail := number[len(number)-1]
			sums = append(sums, tail)
		}
	}
	return sums
}

func CreateWallet(name string, balance Bitcoin) Wallet{
	wallet := Wallet{Name: name, Balance: balance}
	Users = append(Users, wallet)
	
	return wallet
}

func DepositWallet(name string, money Bitcoin) Bitcoin{
	var wallet Wallet
	for i, user := range(Users){
		if user.Name == "Ziv"{
			wallet = Users[i]
		}
	}
	wallet.Deposit(money)
	return wallet.Balance
}

func WithdrawWallet(name string, money Bitcoin) Bitcoin{
	var wallet Wallet
	for i, user := range(Users){
		if user.Name == "Ziv"{
			wallet = Users[i]
		}
	}
	wallet.Withdraw(money)
	return wallet.Balance
}


func main(){
}