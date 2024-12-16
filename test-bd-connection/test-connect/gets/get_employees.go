package gets

import (
	"test-connect/database"
)

type Employees struct {
	Employee_Id   string
	First_Name    string
	Last_Name     string
	Email         string
	Password      string
	Department_Id string
	Hired_Date    string
}

func GetEmployees() ([]Employees, error) {
	rows, err := database.DB.Query("select employee_id, first_name, last_name, email, \"password\", department_id, hired_date from \"PPV3\".employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employees
	for rows.Next() {
		var e Employees
		if err := rows.Scan(&e.Employee_Id, &e.First_Name, &e.Last_Name, &e.Email, &e.Password, &e.Department_Id, &e.Hired_Date); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return employees, nil
}
