package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddStatuses(Status_Name string) error {
	_, err := database.DB.Exec("insert into statuses (status_name) values ($1)", Status_Name)
	return err
}
