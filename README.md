# 📌 TODO REST API (Golang + Fiber + PostgreSQL)

### 📖 Описание проекта

Этот проект представляет собой REST API для управления задачами (TODO-лист) на Golang с использованием Fiber, PostgreSQL и pgx.

### 📌 Функционал API:

✅ Создание задач (POST /tasks)

✅ Получение списка задач (GET /tasks)

✅ Обновление задач (PUT /tasks/:id)

✅ Удаление задач (DELETE /tasks/:id)

✅ Поддержка Graceful Shutdown

✅ Автоматическое логирование запросов

✅ Миграции базы данных с помощью goose

🚀 Установка и запуск

### 1️⃣ Клонирование репозитория

```bash
git clone https://gitlab.com/readyblast/skillsrock-todo.git
cd skillsrock-todo
```

### 2️⃣ Настройка переменных окружения

Создайте файл .env в папке config/ и добавьте параметры:

```env
DB_USER=admin
DB_PASSWORD=password
DB_HOST=host
DB_PORT=5432
DB_NAME=databasename
DB_SSL_MODE=disable
```

### 3️⃣ Установка зависимостей

```bash
go mod tidy
```

### 4️⃣ Запуск миграций

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir migrations postgres "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSL_MODE" up
```

### 5️⃣ Запуск сервера

```bash
go run cmd/main.go
```

Сервер запустится на http://localhost:3000 🚀

### 🐳 Запуск с Docker (опционально)

🔹 1. Соберите и запустите контейнеры

```docker
docker-compose up --build
```

🔹 2. Остановка контейнеров

```docker
docker-compose down
```

### 📜 API Эндпоинты

➤ Создание задачи

POST /tasks

🔹 Запрос:

```json
{
  "title": "New Task",
  "description": "Description of the task"
}
```

🔹 Ответ:

```json
{
  "id": 1,
  "title": "New Task",
  "description": "Description of the task",
  "status": "new",
  "created_at": "2025-03-20T17:10:00Z",
  "updated_at": "2025-03-20T17:10:00Z"
}
```

➤ Получение списка задач

GET /tasks

🔹 Ответ:

```json
[
  {
    "id": 1,
    "title": "New Task",
    "description": "Description of the task",
    "status": "new",
    "created_at": "2025-03-20T17:10:00Z",
    "updated_at": "2025-03-20T17:10:00Z"
  }
]
```

➤ Обновление задачи

PUT /tasks/1

🔹 Запрос:

```json
{
  "title": "Updated Task",
  "status": "done"
}
```

🔹 Ответ:

```json
{
  "id": 1,
  "title": "Updated Task",
  "description": "Description of the task",
  "status": "done",
  "created_at": "2025-03-20T17:10:00Z",
  "updated_at": "2025-03-20T18:00:00Z"
}
```

➤ Удаление задачи

DELETE /tasks/1

🔹 Ответ:

204 No Content

### 🛠️ Используемые технологии

🚀 Golang

⚡ Fiber (быстрая альтернатива Express)

🗄️ PostgreSQL (через pgx)

🔀 Goose (миграции)

🐳 Docker (контейнеризация, опционально)
