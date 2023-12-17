// Calculate earnings before tax (EBT) and earnings after tax (profit)
// Calculate ratio (EBT / profit)
// Output EBT, profit and the ratio


package buildingprofitcalculator

import "fmt"

func buildingprofitcalculator() {
	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expense: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

}