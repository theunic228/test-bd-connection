package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddManagments(Managment_Name string) error {
	_, err := database.DB.Exec("insert into managments (managment_name) values ($1)", Managment_Name)
	return err
}
