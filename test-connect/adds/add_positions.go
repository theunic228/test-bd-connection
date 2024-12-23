package adds

import (
	"test-connect/database"

	_ "github.com/lib/pq"
)

func AddPositions(Position_Name string) error {
	_, err := database.DB.Exec("insert into positions (position_name) values ($1)", Position_Name)
	return err
}
