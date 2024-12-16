package gets

import (
	"test-connect/database"
)

type Statuses struct {
	Status_Id   string
	Status_Name string
	Description string
}

func GetStatuses() ([]Statuses, error) {
	rows, err := database.DB.Query("select status_id, status_name, description from \"PPV3\".statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []Statuses
	for rows.Next() {
		var s Statuses
		if err := rows.Scan(&s.Status_Id, &s.Status_Name, &s.Description); err != nil {
			return nil, err
		}
		statuses = append(statuses, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return statuses, nil
}
