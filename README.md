### Install deps
`go mod tidy`

### Create Migrations
`migrate create -ext sql -dir ./migrations -seq create_users_table`

### Execute Migrations
```migrate -path ./migrations -database "postgres://postgres:sameer@localhost:5432/postgres?sslmode=disable" up```

### Run server
`go run cmd/server/main.go`