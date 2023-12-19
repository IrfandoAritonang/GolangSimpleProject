package main

import "fmt"

func main() {

	//data
	var accountBalance = 1000.0
 
	//akhir//perulangan for --> agar program tidak berhenti karena 1 perintah terpenuhi
	for i := 0; i < 200; i++ {
		fmt.Println("Welcone to Bank BRI :)")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Whitdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)


	if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice ==2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			//cek agar tidak ada jumlah negatif yg disetorkan atau ditarik
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return 
			}
			accountBalance += depositAmount// accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New Amount: ", accountBalance)
		}else	if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var WithdrawalAmount float64
			fmt.Scan(&WithdrawalAmount)
			//cek agar tidak ada jumlah negatif yg disetorkan atau ditarik
			if WithdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				return 
			}
			
			if WithdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				return
			}

			accountBalance -= WithdrawalAmount// accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
		}else {
			fmt.Println("Thanks and Goodbye !")
		}
	}
}
