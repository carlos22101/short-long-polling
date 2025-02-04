package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error al conectar con la base de datos: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Error al hacer ping a la base de datos: %v", err)
    }

    fmt.Println("Conexión a la base de datos exitosa")
}