package gets

import (
	"fmt"
	"test-connect/database"
)

type TaskComments struct {
	Comment_Id   string
	Task_Id      string
	Author_Id    string
	Comment_Text string
	Created_At   string
}

func GetTaskComments() ([]TaskComments, error) {
	rows, err := database.DB.Query("select * from task_comments")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var task_comments []TaskComments
	for rows.Next() {
		var tc TaskComments
		if err := rows.Scan(&tc.Comment_Id, &tc.Task_Id, &tc.Author_Id, &tc.Comment_Text, &tc.Created_At); err != nil {
			fmt.Println(err)
			return nil, err
		}
		task_comments = append(task_comments, tc)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return task_comments, nil
}
