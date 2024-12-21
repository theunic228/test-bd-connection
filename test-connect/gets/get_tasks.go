package gets

import (
	"test-connect/database"
)

type Tasks struct {
	Task_Id     string
	Title       string
	Description string
	Status      string
	Created_By  string
	Created_At  string
	Due_Date    string
}

func GetTasks() ([]Tasks, error) {
	rows, err := database.DB.Query("select task_id, title, description, (select s.status_name from statuses s where s.status_id = t.status), (select e.last_name from employees e where e.employee_id = t.created_by), t.created_at, t.due_date from tasks t")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Tasks
	for rows.Next() {
		var t Tasks
		if err := rows.Scan(&t.Task_Id, &t.Title, &t.Description, &t.Status, &t.Created_By, &t.Created_At, &t.Due_Date); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
