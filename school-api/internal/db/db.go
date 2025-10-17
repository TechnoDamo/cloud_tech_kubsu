package db

import (
    "fmt"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Connect returns a gorm DB connection to Postgres using env vars.
// Expected env variables:
//   DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_SSLMODE
func Connect() (*gorm.DB, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASSWORD")
    name := os.Getenv("DB_NAME")
    ssl := os.Getenv("DB_SSLMODE")
    if ssl == "" {
        ssl = "disable"
    }
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, pass, name, ssl,
    )
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}


