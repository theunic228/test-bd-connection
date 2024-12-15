package gets

import (
	"test-connect/database"
)

type HistoryBase struct {
	History_Base_Id int
	Message_History string
	Create_Date     string
}

func GetHistoryBase() ([]HistoryBase, error) {
	rows, err := database.DB.Query("select history_base_id, message_history, create_date from \"PPV2\".history_base")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history_base []HistoryBase
	for rows.Next() {
		var hb HistoryBase
		if err := rows.Scan(&hb.History_Base_Id, &hb.Message_History, &hb.Create_Date); err != nil {
			return nil, err
		}
		history_base = append(history_base, hb)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return history_base, nil
}
