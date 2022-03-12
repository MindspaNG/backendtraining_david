package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Gets input from user
func getInput(PromptOption string, r *bufio.Reader) (string, error) {
	fmt.Print(PromptOption)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

//function to create an account for a user
func CreateAccountUser(user map[int64]string) AccountUser {
	reader := bufio.NewReader(os.Stdin)
	Option, _ := getInput("\nWhich of the type of account do you prefer(1 = Default , 2 = VIP and 3 = Exit): ", reader)
	var account AccountUser

	//decisiom-taking process whether the user meets the requirements
	//the type of account to create

	switch Option {
	case "1":
		fmt.Println("You chose option 1 for a Default Account type")
		amount, _ := getInput("\nYou are required to deposit at least 10000 to get access to creating an account: ", reader)

		amt, err := strconv.ParseFloat(amount, 64)

		if err != nil {
			fmt.Println("Deposit must be an integer value")
			CreateAccountUser(user)
		} else if amt < 10000 {
			fmt.Println("Deposit at least an amount of 10000 to continue")
			CreateAccountUser(user)
		} else {
			userName, _ := getInput("Input an Account ID:", reader)
			account = newAccount(userName)
			account.balance += amt

			fmt.Println("\nCongrats, Your account creation process was successful")
			account.save()
		}

	case "2":
		fmt.Println("You chose option 2 for a VIP Account type")
		userName := accountValidation(user)

		switch userName {
		case "":
			account.AccountID = ""
		default:
			account.AccountID = userName
			account.balance += 50000
			fmt.Println(account.format())
			account.save()
			fmt.Println()
		}

	case "3":
		fmt.Println("\nProcess terminated")
		fmt.Println()

	}
	return account
}

// Validation process concerning the VIP user occurs here
//accID = account user name , pin = Secret pin
func accountValidation(user map[int64]string) string {

	reader := bufio.NewReader(os.Stdin)
	SecretPin, _ := getInput("\nInput your SecretPin: ", reader)
	pin, err := strconv.ParseInt(SecretPin, 10, 64)
	accID, available := user[pin]
	VipUser := ""

	if err != nil {
		fmt.Println("An integer value is required for the process to be in progress")
		accountValidation(user)
	} else {
		switch available {
		case true:
			fmt.Println("The account ID :", accID, "is valid")
			VipUser = accID
		case false:
			fmt.Println("The Secret Pin is not found")
		}
	}
	return VipUser
}

// function to deposit into account
func deposit(account *AccountUser) {

	reader := bufio.NewReader(os.Stdin)
	DepoWith := AccDecision("Do you want to make a deposit?")

	switch DepoWith {
	case true:
		Amount, _ := getInput("\nEnter Amount to Deposit: ", reader)
		amt, err := strconv.ParseFloat(Amount, 64)
		if err != nil {
			fmt.Println("Deposit amount must be a number!")
			deposit(account)
		} else {
			account.addAmt(amt)
			account.save()
		}
	case false:
		fmt.Println("Process Terminated")
	}
}

// function to withdraw from account
func withdraw(account *AccountUser) {

	reader := bufio.NewReader(os.Stdin)
	DepoWith := AccDecision("Do you want to make a withdrawal?")

	switch DepoWith {
	case true:
		Amount, _ := getInput("Input Amount to Withdraw: ", reader)
		amt, err := strconv.ParseFloat(Amount, 64)
		if err != nil {
			fmt.Println("Input a an integer value for withdrawal")
			withdraw(account)
		} else if amt > account.balance {
			fmt.Println("You have an insufficient balance  ", account.balance)
			withdraw(account)
			account.save()
		} else {
			account.subAmt(amt)
			account.save()
		}
	case false:
		fmt.Println("Process Terminated")
	}
}

//function to get Yes or No input from console
func AccDecision(msg string) bool {
	reader := bufio.NewReader(os.Stdin)
	Decis, _ := getInput(msg+"1 - Yes,  0 - No  :", reader)
	Dn := strings.ToUpper(Decis)
	DepoAmt := true

	switch Dn {
	case "1":
		DepoAmt = true
	case "0":
		DepoAmt = false
	default:
		fmt.Println()
	}
	return DepoAmt

}
func main() {
	//Pre-defined V.I.P User with corresponding PIN
	SecretPin := map[int64]string{
		1421: "Audrey Allen",
		2072: "Keren Davis",
		3421: "Mylves Ken",
	}

	// Create new accounts, verify PIN before creating VIP user
	UserAccount := CreateAccountUser(SecretPin)
	fmt.Println(UserAccount.format())
	withdraw(&UserAccount)
	deposit(&UserAccount)
	fmt.Println(UserAccount.format())

}
