package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddTasks(Title, Description, Status, Created_By, Due_Date string) error {
	_, err := database.DB.Exec("insert into tasks (Title, Description, Status, Created_By, Due_Date) values ($1, $2, $3, $4, $5)", Title, Description, Status, Created_By, Due_Date)
	return err
}
