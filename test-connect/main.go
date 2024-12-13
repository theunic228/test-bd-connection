package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

var db *sql.DB

// Структура для продукта
type Product struct {
	ID    int
	Name  string
	Price float64
}

type Test struct {
	ID      int
	Name    string
	Surname string
}

// Инициализация соединения с базой данных
func init() {
	var err error
	// Строка подключения к PostgreSQL
	connStr := "user=postgres dbname=postgres password=password host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Функция для извлечения данных из базы данных
func getProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func getTest() ([]Test, error) {
	rows, err := db.Query("SELECT id, name, surname FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []Test
	for rows.Next() {
		var t Test
		if err := rows.Scan(&t.ID, &t.Name, &t.Surname); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tests, nil
}

// Функция для добавления нового продукта в базу данных
func addProduct(name string, price float64) error {
	_, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", name, price)
	return err
}

func addTest(name, surname string) error {
	_, err := db.Exec("INSERT INTO test (name, surname) VALUES ($1, $2)", name, surname)
	return err
}

// Обработчик для отображения продуктов на веб-странице
func productsHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные о продуктах
	products, err := getProducts()
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

func testHandler(w http.ResponseWriter, r *http.Request) {
	test, err := getTest()
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
				<label for="price">Price:</label>
				<input type="surname" id="surname" name="surname" required><br>
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

// Обработчик для добавления нового продукта
func addProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	price := r.FormValue("price")
	if name == "" || price == "" {
		http.Error(w, "Name and price are required", http.StatusBadRequest)
		return
	}

	priceValue, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	err = addProduct(name, priceValue)
	if err != nil {
		http.Error(w, "Error adding product", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}

func addTestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	if name == "" || surname == "" {
		http.Error(w, "Name and surname are required", http.StatusBadRequest)
		return
	}

	err := addTest(name, surname)
	if err != nil {
		http.Error(w, "Error adding test", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/test", http.StatusSeeOther)
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Главная страница</title>
		</head>
		<body>
			<h1>Добро пожаловать!</h1>
			<p>Выберите таблицу, нажав на кнопку ниже:</p>
			<form action="http://localhost:8080/products" method="get" style="display:inline;">
				<button type="submit">products</button>
			</form>
			<form action="http://localhost:8080/test" method="get" style="display:inline;">
				<button type="submit">test</button>
			</form>
		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при обработке шаблона", http.StatusInternalServerError)
	}
}

func main() {
	// Запуск HTTP-сервера
	http.HandleFunc("/", mainPageHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/add", addProductHandler)

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/add_t", addTestHandler)

	// Начало работы сервера на порту 8080
	fmt.Println("Server started at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
