package gets

import (
	"fmt"
	"test-connect/database"
)

type Employees struct {
	Employee_Id   string
	First_Name    string
	Last_Name     string
	Patronymic    string
	Email         string
	Password      string
	Department_Id string
	Position_id   string
}

func GetEmployees() ([]Employees, error) {
	rows, err := database.DB.Query("select employee_id, first_name, last_name, patronymic, email, \"password\", (select department_name from departments d where department_id = e.department_id), (select position_name from positions p where position_id = e.position_id)  from employees e")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employees
	for rows.Next() {
		var e Employees
		if err := rows.Scan(&e.Employee_Id, &e.First_Name, &e.Last_Name, &e.Patronymic, &e.Email, &e.Password, &e.Department_Id, &e.Position_id); err != nil {
			fmt.Println(err)
			return nil, err
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return employees, nil
}
