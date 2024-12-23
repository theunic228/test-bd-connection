package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddEmployees(First_Name, Last_Name, Patronymic, Email, Password, Department_Id, Position_id string) error {
	_, err := database.DB.Exec("insert into employees (first_name, last_name, patronymic, email, \"password\", department_id, position_id) values ($1, $2, $3, $4, $5, $6, $7)", First_Name, Last_Name, Patronymic, Email, Password, Department_Id, Position_id)
	return err
}
