package main

import (
	"fmt"
	"os"
)

//menyimpan data ke file agar data tetap tersimpan
func writeBalanceToFile(balance float64) {
	//sprint --> gunannya untuk menghasilkan string baru
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)//berikan nilai 0664 sebagai izin
}

func main() {

	//data
	var accountBalance = 1000.0

	fmt.Println("Welcome to Bank BRI :)")
 
	//akhir//perulangan for --> agar program tidak berhenti karena 1 perintah terpenuhi
	//dan buat program ini dijalankan tanpa batas/limit
	for {
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Whitdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)


	switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
			break
		case 2:
			fmt.Print("Your Deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)
				//cek agar tidak ada jumlah negatif yg disetorkan atau ditarik
				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					//return 
					continue
				}
				accountBalance += depositAmount// accountBalance = accountBalance + depositAmount
				fmt.Println("Balance updated! New Amount: ", accountBalance)
				writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var WithdrawalAmount float64
			fmt.Scan(&WithdrawalAmount)
			//cek agar tidak ada jumlah negatif yg disetorkan atau ditarik
			if WithdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue 
			}
			
			if WithdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= WithdrawalAmount// accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye !")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
