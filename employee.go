// employee.go
package main

import (
	"strings"
	"sort"
)

// Employee struct stores information about an employee.
type Employee struct {
	Name   string
	Age    int
	Salary float64
}

// SortEmployeesByName sorts the slice of employees by their names.
func SortEmployeesByName(employees []Employee) {
	sort.Slice(employees, func(i, j int) bool {
		return strings.ToLower(employees[i].Name) < strings.ToLower(employees[j].Name)
	})
}
