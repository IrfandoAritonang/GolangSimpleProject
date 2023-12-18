package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

 revenue :=  getUserInput("Revenue: ")
 expenses :=  getUserInput("Expense: ")
 taxRate :=  getUserInput("Tax Rate: ")

 ebt, profit, ratio := Calculation(revenue, expenses, taxRate)



	// fmt.Println("ebt: ", ebt)
	// fmt.Println("profit: ", profit)
	// fmt.Println("ratio: ",ratio)
	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3fn", ratio)

}


//implementasi function
//ubah ke float64 karena ada bilangan desimal
func getUserInput(catchInput string) float64  {
	var userInput float64
	fmt.Print(catchInput)
	fmt.Scan(&userInput)
	return userInput
}


//isi paramaternya dan tipe data nya
//karena mau mengembalikan float tambahkan () untuk menerima nilai kembalian/ouput
//setelah itu tambahan return untuk mencetak
func Calculation(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := 	revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
