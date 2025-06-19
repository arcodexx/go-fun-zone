package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	
	var balance int = 1000
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Menu :");
		fmt.Println("1. View Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		choiceString, _ := reader.ReadString('\n')
		choice,_ := strconv.Atoi(strings.TrimSpace(choiceString))

		if choice < 1 || choice > 4 {
			fmt.Println("Invalid choice. Please select a valid option (1-4).")
			continue
		} else if choice == 4 {
			fmt.Println("Exiting the program. Thank you!")
			break
		}

		switch choice {
		case 1:
			fmt.Println("Your current balance is: ₹", balance)
		case 2:
			fmt.Println("Enter amount to deposit:")
			amountString, _ := reader.ReadString('\n')
			amount, err := strconv.Atoi(strings.TrimSpace(amountString))
			if err != nil || amount <= 0 {
				fmt.Println("Invalid deposit amount. Please enter a positive number.")
			} else {
				balance += amount
				fmt.Println("Successfully deposited ₹", amount, ". New balance is: ₹", balance)
			}
		case 3:
			if( balance == 0) {
				fmt.Println("Your balance is zero. You cannot withdraw money.")
				continue
			}
			fmt.Println("Enter amount to withdraw:")
			amountString, _ := reader.ReadString('\n')
			amount, err := strconv.Atoi(strings.TrimSpace(amountString))
			if err != nil || amount <= 0 {
				fmt.Println("Invalid withdrawal amount. Please enter a positive number.")
			} else if amount > balance {
				fmt.Println("Insufficient balance. You cannot withdraw more than your current balance of ₹", balance)
			} else {
				balance -= amount
				fmt.Println("Successfully withdrew ₹", amount, ". New balance is: ₹", balance)
			}
		}
	}
}