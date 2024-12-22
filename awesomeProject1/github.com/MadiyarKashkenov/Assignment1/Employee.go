package main

import (
	"bufio"
	"fmt"
	"os"
)

type Employee interface {
	GetDetails() string
}
type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d Tenge", f.ID, f.Name, f.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (p PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d Tenge, Hours Worked: %.2f", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

type Company struct {
	Employees map[string]Employee
}

var comp Company

func AddEmployee(e Employee) {
	switch emp := e.(type) {
	case FullTimeEmployee:
		comp.Employees[emp.Name] = emp
	case PartTimeEmployee:
		comp.Employees[emp.Name] = emp
	}
}
func ListEmployees() {
	for _, emp := range comp.Employees {
		fmt.Println(emp.GetDetails())
	}
}

func main() {
	comp.Employees = make(map[string]Employee)
	var ch int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1: Add Employee:")
		fmt.Println("2: List Employees:")
		fmt.Println("3: Exit:")
		fmt.Println(" ")
		fmt.Print("Enter a request:")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			fmt.Println("FullTimeEmployee or PartTimeEmployee:")
			var cch int
			var id uint64
			var name string
			var salary uint32
			fmt.Scan(&cch)
			switch cch {
			case 1:
				fmt.Print("Enter ID:")
				fmt.Scan(&id)
				fmt.Print("Enter Name:")
				scanner.Scan()
				name = scanner.Text()
				fmt.Print("Enter Salary:")
				fmt.Scan(&salary)
				AddEmployee(FullTimeEmployee{ID: id, Name: name, Salary: salary})

			case 2:
				var hourlyRate uint64
				var hoursWorked float32
				fmt.Print("Enter ID:")
				fmt.Scan(&id)
				fmt.Print("Enter Name:")
				scanner.Scan()
				name = scanner.Text()
				fmt.Print("Hourly Rate:")
				fmt.Scan(&hourlyRate)
				fmt.Print("Hours Worked:")
				fmt.Scan(&hoursWorked)
				AddEmployee(PartTimeEmployee{ID: id, Name: name, HourlyRate: hourlyRate, HoursWorked: hoursWorked})

			}
		case 2:
			for _, emp := range comp.Employees {
				fmt.Println(emp.GetDetails())
			}
		case 3:
			break
		}
	}
}
