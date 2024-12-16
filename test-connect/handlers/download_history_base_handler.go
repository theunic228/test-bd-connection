package handlers

import (
	"encoding/csv"

	"net/http"
	"test-connect/gets"
	"time"
)

func DownloadCSVHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из базы данных
	history_base, err := gets.GetHistoryBase()
	if err != nil {
		http.Error(w, "Error getting history base", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для скачивания файла
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=logs.csv")

	// Создаем новый CSV writer
	// Добавляем BOM для корректной кодировки UTF-8 в Excel
	_, err = w.Write([]byte("\xEF\xBB\xBF"))
	if err != nil {
		http.Error(w, "Error writing BOM", http.StatusInternalServerError)
		return
	}

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Записываем заголовки в CSV
	err = writer.Write([]string{"History_Base_Id", "Message_History", "Create_Date"})
	if err != nil {
		http.Error(w, "Error writing CSV header", http.StatusInternalServerError)
		return
	}

	// Записываем данные из базы в CSV
	for _, record := range history_base {
		// Преобразуем History_Base_Id в строку
		historyBaseId := record.History_Base_Id

		// Проверяем тип Create_Date (если это строка, оставляем как есть)
		var createDate string
		if t, err := time.Parse(time.RFC3339, record.Create_Date); err == nil {
			// Если это time.Time, то форматируем
			createDate = t.Format(time.RFC3339)
		} else {
			// Если это не время, то просто оставляем строку
			createDate = record.Create_Date
		}

		// Создаем строку для записи в CSV
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
