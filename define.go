package main

import (
	"errors"
	"fmt"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct{
	Height float64
	Width float64
}

func (r Rectangle) Perimeter() float64{
	return (r.Height + r.Width)*2
}

func (r Rectangle) Area()float64{
	return r.Height*r.Width
}

type Circle struct{
	Radius float64
}

func (c Circle) Perimeter() float64{
	return (c.Radius)*2*3.14
}

func (c Circle) Area() float64{
	return c.Radius*c.Radius*3.14
}

type Triangle struct{
	Height float64
	Side float64
}

func (tri Triangle) Perimeter() float64{
	return tri.Side*3
}

func (tri Triangle) Area() float64{
	return tri.Height*tri.Side/2
}

type Wallet struct{
	Name string
	Balance Bitcoin
}

var Users []Wallet
type Bitcoin float64

//Wallet Method
func (w *Wallet) Deposit(money Bitcoin) Bitcoin{
	w.Balance += money
	return w.Balance
}

func (w *Wallet) Withdraw(money Bitcoin) error{
	if money > w.Balance{
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.Balance -= money
	return nil
}

type Stringer interface{
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%f BTC", b)
}

type Dictionary map[string]string

func (d Dictionary)Search(word string) string{
	if d[word] != ""{
		return d[word]
	}
	return "Not found"
}

func (d Dictionary)Add(key , value string){
	d[key] = value
}