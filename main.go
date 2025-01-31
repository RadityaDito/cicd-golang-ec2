package main

import (
    "golang-echo-crud/config"
    "golang-echo-crud/database"
    "golang-echo-crud/routes"

    "github.com/labstack/echo/v4"
)

func main() {
    config.InitDB()
    database.MigrateDB()

    e := echo.New()
    routes.InitRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
