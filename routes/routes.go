package routes

import (
	"golang-echo-crud/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.Home)
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUserByID)
	e.POST("/users", handlers.CreateUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)
}
