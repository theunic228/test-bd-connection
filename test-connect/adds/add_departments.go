package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddDepartments(Department_Name, Managment_Id string) error {
	_, err := database.DB.Exec("insert into departments (department_name, managment_id) values ($1, (select managment_id from managments m where m.managment_name = $2))", Department_Name, Managment_Id)
	return err
}
