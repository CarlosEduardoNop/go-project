package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func OpenConnection() (*sql.DB, error) {
	godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	sc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), port, os.Getenv("DATABASE_NAME"))

	conn, err := sql.Open("mysql", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
