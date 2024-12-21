package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddDepartments(Name, Description string) error {
	_, err := database.DB.Exec("INSERT INTO \"PPV3\".departments (name, description) VALUES ($1, $2)", Name, Description)
	return err
}
