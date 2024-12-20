package gets

import (
	"fmt"
	"test-connect/database"
)

type Positions struct {
	Position_Id   string
	Position_Name string
}

func GetPositions() ([]Positions, error) {
	rows, err := database.DB.Query("select position_id, position_name from positions")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var positions []Positions
	for rows.Next() {
		var p Positions
		if err := rows.Scan(&p.Position_Id, &p.Position_Name); err != nil {
			return nil, err
		}
		positions = append(positions, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return positions, nil
}
