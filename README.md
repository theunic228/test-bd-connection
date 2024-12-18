# *структура проекта* 

# test-connect
## adds
### add_table.go
## database
### db.go
## gets
### get_table.go
## handlers
### add_table_handler.go
### table_handler.go
# go.mod
# go.sum
# main.go

# *описание основного функционала*

В точке входа main  объявляются `http.HandleFunc()` — это функция, которая используется для регистрации обработчика HTTP-запросов для определенного маршрута. 
`http.ListenAndServe` - запускает сервер. (если произойдет ошибка запуска, то будет обработана ошибка).  

```go
func main() {
	http.HandleFunc("/", handlers.MainPageHandler)
	http.HandleFunc("/departments", handlers.DepartmentsHandler)
	http.HandleFunc("/add_departments",handlers.AddDepartmentsHandler)
	http.HandleFunc("/employees", handlers.EmployeesHandler)
	http.HandleFunc("/add_employees",handlers.AddEmployeesHandler)
	http.HandleFunc("/files", handlers.FilesHandler)
	http.HandleFunc("/history_base", handlers.HistoryBaseHandler)
	http.HandleFunc("/download-csv", handlers.DownloadCSVHandler)
	http.HandleFunc("/statuses", handlers.StatusesHandler)
	http.HandleFunc("/task_assignees", handlers.TaskAssigneesHandler)
	http.HandleFunc("/task_comments", handlers.TaskCommentsHandler)
	http.HandleFunc("/task_history", handlers.TaskHistoryHandler)
	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/add_tasks", handlers.AddTasksHandler)

	fmt.Println("Server started at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // логирование ошибки
	}
}
```

В файле db.go, выполняется подключение к базе данных Postgres.
```go
var DB *sql.DB
// Инициализация соединения с базой данных
func init() {
	var err error
	// Строка подключения к PostgreSQL
	connStr := "user=*** dbname=*** password=*** host=*** port=5442 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Проверка соединения

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
```

Основные компоненты - "Хендлеры" (handlers). В данном проекте это типовые элементы, отвечающие за получение полей таблицы и передачи их в html-шаблон. 
```go
func DepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	departments, err := gets.GetDepartments() // получение таблицы отделов
	if err != nil {
		http.Error(w, "Error getting departments",http.StatusInternalServerError)
		return
	}
	
	// создание шаблона
	tmpl, err := template.New("departments").Parse(` 
	//тут разметка страницы
	`)
	
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	
	err = tmpl.Execute(w, departments) // передача полей таблицы в шаблон 
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
```

В файлах типа "get_",  происходит передача полей таблицы в структуру. Используя database.DB.Query(), выбираются поля и передаются в структуру (через цикл). В конце возвращается массив структур, передающийся в "Хендлер".  
```go
type Departments struct {
	Department_Id string
	Name string
	Description string
}

func GetDepartments() ([]Departments, error) {
	rows, err := database.DB.Query("SELECT department_id, name, description FROM \"PPV3\".departments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []Departments

	for rows.Next() {
		var d Departments
		if err := rows.Scan(&d.Department_Id, &d.Name, &d.Description); err != nil {
			return nil, err
		}
		departments = append(departments, d)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return departments, nil
}
```

У некоторых таблиц, есть возможность добавления. Для этого были созданы дополнительные "Хендлеры". В них обрабатываются поля, полученные из формы. `Name := r.FormValue("Name")`. Затем эти переменные, отвечающие за определенные поля таблиц, передаются в `AddDepartments`. 
```go
func AddDepartmentsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	Name := r.FormValue("Name")
	Description := r.FormValue("Description")

	if Name == "" || Description == "" {
		http.Error(w, "Name and Description are required", http.StatusBadRequest)
		return
	}

	err := adds.AddDepartments(Name, Description)

	if err != nil {
		http.Error(w, "Error adding Department", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/departments", http.StatusSeeOther)
}
```

В `AddDepartments` выполняется передача переменных с помощью запроса в базу данных.
```go
func AddDepartments(Name, Description string) error {
	_, err := database.DB.Exec("INSERT INTO \"PPV3\".departments (name, description) VALUES ($1, $2)", Name, Description)
	return err
}
```

# *установка*

![alt text](https://i.pinimg.com/236x/67/a9/21/67a921e5c26caf1845a97f6033536eff.jpg)

-- Hello, world! -- 

## **Step 1**
clone this repo

## **Step 2**
build main file -> go build 

## **Step 3**
start builded file... 
