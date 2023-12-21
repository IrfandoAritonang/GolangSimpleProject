package main

import (
	"errors"
	"fmt"
	"os"
	"strconv" //"string conversion" yang digunakan untuk melakukan konversi antara nilai numerik dan representasi stringnya, serta untuk melakukan manipulasi pada string.
)

const accountBalanceFile ="balance.txt"

//membaca data dari file yang sudah ada atau dibuat
//_ untuk meletakkan variabel yang akan datang karena belum digunakan
func getBalanceFromFile() (float64, error) {
	data, err :=os.ReadFile(accountBalanceFile)
	//nil --> null
	// jika file tidak ada maka akan munncul pesan eror
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)//64 merupakan nilai float default 
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

//menyimpan data ke file agar data tetap tersimpan
func writeBalanceToFile(balance float64) {
	//sprint --> gunannya untuk menghasilkan string baru
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)//berikan nilai 0664 sebagai izin
}

func main() {

	//data
	// var accountBalance = 1000.0
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't continue, sorry")//karena memang itu dimaksudkan untuk digunakan dalam kasus-kasus di mana Anda tidak bisa melanjutkan. Ini juga akan mengakhiri eksekusi aplikasi,
	}

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
