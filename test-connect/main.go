package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Подключение к базе данных
	connStr := "postgres://postgres:123@localhost:5432/postgres"
	dbpool := connectToDB(connStr)
	defer dbpool.Close()

	// Получение и вывод таблиц и данных
	processTables(dbpool)

	// Ввод данных в таблицу
	fmt.Println("\nВведите название таблицы для добавления данных:")
	reader := bufio.NewReader(os.Stdin)
	tableName, _ := reader.ReadString('\n')
	tableName = strings.TrimSpace(tableName)
	insertIntoTable(dbpool, tableName)
}

// Подключение к базе данных PostgreSQL
func connectToDB(connStr string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v\n", err)
	}
	return dbpool
}

// Обработка таблиц: вывод имен таблиц и их данных
func processTables(dbpool *pgxpool.Pool) {
	// Запрос для получения списка таблиц
	queryTables := `
		SELECT table_name
		FROM information_schema.tables
		WHERE table_schema = 'public'
	`

	// Выполняем запрос
	rows, err := dbpool.Query(context.Background(), queryTables)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса для получения таблиц: %v\n", err)
	}
	defer rows.Close()

	// Перебираем таблицы
	fmt.Println("Значения таблиц в схеме 'public':")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Ошибка чтения названия таблицы: %v\n", err)
		}

		// Обрабатываем текущую таблицу
		processTableData(dbpool, tableName)
	}

	// Проверяем ошибки итерации
	if rows.Err() != nil {
		log.Fatalf("Ошибка при итерации по таблицам: %v\n", rows.Err())
	}
}

// Обработка данных для конкретной таблицы
func processTableData(dbpool *pgxpool.Pool, tableName string) {
	fmt.Printf("\nТаблица: %s\n", tableName)
	fmt.Println("Данные (первые 10 строк):")

	// Формируем запрос для получения данных таблицы
	queryData := fmt.Sprintf("SELECT * FROM %s LIMIT 10", tableName)

	// Выполняем запрос
	dataRows, err := dbpool.Query(context.Background(), queryData)
	if err != nil {
		log.Printf("Ошибка получения данных из таблицы %s: %v\n", tableName, err)
		return
	}
	defer dataRows.Close()

	// Обрабатываем строки данных
	printTableData(dataRows)
}

// Вывод данных строки таблицы
func printTableData(dataRows pgx.Rows) {
	// Выводим названия колонок
	columns := dataRows.FieldDescriptions()
	for _, col := range columns {
		fmt.Printf("%s\t", col.Name)
	}
	fmt.Println()

	// Выводим строки данных
	for dataRows.Next() {
		values, err := dataRows.Values()
		if err != nil {
			log.Fatalf("Ошибка чтения значений строки: %v\n", err)
		}
		for _, value := range values {
			fmt.Printf("%v\t", value)
		}
		fmt.Println()
	}

	// Проверяем ошибки итерации
	if dataRows.Err() != nil {
		log.Fatalf("Ошибка при итерации по данным: %v\n", dataRows.Err())
	}
}

// Ввод новых данных в таблицу
func insertIntoTable(dbpool *pgxpool.Pool, tableName string) {
	reader := bufio.NewReader(os.Stdin)

	// Получение информации о колонках таблицы
	queryColumns := fmt.Sprintf(`
		SELECT column_name 
		FROM information_schema.columns 
		WHERE table_name = '%s'
	`, tableName)

	rows, err := dbpool.Query(context.Background(), queryColumns)
	if err != nil {
		log.Fatalf("Ошибка получения колонок таблицы %s: %v\n", tableName, err)
	}
	defer rows.Close()

	columns := []string{}
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			log.Fatalf("Ошибка чтения имени колонки: %v\n", err)
		}
		columns = append(columns, columnName)
	}

	// Формирование INSERT-запроса
	values := []string{}
	for _, column := range columns {
		fmt.Printf("Введите значение для колонки %s: ", column)
		value, _ := reader.ReadString('\n')
		values = append(values, strings.TrimSpace(value))
	}

	placeholders := []string{}
	for i := range values {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	// Выполнение INSERT-запроса
	_, err = dbpool.Exec(context.Background(), insertQuery, interfaceSlice(values)...)
	if err != nil {
		log.Fatalf("Ошибка вставки данных в таблицу %s: %v\n", tableName, err)
	}

	fmt.Println("Данные успешно добавлены.")
}

// Преобразование []string в []interface{} для Exec
func interfaceSlice(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
