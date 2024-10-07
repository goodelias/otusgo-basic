package fixapp

import "fmt"

type Employee struct {
	UserID       int    `json:"userId"`
	Age          int    `json:"age"`
	Name         string `json:"name"`
	DepartmentID int    `json:"departmentId"`
}

func NewEmployee(userID int, age int, name string, departmentID int) Employee {
	return Employee{
		UserID:       userID,
		Age:          age,
		Name:         name,
		DepartmentID: departmentID,
	}
}

func (e Employee) String() string {
	return fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", e.UserID, e.Age, e.Name, e.DepartmentID)
}
