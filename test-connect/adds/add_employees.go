package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func AddEmployees(First_Name, Last_Name, Email, Password, Department_Id, Hired_Date string) error {
	_, err := database.DB.Exec("insert into \"PPV2\".employees (first_name, last_name, email, password, department_id, hired_date) values ($1, $2, $3, $4, $5, $6)", First_Name, Last_Name, Email, Password, Department_Id, Hired_Date)
	return err
}
