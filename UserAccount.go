package main

import (
	"fmt"
	"log"
	"os"
)

//User Defined Account Structure
type AccountUser struct {
	VIPuser     string
	DefaultUser string
	AccountID   string
	balance     float64
	SecretPin   int
}

//function initiate a new account process
func newAccount(name string) AccountUser {
	Account := AccountUser{
		AccountID: name,
		balance:   0,
	}
	return Account
}

//Debit process using receiver function
func (account *AccountUser) subAmt(amount float64) {
	account.balance -= amount
	fmt.Println("You have successfully debited your account with", amount)
	fmt.Println("Your new account balance is ", account.balance)
}

//Credit process using receiver function
func (account *AccountUser) addAmt(amount float64) {
	account.balance += amount
	fmt.Println("You have successfully credited your account with", amount)
	fmt.Println("Your account balance is ", account.balance)
}

//Account User's details or info.
func (account *AccountUser) format() string {
	fs := ""
	switch account.AccountID {
	case "":
		fs = "\nYou don't have an account yet\n"
	default:
		fs += fmt.Sprintf("%v: %v \n", "Name", account.AccountID)
		fs += fmt.Sprintf("%v: %v \n", "Account balance", account.balance)
	}
	return fs
}

//receiver function on accountholder to save accountHolder information to txt file
func (account *AccountUser) save() {
	data := []byte(account.format()) //turns it into a byte slice to be stored in the variable

	err := os.WriteFile("C:/Users/HP/Desktop/AccountLog/accountlog"+account.AccountID+".txt", data, 0644)
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
	fmt.Println("Account Information was saved successfully")
}
