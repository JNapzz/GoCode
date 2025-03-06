package main

import (
	"fmt"
	"budget-app/budgetapp"
)

func main() {
	// Create a new budget app instance
	app := budgetapp.BudgetApp{}

	// Load existing transactions from the file
	err := app.LoadTransactions("transactions.txt")
	if err != nil {
		fmt.Println("Error loading transactions:", err)
	}

	// Loop to display menu and get user input
	for {
		// Display Menu
		fmt.Println("\n--- Budget App Menu ---")
		fmt.Println("1. Add Income")
		fmt.Println("2. Add Expense")
		fmt.Println("3. View Summary")
		fmt.Println("4. View All Incomes")
		fmt.Println("5. View All Expenses")
		fmt.Println("6. Edit Transaction")
		fmt.Println("7. Delete Transaction")
		fmt.Println("8. Export to CSV")
		fmt.Println("9. View Monthly Summary")
		fmt.Println("10. Exit")
		fmt.Print("Please choose an option (1-10): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Add Income
			var incomeAmount float64
			fmt.Print("Enter income amount: $")
			fmt.Scan(&incomeAmount)
			app.AddTransaction("Income", incomeAmount)
			fmt.Printf("Income of $%.2f added.\n", incomeAmount)
		case 2:
			// Add Expense
			var expenseAmount float64
			fmt.Print("Enter expense amount: $")
			fmt.Scan(&expenseAmount)
			app.AddTransaction("Expense", expenseAmount)
			fmt.Printf("Expense of $%.2f added.\n", expenseAmount)
		case 3:
			// Show Budget Summary
			app.ShowSummary()
		case 4:
			// View All Incomes
			app.ViewTransactions("Income")
		case 5:
			// View All Expenses
			app.ViewTransactions("Expense")
		case 6:
			// Edit Transaction
			app.EditTransaction()
		case 7:
			// Delete Transaction
			app.DeleteTransaction()
		case 8:
			// Export to CSV
			err := app.ExportToCSV("transactions.csv")
			if err != nil {
				fmt.Println("Error exporting to CSV:", err)
			} else {
				fmt.Println("Transactions exported to transactions.csv")
			}
		case 9:
			// View Monthly Summary
			app.ViewMonthlySummary()
		case 10:
			// Save transactions to the file before exiting
			err := app.SaveTransactions("transactions.txt")
			if err != nil {
				fmt.Println("Error saving transactions:", err)
			}
			fmt.Println("Exiting the budget app. Goodbye!")
			return
		default:
			// Invalid choice
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
