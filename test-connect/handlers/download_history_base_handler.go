package handlers

import (
	"encoding/csv"

	"net/http"
	"test-connect/gets"
	"time"
)

func DownloadCSVHandler(w http.ResponseWriter, r *http.Request) {
	history_base, err := gets.GetHistoryBase()
	if err != nil {
		http.Error(w, "Error getting history base", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=logs.csv")

	_, err = w.Write([]byte("\xEF\xBB\xBF"))
	if err != nil {
		http.Error(w, "Error writing BOM", http.StatusInternalServerError)
		return
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	err = writer.Write([]string{"History_Base_Id", "Message_History", "Create_Date"})
	if err != nil {
		http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
		return
	}

	for _, record := range history_base {
		historyBaseId := record.History_Base_Id

		var createDate string
		if t, err := time.Parse(time.RFC3339, record.Create_Date); err == nil {
			createDate = t.Format(time.RFC3339)
		} else {
			createDate = record.Create_Date
		}

		row := []string{
			historyBaseId,
			record.Message_History,
			createDate,
		}
		err := writer.Write(row)
		if err != nil {
			http.Error(w, "Error writing CSV record", http.StatusInternalServerError)
			return
		}
	}
}
