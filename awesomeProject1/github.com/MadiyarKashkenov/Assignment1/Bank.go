package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

var bank *BankAccount

func Deposit(amount float64) {
	bank.Balance = bank.Balance + amount
}
func Withdraw(amount float64) {
	if amount > bank.Balance {
		fmt.Println("Incorrect order!")
	} else {
		bank.Balance -= amount
	}
}
func GetBalance() {
	fmt.Println(bank.Balance)
}
func Transaction(account *BankAccount, transactions []float64) {

}
func main() {
	bank = &BankAccount{
		AccountNumber: "00000001",
		HolderName:    "Kashkenov Madiyar",
		Balance:       0.0,
	}
	for {
		var ch int
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get Balance")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		fmt.Scan(&ch)
		var amount float64
		switch ch {
		case 1:
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			Deposit(amount)
			fmt.Println("Deposited:", amount)

		case 2:
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&amount)
			Withdraw(amount)
			fmt.Println("Withdrew:", amount)

		case 3:
			GetBalance()

		case 4:
			return
		}
	}
}
