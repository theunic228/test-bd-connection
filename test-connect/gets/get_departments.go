package gets

import (
	"test-connect/database"
)

type Departments struct {
	Department_Id int
	Name          string
	Description   string
}

func GetDepartments() ([]Departments, error) {
	rows, err := database.DB.Query("SELECT department_id, name, description FROM \"PPV2\".departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []Departments
	for rows.Next() {
		var d Departments
		if err := rows.Scan(&d.Department_Id, &d.Name, &d.Description); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, nil
}
