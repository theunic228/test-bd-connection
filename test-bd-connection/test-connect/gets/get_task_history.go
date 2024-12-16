package gets

import (
	"test-connect/database"
)

type TaskHistory struct {
	History_Id         string
	Task_Id            string
	Action_Time        string
	Action_Description string
	Performed_By       string
}

func GetTaskHistory() ([]TaskHistory, error) {
	rows, err := database.DB.Query("select history_id, task_id, action_time, action_description, performed_by from \"PPV3\".task_history")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task_history []TaskHistory
	for rows.Next() {
		var th TaskHistory
		if err := rows.Scan(&th.History_Id, &th.Task_Id, &th.Action_Time, &th.Action_Description, &th.Performed_By); err != nil {
			return nil, err
		}
		task_history = append(task_history, th)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return task_history, nil
}
