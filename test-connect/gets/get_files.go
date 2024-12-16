package gets

import (
	"test-connect/database"
)

type Files struct {
	File_Id     string
	Task_Id     string
	File_Name   string
	File_Path   string
	Uploaded_At string
}

func GetFiles() ([]Files, error) {
	rows, err := database.DB.Query("select file_id, task_id, file_name, file_path, uploaded_at from \"PPV3\".files ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []Files
	for rows.Next() {
		var f Files
		if err := rows.Scan(&f.File_Id, &f.Task_Id, &f.File_Name, &f.File_Path, &f.Uploaded_At); err != nil {
			return nil, err
		}
		files = append(files, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return files, nil
}
