package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Исправляем вызов функции на существующий GetTest
	test, err := gets.GetTest()
	if err != nil {
		http.Error(w, "Error getting test", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("test").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Test</title>
		</head>
		<body>
			<h1>Test List</h1>
			<table border="1">
				<tr>
					<th>ID</th>
					<th>Name</th>
					<th>Surname</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.ID}}</td>
					<td>{{.Name}}</td>
					<td>{{.Surname}}</td>
				</tr>
				{{end}}
			</table>
			<h2>Add New Test</h2>
			<form action="/add_t" method="POST">
				<label for="name">Name:</label>
				<input type="text" id="name" name="name" required><br>
				<label for="surname">Price:</label>
				<input type="text" id="surname" name="surname" required><br>
				<input type="submit" value="Add Test">
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

	err = tmpl.Execute(w, test)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
