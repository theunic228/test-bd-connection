package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

func AddDepartments(Name, Description string) error {
	_, err := database.DB.Exec("INSERT INTO \"PPV2\".departments (name, description) VALUES ($1, $2)", Name, Description)
	return err
}
