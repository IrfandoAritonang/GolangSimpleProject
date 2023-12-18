// Calculate earnings before tax (EBT) and earnings after tax (profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expense: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := 	revenue - expenses
	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	fmt.Println("ebt: ", ebt)
	fmt.Println("profit: ", profit)
	fmt.Println("ratio: ",ratio)
}