package fixapp

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	ErrEmptyStaff          = errors.New("no employees found")
	ErrZeroOrNegativeField = errors.New("employee has zero or negative value field(s)")
)

func PrintStaff(staff []Employee) (string, error) {
	// Check if the slice is empty
	if len(staff) == 0 {
		return "", fmt.Errorf("%w: %v", ErrEmptyStaff, staff)
	}

	// Remove duplicate employees using a map for uniqueness
	uniqueEmployees := make(map[string]Employee)
	for _, emp := range staff {
		if emp.UserID <= 0 || emp.Age <= 0 || emp.Name == "" || emp.DepartmentID <= 0 {
			return "", fmt.Errorf("%w: %v", ErrZeroOrNegativeField, emp)
		}
		key := fmt.Sprintf("%d-%d-%s-%d", emp.UserID, emp.Age, emp.Name, emp.DepartmentID)
		uniqueEmployees[key] = emp
	}

	// Create a slice of employees for sorting
	employees := make([]Employee, 0, len(uniqueEmployees))
	for _, emp := range uniqueEmployees {
		employees = append(employees, emp)
	}

	// Sort employees by ID
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].UserID < employees[j].UserID
	})

	// Write employees data in the builder
	builder := strings.Builder{}
	for _, emp := range employees {
		builder.WriteString(emp.String() + "\n")
	}

	return builder.String(), nil
}
