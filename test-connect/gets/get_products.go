package gets

import (
	"test-connect/database"
)

// Структура для продукта
type Product struct {
	ID    int
	Name  string
	Price float64
}

// Функция для извлечения данных из базы данных
func GetProducts() ([]Product, error) {
	rows, err := database.DB.Query("SELECT id, name, price FROM products")
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
