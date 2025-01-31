package database

import (
    "fmt"
    "golang-echo-crud/config"
    "golang-echo-crud/models"
)

func MigrateDB() {
    err := config.DB.AutoMigrate(&models.User{})
    if err != nil {
        fmt.Println("Database migration failed:", err)
    } else {
        fmt.Println("Database migration successful")
    }
}
