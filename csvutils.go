// csvutils.go
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// writeToCSV writes a list of employees to a CSV file.
func writeToCSV(filename string, data []Employee) error {
	// Open the file in append mode, create it if it doesn't exist.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open or create the file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header if the file is empty
	if fileInfo, _ := file.Stat(); fileInfo.Size() == 0 {
		writer.Write([]string{"Name", "Age", "Salary"})
	}

	// Write each employee record as a CSV row
	for _, emp := range data {
		record := []string{
			emp.Name,
			strconv.Itoa(emp.Age),
			fmt.Sprintf("%.2f", emp.Salary),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("could not write record to file: %v", err)
		}
	}

	return nil
}

// readFromCSV reads data from a CSV file and returns a slice of Employee structs.
func readFromCSV(filename string) ([]Employee, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open the file: %v", err)
	}
	defer file.Close()

	// Read the CSV data into a slice of records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read data from CSV file: %v", err)
	}

	// Skip the header row and parse employee data
	var employees []Employee
	for _, record := range records[1:] { // Skip header
		age, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("invalid age value: %v", err)
		}

		salary, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid salary value: %v", err)
		}

		employees = append(employees, Employee{
			Name:   record[0],
			Age:    age,
			Salary: salary,
		})
	}

	return employees, nil
}
