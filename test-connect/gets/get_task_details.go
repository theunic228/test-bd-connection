package gets

import (
	"fmt"
	"test-connect/database"
)

type TaskDetails struct {
	Task_Id            string
	Title              string
	Description        string
	Status             string
	Created_By         string
	Task_Created_At    string
	Due_Date           string
	Author_Id          string
	Comment_Text       string
	Comment_Created_At string
	File_Name          string
	File_Path          string
	Uploaded_At        string
}

func GetTaskDetails() ([]TaskDetails, error) {
	rows, err := database.DB.Query("select task_id, COALESCE(task_title, 'не заполнено'), COALESCE(task_description, 'не заполнено'), COALESCE((select status_name from statuses s where s.status_id = task_status), 'не заполнено'), COALESCE((select e.last_name from employees e where e.employee_id = task_creator), 'не заполнено'), task_created_at, NULLIF(task_due_date, '2000-01-01')::DATE, COALESCE((select e.last_name from employees e where e.employee_id = comment_author), 'не заполнено'), COALESCE(comment_text, 'не заполнено'), comment_created_at, COALESCE(file_name, 'не заполнено'), COALESCE(file_path, 'не заполнено'), file_uploaded_at from task_details ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var taskDetails []TaskDetails
	for rows.Next() {
		var td TaskDetails
		if err := rows.Scan(&td.Task_Id, &td.Title, &td.Description, &td.Status, &td.Created_By, &td.Task_Created_At, &td.Due_Date, &td.Author_Id, &td.Comment_Text, &td.Comment_Created_At, &td.File_Name, &td.File_Path, &td.Uploaded_At); err != nil {
			return nil, err
		}
		taskDetails = append(taskDetails, td)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return taskDetails, nil
}
