package budgetapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"time"
	"budget-app/transaction"
	"strconv"
)

// Define the structure for the Budget App
type BudgetApp struct {
	Balance      float64
	Transactions []transaction.Transaction
}

// Method to Add a transaction (either Income or Expense)
func (app *BudgetApp) AddTransaction(transactionType string, amount float64) {
	transaction := transaction.Transaction{
		Type:   transactionType,
		Amount: amount,
	}

	// Add transaction to the list
	app.Transactions = append(app.Transactions, transaction)

	// Update the balance
	if strings.ToLower(transactionType) == "income" {
		app.Balance += amount
	} else if strings.ToLower(transactionType) == "expense" {
		app.Balance -= amount
	}
}

// Method to display the current budget summary (Balance and Transactions)
func (app *BudgetApp) ShowSummary() {
	fmt.Printf("\n--- Budget Summary ---\n")
	fmt.Printf("Current Balance: $%.2f\n", app.Balance)
	fmt.Println("Transactions:")
	for i, transaction := range app.Transactions {
		fmt.Printf("%d. %s: $%.2f\n", i+1, transaction.Type, transaction.Amount)
	}
	fmt.Println("----------------------")
}

// Method to save transactions to a file
func (app *BudgetApp) SaveTransactions(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, transaction := range app.Transactions {
		line := fmt.Sprintf("%s,%f\n", transaction.Type, transaction.Amount)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

// Method to load transactions from a file
func (app *BudgetApp) LoadTransactions(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		amount, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return err
		}
		app.AddTransaction(parts[0], amount)
	}
	return scanner.Err()
}

// Method to view all transactions of a specific type (Income or Expense)
func (app *BudgetApp) ViewTransactions(transactionType string) {
	fmt.Printf("\n--- View %s ---\n", transactionType)
	for i, transaction := range app.Transactions {
		if strings.EqualFold(transaction.Type, transactionType) {
			fmt.Printf("%d. %s: $%.2f\n", i+1, transaction.Type, transaction.Amount)
		}
	}
	fmt.Println("------------------")
}

// Method to edit an existing transaction
func (app *BudgetApp) EditTransaction() {
	fmt.Println("Choose the transaction number to edit:")
	var num int
	fmt.Scan(&num)

	if num < 1 || num > len(app.Transactions) {
		fmt.Println("Invalid transaction number.")
		return
	}

	transaction := &app.Transactions[num-1]
	fmt.Printf("Editing transaction #%d: %s $%.2f\n", num, transaction.Type, transaction.Amount)
	
	var newType string
	var newAmount float64
	fmt.Print("Enter new type (Income/Expense): ")
	fmt.Scan(&newType)
	fmt.Print("Enter new amount: $")
	fmt.Scan(&newAmount)
	
	transaction.Type = newType
	transaction.Amount = newAmount

	// Update the balance
	app.UpdateBalance()
	fmt.Printf("Transaction #%d updated to %s $%.2f\n", num, transaction.Type, transaction.Amount)
}

// Method to delete a transaction
func (app *BudgetApp) DeleteTransaction() {
	fmt.Println("Choose the transaction number to delete:")
	var num int
	fmt.Scan(&num)

	if num < 1 || num > len(app.Transactions) {
		fmt.Println("Invalid transaction number.")
		return
	}

	app.Transactions = append(app.Transactions[:num-1], app.Transactions[num:]...)
	fmt.Println("Transaction deleted.")
	app.UpdateBalance()
}

// Method to update the balance after editing or deleting transactions
func (app *BudgetApp) UpdateBalance() {
	app.Balance = 0
	for _, transaction := range app.Transactions {
		if strings.ToLower(transaction.Type) == "income" {
			app.Balance += transaction.Amount
		} else if strings.ToLower(transaction.Type) == "expense" {
			app.Balance -= transaction.Amount
		}
	}
}

// Method to export transactions to a CSV file
func (app *BudgetApp) ExportToCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("Type,Amount\n")
	if err != nil {
		return err
	}

	for _, transaction := range app.Transactions {
		line := fmt.Sprintf("%s,%.2f\n", transaction.Type, transaction.Amount)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

// Method to view a monthly summary of transactions
func (app *BudgetApp) ViewMonthlySummary() {
	//currentMonth := time.Now().Month()
	incomeTotal := 0.0
	expenseTotal := 0.0

	for _, transaction := range app.Transactions {
		// Example: Assume each transaction is in the current month (for simplicity)
		if transaction.Type == "Income" {
			incomeTotal += transaction.Amount
		} else if transaction.Type == "Expense" {
			expenseTotal += transaction.Amount
		}
	}

	fmt.Printf("\n--- Monthly Summary ---\n")
	fmt.Printf("Total Income: $%.2f\n", incomeTotal)
	fmt.Printf("Total Expense: $%.2f\n", expenseTotal)
	fmt.Printf("Net Balance: $%.2f\n", incomeTotal-expenseTotal)
	fmt.Println("------------------------")
}
