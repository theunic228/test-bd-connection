package gets

import (
	"fmt"
	"test-connect/database"
)

type Managments struct {
	Managment_Id   string
	Managment_Name string
}

func GetManagments() ([]Managments, error) {
	rows, err := database.DB.Query("select managment_id, managment_name from managments")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var managments []Managments
	for rows.Next() {
		var m Managments
		if err := rows.Scan(&m.Managment_Id, &m.Managment_Name); err != nil {
			fmt.Println(err)
			return nil, err
		}
		managments = append(managments, m)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return managments, nil
}
