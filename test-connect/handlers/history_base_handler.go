package handlers

import (
	"encoding/csv"
	"html/template"
	"net/http"
	"strconv"
	"test-connect/gets"
	"time"
)

func HistoryBaseHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetHistoryBase
	history_base, err := gets.GetHistoryBase()
	if err != nil {
		http.Error(w, "Error getting departments", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("history_base").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>history_base</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f9f9f9;
                    color: #333;
                    margin: 0;
                    padding: 20px;
                    display: flex;
                    flex-direction: column;
                    align-items: center;
                }
                h1 {
                    color: #4CAF50;
                    margin-bottom: 20px;
                }
                table {
                    border-collapse: collapse;
                    width: 80%;
                    margin-bottom: 20px;
                    background-color: white;
                    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
                }
                th, td {
                    border: 1px solid #ddd;
                    padding: 8px;
                    text-align: left;
                }
                th {
                    background-color: #4CAF50;
                    color: white;
                }
                tr:nth-child(even) {
                    background-color: #f2f2f2;
                }
                tr:hover {
                    background-color: #ddd;
                }
                button {
                    background-color: #4CAF50;
                    color: white;
                    border: none;
                    padding: 10px 20px;
                    font-size: 16px;
                    border-radius: 5px;
                    cursor: pointer;
                    transition: background-color 0.3s ease, transform 0.2s ease;
                }
                button:hover {
                    background-color: #45a049;
                    transform: translateY(-2px);
                }
                button:active {
                    transform: translateY(1px);
                }
                form {
                    margin-top: 20px;
                }
            </style>
		</head>
		<body>
			<h1>History Base List</h1>
			<table border="1">
				<tr>
					<th>History_Base_Id</th>
					<th>Message_History</th>
					<th>Create_Date</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.History_Base_Id}}</td>
					<td>{{.Message_History}}</td>
					<td>{{.Create_Date}}</td>
				</tr>
				{{end}}
			</table>
			<br>
            <form action="/download-csv" method="get" style="display:inline;">
				<button type="submit">Download CSV</button>
			</form>
			<form action="http://localhost:8080/" method="get" style="display:inline;">
				<button type="submit">main</button>
			</form>
		</body>
		</html>
	`)

	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, history_base)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

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
		historyBaseId := strconv.Itoa(record.History_Base_Id)

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
