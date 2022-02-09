package main

import (
	"fmt"
	"testing"
)

func TestCreateWallet(t *testing.T){
	tests := []struct{
		Name string
		Balance Bitcoin
	}{
		{"Ziv", 80.0},
		{"Zack", 80.0},
		{"Zoe", 80.0},
	}

	for _, test := range(tests){
		CreateWallet(test.Name, test.Balance)
	}
	
	got := fmt.Sprintf("%v", Users) //Let struct turn into string
	want := "[{Ziv 80.000000 BTC} {Zack 80.000000 BTC} {Zoe 80.000000 BTC}]"
	assertCorrectMessage(t, got, want)
}

func TestWallet(t *testing.T){
	
	t.Run("Deposit", func(t *testing.T){
		got := DepositWallet("Ziv", 20.0)
		want := Bitcoin(100.0)
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("Withdraw", func (t *testing.T){
		got := WithdrawWallet("Ziv", 20.0)
		want := Bitcoin(60.0)
		assertCorrectMessage(t, got, want)
	})

	t.Run("Withdraw but not enough money", func(t *testing.T){
		wallet := CreateWallet("Dumb", 20.0)
		err := wallet.Withdraw(99.0)
		assertCorrectMessage(t, wallet.Balance, Bitcoin(20.0))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

