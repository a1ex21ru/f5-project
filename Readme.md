# F5 Notes API

REST API для управления заметками с веб-интерфейсом, построенный на Go, Gin, GORM и PostgreSQL.

## 🚀 Возможности

- ✅ CRUD операции для заметок (Create, Read, Update, Delete)
- ✅ RESTful API с JSON
- ✅ Swagger документация
- ✅ Веб-интерфейс для управления заметками
- ✅ PostgreSQL база данных
- ✅ Docker контейнеризация
- ✅ Auto-миграции с GORM
- ✅ Health check эндпоинт

## 📋 Требования

- Docker & Docker Compose
- Go 1.24+ (для локальной разработки)
- Make (опционально)

## 🏗️ Структура проекта

```
f5-project/
├── cmd/
│   ├── app/
│   │   └── main.go           # Точка входа приложения
│   └── server/
│       └── server.go         # HTTP сервер
├── internal/
│   ├── adapters/
│   │   └── postgres/
│   │       └── client.go     # PostgreSQL клиент и миграции
│   ├── handlers/
│   │   └── handler.go        # HTTP обработчики
│   ├── models/
│   │   └── note.go           # Модели данных
│   ├── repository/
│   │   └── repo.go           # Слой доступа к данным
│   ├── routes/
│   │   └── routes.go         # Маршруты API
│   └── service/
│       └── service.go        # Бизнес-логика
├── static/
│   ├── index.html            # Главная страница
│   ├── create.html           # Страница создания заметки
│   └── edit.html             # Страница редактирования
├── docs/                     # Swagger документация
├── db/migrations/            # SQL миграции
├── docker-compose.yml
├── Dockerfile
└── go.mod
```

## 🚀 Быстрый старт

### Запуск с Docker Compose

1. **Клонируйте репозиторий:**
```bash
git clone <repository-url>
cd f5-project
```

2. **Запустите приложение:**
```bash
docker-compose up -d
```

3. **Проверьте статус:**
```bash
docker-compose ps
```

4. **Откройте в браузере:**
- Веб-интерфейс: http://localhost:8080
- Swagger документация: http://localhost:8080/swagger/index.html
- Health check: http://localhost:8080/healthcheck

### Локальная разработка

1. **Запустите PostgreSQL:**
```bash
docker-compose up -d postgres
```

2. **Установите зависимости:**
```bash
go mod download
```

3. **Сгенерируйте Swagger документацию:**
```bash
swag init -g cmd/app/main.go
```

4. **Запустите приложение:**
```bash
go run cmd/app/main.go
```

## 📚 API Документация

### Endpoints

#### Notes

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/notes` | Получить все заметки |
| GET | `/api/notes/:id` | Получить заметку по ID |
| POST | `/api/notes` | Создать новую заметку |
| PATCH | `/api/notes/:id` | Обновить заметку |
| DELETE | `/api/notes/:id` | Удалить заметку |

#### System

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/healthcheck` | Проверка здоровья API |
| GET | `/swagger/*` | Swagger UI документация |

### Примеры запросов

**Создать заметку:**
```bash
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My First Note",
    "description": "This is a test note"
  }'
```

**Получить все заметки:**
```bash
curl http://localhost:8080/api/notes
```

**Получить заметку по ID:**
```bash
curl http://localhost:8080/api/notes/1
```

**Обновить заметку:**
```bash
curl -X PATCH http://localhost:8080/api/notes/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Note",
    "description": "Updated description"
  }'
```

**Удалить заметку:**
```bash
curl -X DELETE http://localhost:8080/api/notes/1
```

## 🗄️ Модель данных

### Note

```go
type Note struct {
    ID          uint      `json:"ID" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"column:name"`
    Description string    `json:"description,omitempty" gorm:"column:description"`
    CreatedAt   time.Time `json:"CreatedAt"`
    UpdatedAt   time.Time `json:"UpdatedAt"`
    DeletedAt   *time.Time `json:"DeletedAt,omitempty" gorm:"index"`
}
```

## 🔧 Конфигурация

### Переменные окружения

Настройки в `docker-compose.yml`:

```yaml
environment:
  POSTGRES_USER: user
  POSTGRES_PASSWORD: admin
  POSTGRES_DB: practice
  POSTGRES_HOST: f5-postgres
  GIN_MODE: release
  port: "8080"
```

### PostgreSQL

База данных автоматически создается и мигрируется при запуске контейнера.

**Строка подключения:**
```
postgres://user:admin@f5-postgres:5432/practice?sslmode=disable
```

## 🛠️ Разработка

### Makefile команды

```bash
# Запуск миграций вручную (goose)
make migrateUp

# Очистка всех Docker ресурсов
make removeAll
```

### Обновление Swagger документации

После изменения API аннотаций:

```bash
swag init -g cmd/app/main.go
```

### Пересборка контейнера

```bash
docker-compose down
docker-compose build --no-cache
docker-compose up -d
```

## 📝 Логирование

Просмотр логов:

```bash
# Все сервисы
docker-compose logs -f

# Только API
docker-compose logs -f f5-api

# Только PostgreSQL
docker-compose logs -f postgres
```

## 🐛 Отладка

### Вход в контейнер

```bash
# API контейнер
docker exec -it f5-api sh

# PostgreSQL контейнер
docker exec -it f5-postgres psql -U user -d practice
```

### Проверка базы данных

```docker
-- Войти в PostgreSQL
docker exec -it f5-postgres psql -U user -d practice

-- Показать все таблицы
\dt

-- Показать все заметки
SELECT * FROM notes;

-- Выйти
\q
```

## 🏛️ Архитектура

Проект использует Clean Architecture с разделением на слои:

1. **Handlers** - HTTP обработчики запросов
2. **Service** - Бизнес-логика и валидация
3. **Repository** - Работа с базой данных
4. **Models** - Структуры данных
5. **Routes** - Маршрутизация

```
Request → Handler → Service → Repository → Database
                                    ↓
Response ← Handler ← Service ← Repository
```

## 🔐 Безопасность

- ✅ Непривилегированный пользователь в Docker
- ✅ Минимальный Alpine образ
- ✅ Статически слинкованный бинарник
- ✅ Soft delete для заметок (DeletedAt)
- ✅ SQL инъекции защищены GORM

## 📦 Зависимости

### Основные

- **Gin** - HTTP веб-фреймворк
- **GORM** - ORM для Go
- **PostgreSQL Driver** - Драйвер БД
- **Swag** - Swagger документация

### Полный список

Смотрите `go.mod` для всех зависимостей.

## 🚨 Troubleshooting

### Swagger не загружается

1. Проверьте, что docs/ скопированы в контейнер:
```bash
docker exec f5-api ls -la /app/docs
```

2. Пересгенерируйте документацию:
```bash
swag init -g cmd/app/main.go
docker-compose build --no-cache
docker-compose up -d
```

### База данных не подключается

1. Проверьте health check PostgreSQL:
```bash
docker-compose ps
```

2. Проверьте логи:
```bash
docker-compose logs postgres
```

3. Убедитесь, что контейнер запущен:
```bash
docker ps | grep postgres
```

### Порт 8080 уже занят

Измените порт в `docker-compose.yml`:
```yaml
ports:
  - "3000:8080"  # 3000 на хосте, 8080 в контейнере
```

---

**Полезные ссылки:**
- [Gin Documentation](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [Swagger Documentation](https://swagger.io/docs/)
- [Docker Documentation](https://docs.docker.com/)