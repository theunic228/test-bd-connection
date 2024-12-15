package gets

import (
	"test-connect/database"
)

type TaskComments struct {
	Comment_Id   int
	Task_Id      int
	Author_Id    int
	Comment_Text string
	Created_At   string
}

func GetTaskComments() ([]TaskComments, error) {
	rows, err := database.DB.Query("select comment_id, task_id, author_id, comment_text, created_at from \"PPV2\".task_comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task_comments []TaskComments
	for rows.Next() {
		var tc TaskComments
		if err := rows.Scan(&tc.Comment_Id, &tc.Task_Id, &tc.Author_Id, &tc.Comment_Text, &tc.Created_At); err != nil {
			return nil, err
		}
		task_comments = append(task_comments, tc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return task_comments, nil
}
