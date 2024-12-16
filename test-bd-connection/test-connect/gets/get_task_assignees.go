package gets

import (
	"test-connect/database"
)

type TaskAssignees struct {
	Task_Id     string
	Employee_Id string
}

func GetTaskAssignees() ([]TaskAssignees, error) {
	rows, err := database.DB.Query("select task_id, employee_id from \"PPV3\".task_assignees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task_assignees []TaskAssignees
	for rows.Next() {
		var ta TaskAssignees
		if err := rows.Scan(&ta.Task_Id, &ta.Employee_Id); err != nil {
			return nil, err
		}
		task_assignees = append(task_assignees, ta)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return task_assignees, nil
}
