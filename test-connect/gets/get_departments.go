package gets

import (
	"fmt"
	"test-connect/database"
)

type Departments struct {
	Department_Id   string
	Department_Name string
	Managment_Id    string
}

func GetDepartments() ([]Departments, error) {
	rows, err := database.DB.Query("select department_id, department_name, (select managment_name from managments where managment_id = d.managment_id) from departments d")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var departments []Departments
	for rows.Next() {
		var d Departments
		if err := rows.Scan(&d.Department_Id, &d.Department_Name, &d.Managment_Id); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, nil
}
