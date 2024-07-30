    ###  basic api server

если запускать без докера а локально на компе нужно экспортировать либо добавлять в голенде в параметрах запуска envirements
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USERNAME=postgres
export DB_NAME=postgres
export DB_SSLMODE=disable
export DB_PASSWORD=qwerty
```

```bash
docker-compose up -d
```

```bash
migrate -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" -path migrations up

```