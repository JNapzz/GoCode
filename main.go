// main.go
package main

import (
	"fmt"
	"log"
)

func main() {
	// Specify the file name
	filename := "employees.csv"

	// Data to write to the CSV file
	employeesToWrite := []Employee{
		{"John Doe", 30, 55000.50},
		{"Alice Smith", 25, 60000.00},
		{"Bob Johnson", 40, 45000.00},
		{"Charlie Brown", 35, 70000.00},
	}

	// Writing data to CSV
	err := writeToCSV(filename, employeesToWrite)
	if err != nil {
		log.Fatalf("Error writing to CSV: %v", err)
	}

	// Reading data from CSV
	employees, err := readFromCSV(filename)
	if err != nil {
		log.Fatalf("Error reading from CSV: %v", err)
	}

	// Sorting employees by name
	SortEmployeesByName(employees)

	// Display sorted employees
	fmt.Println("Employees sorted by name:")
	for _, emp := range employees {
		fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n", emp.Name, emp.Age, emp.Salary)
	}
}
