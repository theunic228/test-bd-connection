package handlers

import (
	"html/template"
	"net/http"
	"test-connect/gets"
)

// Обработчик для отображения продуктов на веб-странице
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные о продуктах
	products, err := gets.GetProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	// Используем шаблон для отображения данных
	tmpl, err := template.New("products").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Products</title>
		</head>
		<body>
			<h1>Product List</h1>
			<table border="1">
				<tr>
					<th>ID</th>
					<th>Name</th>
					<th>Price</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.ID}}</td>
					<td>{{.Name}}</td>
					<td>{{.Price}}</td>
				</tr>
				{{end}}
			</table>
			<h2>Add New Product</h2>
			<form action="/add" method="POST">
				<label for="name">Name:</label>
				<input type="text" id="name" name="name" required><br>
				<label for="price">Price:</label>
				<input type="number" id="price" name="price" step="0.01" required><br>
				<input type="submit" value="Add Product">
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

	// Рендерим шаблон с данными
	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
