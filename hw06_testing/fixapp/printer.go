package fixapp

import (
	"sort"
	"strings"
)

func PrintStaff(staff []Employee) (string, error) {
	// Create a slice of employees for sorting
	employees := make([]Employee, 0, len(staff))

	employees = append(employees, staff...)

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
