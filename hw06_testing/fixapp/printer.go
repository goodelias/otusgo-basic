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
	// Create a slice of employees for sorting
	employees := make([]Employee, 0, len(staff))
	for _, emp := range staff {
		if emp.UserID <= 0 || emp.Age <= 0 || emp.Name == "" || emp.DepartmentID <= 0 {
			return "", fmt.Errorf("%w: %v", ErrZeroOrNegativeField, emp)
		}
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
