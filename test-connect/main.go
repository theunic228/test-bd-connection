package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test-connect/database"
	"test-connect/handlers"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

// Функция для добавления нового продукта в базу данных
func addProduct(name string, price float64) error {
	_, err := database.DB.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", name, price)
	return err
}

func addTest(name, surname string) error {
	_, err := database.DB.Exec("INSERT INTO test (name, surname) VALUES ($1, $2)", name, surname)
	return err
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

func main() {
	// Запуск HTTP-сервера
	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/add", addProductHandler)

	http.HandleFunc("/test", handlers.TestHandler)
	http.HandleFunc("/add_t", addTestHandler)

	// Начало работы сервера на порту 8080
	fmt.Println("Server started at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
