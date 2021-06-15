package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"open-prr/pkg/envs"
)

var connection *sql.DB

func GetConnection() (*sql.DB, error) {
	if connection == nil {
		fmt.Println("Creating a first connection with postgres database")
		host := envs.Getenv("POSTGRES_HOST", "0.0.0.0")
		port := envs.Getenv("POSTGRES_PORT", "5432")
		user := envs.Getenv("POSTGRES_USER", "user")
		password := envs.Getenv("POSTGRES_PASSWORD", "pass")
		database := envs.Getenv("POSTGRES_DB", "open-prr")
		sslMode := envs.Getenv("POSTGRES_SSL_MODE", "disable")
		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, database, sslMode)
		fmt.Println(connectionString)

		conn, err := sql.Open("postgres", connectionString)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		connection = conn
	}
	return connection, nil
}
