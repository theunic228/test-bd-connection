package gets

import (
	"test-connect/database"
)

type Test struct {
	ID      int
	Name    string
	Surname string
}

func GetTest() ([]Test, error) {
	rows, err := database.DB.Query("SELECT id, name, surname FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var t Test
		if err := rows.Scan(&t.ID, &t.Name, &t.Surname); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tests, nil
}
