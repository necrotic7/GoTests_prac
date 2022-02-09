package main

import (
	"testing"
)

func TestCreateWallet(t *testing.T){
	got := CreateWallet("Ziv", 80.0)
	want := Wallet{"Ziv", 80.0}
	assertCorrectMessage(t, got, want)
}

func TestWallet(t *testing.T){

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
}

