# Basic API Server

Этот проект представляет собой базовый API-сервер, написанный на Go.

## Содержание
- [Требования](#требования)
- [Настройка](#настройка)
- [Запуск](#запуск)
- [Команды](#команды)
- [Миграции](#миграции)
- [Docker](#docker)
- [Postman](#postman)
- [Swagger](#swagger)

## Требования

- Go 1.22
- Docker и Docker Compose
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate)

## Настройка

1. Спулить репозиторий:

```bash
git clone https://github.com/dewi911/serverfn.git
```

2. Создайте файл `.env` из примера:

```bash
make env
```

3. Настройте переменные окружения в файле `.env` если нужный порт занят.


## Запуск

### Локальный запуск

Для запуска сервера локально, установите следующие переменные окружения:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USERNAME=postgres
export DB_NAME=postgres
export DB_SSLMODE=disable
export DB_PASSWORD=qwerty
```
Затем выполните:
```bash
make run
```
Запуск с Docker
```bash
make docker-compose-up
```

## Команды

В проекте доступны следующие команды:

##### `make build:` Сборка проекта
##### `make run:` Запуск сервера
##### `make test:` Запуск тестов
##### `make docker-build:` Сборка Docker-образа
##### `make docker-compose-up:` Запуск сервисов через Docker Compose
##### `make docker-compose-down:` Остановка сервисов Docker Compose
##### `make swag-init:` Инициализация Swagger-документации
##### `make clean:` Очистка бинарных файлов
##### `make env:` Создание файла .env из примера


## Миграции

Для управления миграциями базы данных используются следующие команды:

Применение миграций:

```bash
make migrate-up
```

Откат миграций:

```bash
make migrate-down
```

Применение миграций без Docker:

```bash
migrate -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" -path migrations up
```

## Docker

Для работы с Docker используйте следующие команды:

Запуск сервисов:

```bash
docker-compose up -d
```

Остановка сервисов:

```bash
docker-compose down
```

## Postman

Для отправки запросов используйте следующие примеры:

```json
{
  "method": "GET",
  "url": "https://google.com",
  "headers": {
    "authentication": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "headers": {
      "Content-Type": "application/json",
      "Accept": "application/json",
      "User-Agent": "CustomClient/1.0",
      "X-Request-ID": "550e8400-e29b-41d4-a716-446655440000"
    },
    "HTTPStatusCode": 200,
    "responseHeaders": {},
    "responseLength": 55,
    "error": ""
  }
}
```
###### CRUD реализация тасков:

- Отправка С методом POST и пример запроса выше для создания таска
- Отправка С методом GET для получения списка всех тасков
```bash
http://localhost:8080/task
```
- Отправка С методом GET и айди таска для получения только этого таска
- Отправка С методом DELETE для удаления таска
- Отправка С методом PUT для обновления статуса таска
```bash
http://localhost:8080/task/1
```

## Swagger

Для наглядной проверки работоспособности сервера использовать сваггер:

```http request
http://localhost:8080/swagger/index.html
```



###### Сервер работает по принципу создания таска с ссылкой и методом запроса и после пинг на этот сайт где каждый таск обрабатывает отдельная горутинав которой есть возможность изменения количество активных горутин с реализацией очереди на каналах с обновлением статуса в базу данных при успешном пинге сайта

