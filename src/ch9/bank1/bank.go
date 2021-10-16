// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)  // send amount to deposit
var balances = make(chan int)  // receive balance
var withdraws = make(chan int) // send amount to withdraw

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	withdraws <- amount
	if amount1 := <-withdraws; amount1 == amount {
		return true
	} else{
		return false
	}
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance < amount {
				withdraws <- -1 * amount
			} else {
				withdraws <- amount
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
