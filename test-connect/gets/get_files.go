package gets

import (
	"fmt"
	"test-connect/database"
)

type Files struct {
	File_Id     string
	Comment_Id  string
	File_Name   string
	File_Path   string
	Uploaded_At string
}

func GetFiles() ([]Files, error) {
	rows, err := database.DB.Query("select file_id, (select (select e.last_name from employees e where e.employee_id = tc.author_id) || ' добавил(-а) комментарий: ' || comment_text  as comments from task_comments tc where comment_id = f.comment_id ), f.file_name, f.file_path, f.uploaded_at from files f")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []Files
	for rows.Next() {
		var f Files
		if err := rows.Scan(&f.File_Id, &f.Comment_Id, &f.File_Name, &f.File_Path, &f.Uploaded_At); err != nil {
			fmt.Println(err)
			return nil, err
		}
		files = append(files, f)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return files, nil
}
