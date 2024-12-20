package gets

import (
	"fmt"
	"test-connect/database"
)

type Statuses struct {
	Status_Id   string
	Status_Name string
}

func GetStatuses() ([]Statuses, error) {
	rows, err := database.DB.Query("select * from statuses")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var statuses []Statuses
	for rows.Next() {
		var s Statuses
		if err := rows.Scan(&s.Status_Id, &s.Status_Name); err != nil {
			fmt.Println(err)
			return nil, err
		}
		statuses = append(statuses, s)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return statuses, nil
}
